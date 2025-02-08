package repository

import (
	"context"
	"database/sql"
	"fmt"
	"server/internal/datastruct"
	"server/internal/model"
	"server/util"
	"strings"
	"time"

	"github.com/blockloop/scan/v2"
	"github.com/bsm/redislock"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type PermissionRepository interface {
	RoleAccessList(ctx context.Context) ([]datastruct.PermissionRoleAccessList, error)
	RuleCreate(ctx context.Context, mdl model.Rule) error
}

type permissionRepository struct {
	sqlDB  *sql.DB
	sqlxDB *sqlx.DB
	cache  *redis.Client
}

func (m *permissionRepository) RoleAccessList(ctx context.Context) ([]datastruct.PermissionRoleAccessList, error) {
	data := make([]datastruct.PermissionRoleAccessList, 0)

	query := `
	SELECT 
		ro."uuid" as id, 
		ro."name", 
		COALESCE(
			subquery.perms, 
			'[]'
		) AS perms
	FROM roles ro
	LEFT JOIN LATERAL (
		SELECT 
			JSON_AGG(
				JSON_BUILD_OBJECT(
					'id', p.uuid,
					'code', p.code,
					'name', p.name,
					'parent', p.parent,
					'full_method', p.full_method
				)
			) AS perms
		FROM permissions p
		INNER JOIN rules ru 
			ON p.full_method = ru.v1 
			AND ru.p_type = 'p' 
			AND ru.v0 = ro.code
	) subquery ON TRUE
	ORDER BY ro."name";
	`

	sqlRows, err := m.sqlDB.QueryContext(ctx, query)
	if err != nil {
		return data, fmt.Errorf("13:%v", err)
	}

	if err := scan.Rows(&data, sqlRows); err != nil {
		return data, fmt.Errorf("5:%v", err)
	}

	for i, d := range data {
		data[i].Permissions = util.UnmarshalConverter[[]datastruct.Permission](d.PermissionsString)
	}

	return data, nil
}

func (m *permissionRepository) RuleCreate(ctx context.Context, mdl model.Rule) error {
	locker := redislock.New(m.cache)
	lock, err := locker.Obtain(ctx, "RuleCreate", 60*time.Second, nil)
	if err != nil {
		return fmt.Errorf("13:%v", err)
	}
	defer lock.Release(ctx)

	sqlRslt, err := m.sqlxDB.NamedExecContext(ctx, `
	insert into rules (p_type, v0, v1, v2, v3, v4, v5) values (:p_type, :v0, :v1, :v2, :v3, :v4, :v5)
	`, mdl)
	if err != nil {
		if strings.Contains(err.Error(), `duplicate key value violates unique constraint`) {
			return fmt.Errorf("6:%v", err)
		}
		return fmt.Errorf("13:%v", err)
	}
	rowsAffected, err := sqlRslt.RowsAffected()
	if err != nil {
		return fmt.Errorf("13:%v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("5:no data entered")
	}

	return nil
}
