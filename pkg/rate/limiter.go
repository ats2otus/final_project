package rate

import (
	"time"
)

type Limiter interface {
	Allow(key string) bool
	Reset(key string)
}

type limiter struct {
	limit    int
	interval time.Duration
}

func (l *limiter) Allow(key string) bool {
	//TODO:
	return true
}

func (l *limiter) Reset(key string) {
	//TODO:
}

func NewLimiter(interval time.Duration, limit int) Limiter {
	return &limiter{
		interval: interval, limit: limit,
	}
}
