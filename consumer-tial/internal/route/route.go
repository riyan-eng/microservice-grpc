package route

import (
	"consumer-trial/internal/entity"
	"consumer-trial/internal/service"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bytedance/sonic"
)

type route struct {
	service service.Notification
}

func NewRoute(service service.Notification) *route {
	return &route{
		service: service,
	}
}

func (m *route) Route(status, messageString string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var message entity.Message
	if err := sonic.UnmarshalString(messageString, &message); err != nil {
		return err
	}

	message.Status = status

	switch message.Type {
	case "WHATSAPP":
		if err := m.service.WhatsApp(ctx, message); err != nil {
			log.Println(err)
			return err
		}
	case "EMAIL":
		if err := m.service.Email(ctx, message); err != nil {
			log.Println(err)
			return err
		}
	case "PUSH_MOBILE":
		if err := m.service.PushMobile(ctx, message); err != nil {
			log.Println(err)
			return err
		}
	default:
		return fmt.Errorf("message type invalid")
	}
	return nil
}
