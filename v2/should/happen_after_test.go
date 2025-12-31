package should_test

import (
	"testing"
	"time"

	"github.com/smarty/gunit/v2/should"
)

func TestShouldHappenAfter(t *testing.T) {
	assert := NewAssertion(t)

	assert.TypeMismatch(1, should.HappenAfter, time.Now())
	assert.TypeMismatch(time.Now(), should.HappenAfter, 1)

	assert.ExpectedCountInvalid(time.Now(), should.HappenAfter)
	assert.ExpectedCountInvalid(time.Now(), should.HappenAfter, time.Now(), time.Now())

	assert.Fail(time.Now(), should.HappenAfter, time.Now())
	assert.Pass(time.Now().Add(time.Second), should.HappenAfter, time.Now())
}
