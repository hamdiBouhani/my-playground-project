package rate_limiter

import (
	"fmt"
	"testing"
	"time"
)

var rateLimiter *RateLimiter

func init() {
	// Create a rate limiter with a rate of 5 requests per second and a bucket capacity of 10 tokens
	rateLimiter = NewRateLimiter(5.0, 10.0)
}

func TestRateLimiter(t *testing.T) {
	// Simulate 15 requests
	for i := 0; i < 15; i++ {
		if rateLimiter.Allow() {
			fmt.Printf("Request %d: Allowed\n", i+1)
		} else {
			fmt.Printf("Request %d: Denied\n", i+1)
		}

		// Introduce a delay between requests
		time.Sleep(time.Second)
	}
}
