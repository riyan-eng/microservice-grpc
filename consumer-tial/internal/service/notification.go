package service

import (
	"consumer-trial/internal/entity"
	"consumer-trial/internal/repository"
	"context"
	"fmt"
	"log"
)

type Notification interface {
	WhatsApp(ctx context.Context, ent entity.Message) error
	Email(ctx context.Context, ent entity.Message) error
	PushMobile(ctx context.Context, ent entity.Message) error
}

type notification struct {
	dao repository.DAO
}

func NewNotification(dao *repository.DAO) Notification {
	return &notification{
		dao: *dao,
	}
}

func (m *notification) WhatsApp(ctx context.Context, ent entity.Message) error {
	log.Println(ent)
	return fmt.Errorf("gagal kirim wa ke:%v :%v", ent.Id, ent.Body.To)
}

func (m *notification) Email(ctx context.Context, ent entity.Message) error {
	if err := m.dao.NewNotificationRepository().WhatsApp(ctx, ent.Id, ent.Status); err != nil {
		return err
	}
	return nil
}

func (m *notification) PushMobile(ctx context.Context, ent entity.Message) error {
	return nil
}
