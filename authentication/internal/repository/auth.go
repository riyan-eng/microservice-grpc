package repository

import (
	"context"
	"database/sql"
	"fmt"
	"server/internal/datastruct"
	"server/util"

	"github.com/blockloop/scan/v2"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type AuthRepository interface {
	Login(ctx context.Context, username string) (*datastruct.AuthLoginData, error)
	Logout(ctx context.Context, userId string) error
	Me(ctx context.Context, userId string) (*datastruct.AuthMe, error)
}

type authRepository struct {
	sqlDB  *sql.DB
	sqlxDB *sqlx.DB
	cache  *redis.Client
}

func (m *authRepository) Login(ctx context.Context, username string) (*datastruct.AuthLoginData, error) {
	data := new(datastruct.AuthLoginData)

	query := fmt.Sprintf(`
	select u."uuid", u.email, u.username, u."password", u.is_active, r.code as role_code, r."name" as role_name 
	from users u 
	left join user_datas ud on ud.user_uuid = u."uuid" 
	left join roles r on r.code = ud.role_code
	where u.email = '%s' and u.deleted_at is null
	limit 1
	`, username)

	sqlRows, err := m.sqlDB.QueryContext(ctx, query)
	if err != nil {
		return data, fmt.Errorf("13:%v", err)
	}

	if err := scan.Row(data, sqlRows); err != nil {
		if err == sql.ErrNoRows {
			return data, fmt.Errorf("5:email or username not registered")
		}
		return data, fmt.Errorf("13:%v", err)
	}

	return data, nil
}

func (m *authRepository) Logout(ctx context.Context, userId string) error {
	if err := m.cache.Del(ctx, fmt.Sprintf("access-token-%s", userId)).Err(); err != nil {
		return fmt.Errorf("13:%v", err)
	}
	if err := m.cache.Del(ctx, fmt.Sprintf("refresh-token-%s", userId)).Err(); err != nil {
		return fmt.Errorf("13:%v", err)
	}
	return nil
}

func (m *authRepository) Me(ctx context.Context, userId string) (*datastruct.AuthMe, error) {
	data := new(datastruct.AuthMe)

	query := fmt.Sprintf(`
	select u."uuid", u.username, r.code as role_code, r."name" as role_name, u.photo_url, (
		select json_agg(p.code) 
		from rules r 
		left join permissions p on r.v1 = p.full_method
		where r.v0 = ud.role_code
	) as permissions
	from users u 
	left join user_datas ud on ud.user_uuid = u.uuid 
	left join roles r on r.code = ud.role_code
	where u.uuid = '%v' 
	limit 1
	`, userId)

	sqlRows, err := m.sqlDB.QueryContext(ctx, query)
	if err != nil {
		return data, fmt.Errorf("13:%v", err)
	}

	if err := scan.Row(data, sqlRows); err != nil {
		return data, fmt.Errorf("5:%v", err)
	}

	data.Permissions = util.UnmarshalConverter[[]string](data.PermissionsStr)
	return data, nil
}
