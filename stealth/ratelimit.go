package stealth

import "time"

// RateLimiter simulates human action limits
type RateLimiter struct {
	MaxActions int
	Count      int
	LastReset  time.Time
}

// NewRateLimiter creates a new limiter
func NewRateLimiter(max int) *RateLimiter {
	return &RateLimiter{
		MaxActions: max,
		LastReset:  time.Now(),
	}
}

// Allow checks if another action is allowed
func (r *RateLimiter) Allow() bool {
	// reset every day
	if time.Since(r.LastReset) > 24*time.Hour {
		r.Count = 0
		r.LastReset = time.Now()
	}

	if r.Count >= r.MaxActions {
		return false
	}

	r.Count++
	return true
}
