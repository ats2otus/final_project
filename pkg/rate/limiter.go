package rate

import (
	"sync"
	"time"
)

type Limiter interface {
	Allow(key string) bool
	Reset(key string)
}

type limiter struct {
	limit    int
	interval time.Duration

	mu      sync.RWMutex
	buckets map[int64]*bucket
}

func (l *limiter) Allow(key string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := l.timestamp()
	if _, ok := l.buckets[now]; !ok {
		l.buckets[now] = &bucket{
			recs: make(map[string]int),
		}
	}

	return l.buckets[now].incr(key) <= l.limit
}

func (l *limiter) Reset(key string) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	l.buckets[l.timestamp()].reset(key)
}

// округляет timestamp до interval, что бы в течении текущего interval результат всегда был один
func (l *limiter) timestamp() int64 {
	return time.Now().Truncate(l.interval).Unix()
}

func (l *limiter) background() {
	for range time.NewTicker(l.interval).C {
		l.cleanup()
	}
}

func (l *limiter) cleanup() {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := l.timestamp()
	for ts := range l.buckets {
		if ts < now {
			delete(l.buckets, ts)
		}
	}
}

func NewLimiter(interval time.Duration, limit int) Limiter {
	l := limiter{
		interval: interval, limit: limit,
		buckets: make(map[int64]*bucket),
	}

	go l.background()

	return &l
}
