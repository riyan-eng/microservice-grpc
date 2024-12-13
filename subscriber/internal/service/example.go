package service

import (
	"context"
	"server/internal/entity"
	"server/internal/model"
	"server/internal/repository"
)

type ExampleService interface {
	Create(ctx context.Context, ent *entity.ExampleCreate) error
	Put(ctx context.Context, ent *entity.ExamplePut) error
}

type exampleService struct {
	dao repository.DAO
}

func NewExampleService(dao *repository.DAO) ExampleService {
	return &exampleService{
		dao: *dao,
	}
}

func (m *exampleService) Create(ctx context.Context, ent *entity.ExampleCreate) error {
	mdl := model.Example{
		Id:     ent.Id,
		Name:   ent.Name,
		Detail: ent.Detail,
	}

	if err := m.dao.NewExampleRepository().Create(ctx, &mdl); err != nil {
		// custom err
		return err
	}

	return nil
}

func (m *exampleService) Put(ctx context.Context, ent *entity.ExamplePut) error {
	mdl := model.Example{
		Id:     ent.Id,
		Name:   ent.Name,
		Detail: ent.Detail,
	}

	if err := m.dao.NewExampleRepository().Put(ctx, &mdl); err != nil {
		// custom err
		return err
	}

	return nil
}
