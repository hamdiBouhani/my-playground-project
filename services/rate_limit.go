package services

import (
	"time"
)

type RateLimiter struct {
	RequestsPerSecond float64 //rate
	capacity          float64 // bucket capacity (tokens)
	tokens            float64 // current number of tokens in the bucket
	LastRequestTime   time.Time
}

func NewRateLimiter(rate float64, capacity float64) *RateLimiter {
	return &RateLimiter{
		RequestsPerSecond: rate,
		capacity:          capacity,
		tokens:            capacity,
	}
}

func (r *RateLimiter) Allow() bool {
	now := time.Now()

	elapsed := now.Sub(r.LastRequestTime).Seconds()
	r.tokens += elapsed * r.RequestsPerSecond

	// Clamp the tokens to the bucket's capacity
	if r.tokens > r.capacity {
		r.tokens = r.capacity
	}

	// Check if enough tokens are available for the request
	if r.tokens >= 1.0 {
		r.tokens -= 1.0
		r.LastRequestTime = now
		return true
	}

	return false
}
