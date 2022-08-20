package rate

import "testing"

func Test_Bucket(t *testing.T) {
	key := "test"
	b := newBucket()

	for i := 1; i <= 10; i++ {
		if val := b.incr(key); val != i {
			t.Fatalf("expected %d got %d", i, val)
		}
	}
	b.reset(key)
	if val := b.incr(key); val != 1 {
		t.Fatalf("expected 1 got %d", val)
	}
}
