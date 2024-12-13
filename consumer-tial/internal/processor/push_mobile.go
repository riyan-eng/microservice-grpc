package processor

import (
	"consumer-trial/internal/entity"
	"context"
	"fmt"
	"log"

	"golang.org/x/time/rate"
)

// Define the rate limiter
var limiterPushMobile = rate.NewLimiter(2, 2) // 200 requests per second with a burst of 200
// var limiterPushMobileOnce sync.Once

func PushMobile(message entity.MessageBody) error {
	// limiterPushMobileOnce.Do(func() {
	// 	log.Println("Rate limiter initialized for 200 calls per second.")
	// })

	// Wait for permission to proceed
	if err := limiterPushMobile.Wait(context.TODO()); err != nil {
		return fmt.Errorf("rate limiter error: %v", err)
	}

	log.Println("sukses kirim push mobile ke:", message.To)
	return nil
}
