package should_test

import (
	"testing"
	"time"

	"github.com/smarty/gunit/v2/assert/should"
)

func TestShouldHappenWithin(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid(time.Now(), should.HappenWithin)
	assert.ExpectedCountInvalid(time.Now(), should.HappenWithin, time.Nanosecond)

	assert.TypeMismatch(1, should.HappenWithin, time.Nanosecond, time.Now())
	assert.TypeMismatch(time.Now(), should.HappenWithin, 1, time.Now())
	assert.TypeMismatch(time.Now(), should.HappenWithin, time.Nanosecond, 1)

	assert.Fail(time.Now(), should.HappenWithin, time.Nanosecond, time.Now().Truncate(time.Millisecond))
	assert.Pass(time.Now(), should.HappenWithin, time.Second, time.Now())
}
