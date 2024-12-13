package service

import (
	"context"
	"fmt"
	"server/internal/entity"
	"server/internal/model"
	"server/internal/repository"

	"github.com/bytedance/sonic"
)

type BackgroundService interface {
	Create(ctx context.Context, ent *entity.BackgroundCreate) error
}

type backgroundService struct {
	dao repository.DAO
}

func NewBackgroundService(dao *repository.DAO) BackgroundService {
	return &backgroundService{
		dao: *dao,
	}
}

func (m *backgroundService) Create(ctx context.Context, ent *entity.BackgroundCreate) error {
	message, err := sonic.Marshal(ent)
	if err != nil {
		return fmt.Errorf("13:%v", err)
	}

	mdl := model.Background{
		Id:     ent.Id,
		Status: "WAITING",
		Type:   ent.Type,
		Body:   string(message),
	}

	if err := m.dao.NewBackgroundRepository().Create(ctx, message, &mdl); err != nil {
		return err
	}
	return nil
}
