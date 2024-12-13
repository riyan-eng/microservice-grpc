package processor

import (
	"consumer-trial/internal/entity"
	"context"
	"fmt"
	"log"

	"golang.org/x/time/rate"
)

// Define the rate limiter
var limiterEmail = rate.NewLimiter(2, 2) // 200 requests per second with a burst of 200
// var limiterEmailOnce sync.Once

func Email(message entity.MessageBody) error {
	// limiterEmailOnce.Do(func() {
	// 	log.Println("Rate limiter initialized for 200 calls per second.")
	// })

	// Wait for permission to proceed
	if err := limiterEmail.Wait(context.TODO()); err != nil {
		return fmt.Errorf("rate limiter error: %v", err)
	}

	// call api
	// log.Println("sukses kirim email ke:", message.To)
	log.Println(message)
	return fmt.Errorf("gagal kirim email ke:%v", message.To)
}
