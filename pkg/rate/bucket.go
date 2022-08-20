package rate

import "sync"

type bucket struct {
	mu   sync.Mutex
	recs map[string]int
}

func (b *bucket) incr(key string) int {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.recs[key]++
	return b.recs[key]
}

func (b *bucket) reset(key string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	delete(b.recs, key)
}

func newBucket() *bucket {
	return &bucket{
		recs: make(map[string]int),
	}
}
