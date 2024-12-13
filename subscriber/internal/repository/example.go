package repository

import (
	"context"
	"database/sql"
	"fmt"
	"server/internal/datastruct"
	"server/internal/model"
	"strings"

	"github.com/blockloop/scan/v2"
	"github.com/jmoiron/sqlx"
)

type ExampleRepository interface {
	List(ctx context.Context, search string, limit, offset int32) ([]datastruct.ExampleList, int32, error)
	Create(ctx context.Context, mdl *model.Example) error
	Put(ctx context.Context, mdl *model.Example) error
	Detail(ctx context.Context, id string) (*datastruct.ExampleDetail, error)
	Delete(ctx context.Context, id string) error
}

type exampleRepository struct {
	sqlDB  *sql.DB
	sqlxDB *sqlx.DB
}

func (m *exampleRepository) Create(ctx context.Context, mdl *model.Example) error {
	sqlRslt, err := m.sqlxDB.NamedExecContext(ctx, `
	insert into example (uuid, detail, name) values (:id, :detail, :name)
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

func (m *exampleRepository) Put(ctx context.Context, mdl *model.Example) error {
	sqlRslt, err := m.sqlxDB.NamedExecContext(ctx, `
	update example set detail=coalesce(:detail, detail), name=coalesce(:name, name), updated_at=now() where uuid=:id
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
		return fmt.Errorf("5:no data updated")
	}

	return nil
}

func (m *exampleRepository) List(ctx context.Context, search string, limit, offset int32) ([]datastruct.ExampleList, int32, error) {
	data := make([]datastruct.ExampleList, 0)
	var countRow int32

	query := fmt.Sprintf(`
	select e."uuid", e.name, e.detail,
	count(e.uuid) over() as total_rows
	from example e 
	where lower(e.name) like lower('%%%v%%') order by e.created_at %v limit %v offset %v
	`, search, "desc", limit, offset)
	sqlRows, err := m.sqlDB.QueryContext(ctx, query)
	if err != nil {
		return data, countRow, fmt.Errorf("13:%v", err)
	}

	if err := scan.Rows(&data, sqlRows); err != nil {
		return data, countRow, fmt.Errorf("13:%v", err)
	}

	for _, d := range data {
		countRow = d.TotalRows
		break
	}

	return data, countRow, nil
}

func (m *exampleRepository) Detail(ctx context.Context, id string) (*datastruct.ExampleDetail, error) {
	data := new(datastruct.ExampleDetail)

	query := fmt.Sprintf(`
	select e."uuid", e.name, e.detail
	from example e
	where e.uuid = '%v'
	`, id)
	sqlRows, err := m.sqlDB.QueryContext(ctx, query)
	if err != nil {
		return data, fmt.Errorf("13:%v", err)
	}

	if err := scan.Row(data, sqlRows); err != nil {
		return data, fmt.Errorf("5:%v", err)
	}

	return data, nil
}

func (m *exampleRepository) Delete(ctx context.Context, id string) error {
	sqlRslt, err := m.sqlxDB.ExecContext(ctx, fmt.Sprintf("delete from example where uuid = '%v'", id))
	if err != nil {
		return fmt.Errorf("13:%v", err)
	}

	rowsAffected, err := sqlRslt.RowsAffected()
	if err != nil {
		return fmt.Errorf("13:%v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("5:no data deleted")
	}

	return nil
}
