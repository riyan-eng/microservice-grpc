package processor

import (
	"consumer-trial/internal/entity"
	"context"
	"fmt"
	"log"

	"golang.org/x/time/rate"
)

// Define the rate limiter
var limiterWhatsApp = rate.NewLimiter(2, 2) // 200 requests per second with a burst of 200
// var limiterWhatsAppOnce sync.Once

func WhatsApp(message entity.MessageBody) error {
	// limiterWhatsAppOnce.Do(func() {
	// 	log.Println("Rate limiter initialized for 200 calls per second.")
	// })

	// Wait for permission to proceed
	if err := limiterWhatsApp.Wait(context.TODO()); err != nil {
		return fmt.Errorf("rate limiter error: %v", err)
	}

	// call api
	log.Println("sukses kirim wa ke:", message.To)
	return nil
}
