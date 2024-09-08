package repository

import (
	"context"
	"database/sql"
	"fmt"
	"server/internal/datastruct"
	"server/internal/model"
	"server/util"
	"strings"

	"github.com/blockloop/scan/v2"
	"github.com/jmoiron/sqlx"
)

type ExampleRepository interface {
	List(ctx *context.Context, search *string, limit, offset *int32) (*[]datastruct.ExampleList, *int32, *util.Error)
	Create(ctx *context.Context, mdl *model.Example) *util.Error
	Put(ctx *context.Context, mdl *model.Example) *util.Error
	Detail(ctx *context.Context, id *string) (*datastruct.ExampleDetail, *util.Error)
	Delete(ctx *context.Context, id *string) *util.Error
}

type exampleRepository struct {
	sqlDB  *sql.DB
	sqlxDB *sqlx.DB
}

func (m *exampleRepository) Create(ctx *context.Context, mdl *model.Example) *util.Error {
	sqlRslt, err := m.sqlxDB.NamedExecContext(*ctx, `
	insert into example (uuid, detail, name) values (:id, :detail, :name)
	`, mdl)
	if err != nil {
		if strings.Contains(err.Error(), `duplicate key value violates unique constraint`) {
			return &util.Error{
				Errors:     err,
				StatusCode: 6,
			}
		}
		return &util.Error{
			Errors:     err,
			StatusCode: 13,
		}
	}
	rowsAffected, err := sqlRslt.RowsAffected()
	if err != nil {
		return &util.Error{
			Errors:     err,
			StatusCode: 13,
		}
	}
	if rowsAffected == 0 {
		return &util.Error{
			Errors:     fmt.Errorf("no data entered"),
			StatusCode: 5,
		}
	}

	return &util.Error{}
}

func (m *exampleRepository) Put(ctx *context.Context, mdl *model.Example) *util.Error {
	sqlRslt, err := m.sqlxDB.NamedExecContext(*ctx, `
	update example set detail=coalesce(:detail, detail), name=coalesce(:name, name), updated_at=now() where uuid=:id
	`, mdl)
	if err != nil {
		if strings.Contains(err.Error(), `duplicate key value violates unique constraint`) {
			return &util.Error{
				Errors:     err,
				StatusCode: 6,
			}
		}
		return &util.Error{
			Errors:     err,
			StatusCode: 13,
		}
	}
	rowsAffected, err := sqlRslt.RowsAffected()
	if err != nil {
		return &util.Error{
			Errors:     err,
			StatusCode: 13,
		}
	}
	if rowsAffected == 0 {
		return &util.Error{
			Errors:     fmt.Errorf("no data updated"),
			StatusCode: 5,
		}
	}

	return &util.Error{}
}

func (m *exampleRepository) List(ctx *context.Context, search *string, limit, offset *int32) (*[]datastruct.ExampleList, *int32, *util.Error) {
	data := make([]datastruct.ExampleList, 0)
	countRow := new(int32)

	query := fmt.Sprintf(`
	select e."uuid", e.name, e.detail,
	count(e.uuid) over() as total_rows
	from example e 
	where lower(e.name) like lower('%%%v%%') order by e.created_at %v limit %v offset %v
	`, *search, "desc", *limit, *offset)
	sqlRows, err := m.sqlDB.QueryContext(*ctx, query)
	if err != nil {
		return &data, countRow, &util.Error{
			Errors:     err,
			StatusCode: 13,
		}
	}

	if err := scan.Rows(&data, sqlRows); err != nil {
		return &data, countRow, &util.Error{
			Errors:     err,
			StatusCode: 13,
		}
	}

	for _, d := range data {
		countRow = &d.TotalRows
		break
	}

	return &data, countRow, &util.Error{}
}

func (m *exampleRepository) Detail(ctx *context.Context, id *string) (*datastruct.ExampleDetail, *util.Error) {
	data := new(datastruct.ExampleDetail)

	query := fmt.Sprintf(`
	select e."uuid", e.name, e.detail
	from example e
	where e.uuid = '%v'
	`, *id)
	sqlRows, err := m.sqlDB.QueryContext(*ctx, query)
	if err != nil {
		return data, &util.Error{
			Errors:     err,
			StatusCode: 13,
		}
	}

	if err := scan.Row(data, sqlRows); err != nil {
		return data, &util.Error{
			Errors:     err,
			StatusCode: 5,
		}
	}

	return data, &util.Error{}
}

func (m *exampleRepository) Delete(ctx *context.Context, id *string) *util.Error {
	sqlRslt, err := m.sqlxDB.ExecContext(*ctx, fmt.Sprintf("delete from example where uuid = '%v'", *id))
	if err != nil {
		return &util.Error{
			Errors:     err,
			StatusCode: 13,
		}
	}

	rowsAffected, err := sqlRslt.RowsAffected()
	if err != nil {
		return &util.Error{
			Errors:     err,
			StatusCode: 13,
		}
	}
	if rowsAffected == 0 {
		return &util.Error{
			Errors:     fmt.Errorf("no data deleted"),
			StatusCode: 5,
		}
	}

	return &util.Error{}
}
