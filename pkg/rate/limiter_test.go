package rate

import (
	"fmt"
	"testing"
	"time"
)

func Test_Limiter(t *testing.T) {
	rate := limiter{
		buckets: map[int64]*bucket{
			1: newBucket(), 2: newBucket(), // заведомо устарешвние бакеты
		},
		interval: time.Hour, limit: 2,
	}

	total := 10
	for i := 0; i < total; i++ {
		rate.Allow(fmt.Sprintf("key-%d", i))
	}

	// бакет текущего интервала
	bucket := rate.buckets[rate.timestamp()]
	if len(bucket.recs) != total {
		t.Fatalf("len of current bucket have to be %d, actual size %d", total, len(rate.buckets))
	}

	// убедимся, что cleanup удалит старые бакеты
	rate.cleanup()
	if len(rate.buckets) != 1 {
		t.Fatalf("after cleanup expected len of buckets 1 actual %d", len(rate.buckets))
	}

	// бакет текущего интервала
	bucket = rate.buckets[rate.timestamp()]
	if len(bucket.recs) != total {
		t.Fatalf("len of current bucket have to be %d, actual size %d", total, len(rate.buckets))
	}

	for i := 0; i < total; i++ {
		rate.Reset(fmt.Sprintf("key-%d", i))
	}
	if len(bucket.recs) != 0 {
		t.Fatalf("len of current bucket have to be %d, actual size %d", 0, len(rate.buckets))
	}
}
