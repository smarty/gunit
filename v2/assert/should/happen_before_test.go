package should_test

import (
	"testing"
	"time"

	"github.com/smarty/gunit/v2/assert/should"
)

func TestShouldHappenBefore(t *testing.T) {
	assert := NewAssertion(t)

	assert.TypeMismatch(1, should.HappenBefore, time.Now())
	assert.TypeMismatch(time.Now(), should.HappenBefore, 1)

	assert.ExpectedCountInvalid(time.Now(), should.HappenBefore)
	assert.ExpectedCountInvalid(time.Now(), should.HappenBefore, time.Now(), time.Now())

	assert.Fail(time.Now().Add(time.Second), should.HappenBefore, time.Now())
	assert.Pass(time.Now(), should.HappenBefore, time.Now())
}
