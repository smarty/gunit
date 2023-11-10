package gunit

import (
	"math"
	"testing"
	"time"
)

func TestEqual(t *testing.T) {
	shouldNotEqual(t, 1, 2)
	shouldEqual(t, 1, 1)
	shouldEqual(t, 1, uint(1))

	now := time.Now()
	shouldEqual(t, now.UTC(), now.In(time.Local))
	shouldNotEqual(t, time.Now(), time.Now())

	shouldNotEqual(t, struct{ A string }{}, struct{ B string }{})
	shouldEqual(t, struct{ A string }{}, struct{ A string }{})

	shouldNotEqual(t, []byte("hi"), []byte("bye"))
	shouldEqual(t, []byte("hi"), []byte("hi"))

	const MAX uint64 = math.MaxUint64
	shouldNotEqual(t, -1, MAX)
	shouldNotEqual(t, MAX, -1)

	shouldEqual(t, returnsNilInterface(), nil)
}
func shouldEqual(t *testing.T, a, b any) {
	if err := equal(a, b); err != nil {
		t.Error(err)
	}
}
func shouldNotEqual(t *testing.T, a, b any) {
	if err := equal(a, b); err == nil {
		t.Error("values were equal (but shouldn't have been):", a, b)
	}
}

func returnsNilInterface() any { return nil }
