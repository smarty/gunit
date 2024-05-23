package should_test

import (
	"math"
	"testing"
	"time"

	"github.com/smarty/gunit/v2/should"
)

func TestShouldEqual(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.Equal)
	assert.ExpectedCountInvalid("actual", should.Equal, "EXPECTED", "EXTRA")

	assert.Fail(1, should.Equal, 2)
	assert.Pass(1, should.Equal, 1)
	assert.Pass(1, should.Equal, uint(1))

	now := time.Now()
	assert.Pass(now.UTC(), should.Equal, now.In(time.Local))
	assert.Fail(time.Now(), should.Equal, time.Now())

	assert.Fail(struct{ A string }{}, should.Equal, struct{ B string }{})
	assert.Pass(struct{ A string }{}, should.Equal, struct{ A string }{})

	assert.Fail([]byte("hi"), should.Equal, []byte("bye"))
	assert.Pass([]byte("hi"), should.Equal, []byte("hi"))

	const max uint64 = math.MaxUint64
	assert.Fail(-1, should.Equal, max)
	assert.Fail(max, should.Equal, -1)

	assert.Pass(returnsNilInterface(), should.Equal, nil)
}

func TestShouldNotEqual(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.NOT.Equal)
	assert.ExpectedCountInvalid("actual", should.NOT.Equal, "EXPECTED", "EXTRA")

	assert.Fail(1, should.NOT.Equal, 1)
	assert.Pass(1, should.NOT.Equal, 2)
}

func returnsNilInterface() any { return nil }
