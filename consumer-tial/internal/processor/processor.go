package processor

import (
	"consumer-trial/internal/entity"
	"fmt"
	"log"

	"github.com/bytedance/sonic"
)

func NotifSender(messageString string) error {
	var message entity.Message
	if err := sonic.UnmarshalString(messageString, &message); err != nil {
		return err
	}
	switch message.Type {
	case "WHATSAPP":
		if err := WhatsApp(message.Body); err != nil {
			log.Println(err)
			return err
		}
	case "EMAIL":
		if err := Email(message.Body); err != nil {
			log.Println(err)
			return err
		}
	case "PUSH_MOBILE":
		if err := PushMobile(message.Body); err != nil {
			log.Println(err)
			return err
		}
	default:
		return fmt.Errorf("message type invalid")
	}
	return nil
}
