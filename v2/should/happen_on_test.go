package should_test

import (
	"testing"
	"time"

	"github.com/smarty/gunit/v2/should"
)

func TestShouldHappenOn(t *testing.T) {
	assert := NewAssertion(t)

	assert.TypeMismatch(1, should.HappenOn, time.Now())
	assert.TypeMismatch(time.Now(), should.HappenOn, 1)

	assert.ExpectedCountInvalid(time.Now(), should.HappenOn)
	assert.ExpectedCountInvalid(time.Now(), should.HappenOn, time.Now(), time.Now())

	now := time.Now()
	assert.Pass(now.UTC(), should.HappenOn, now.In(time.Local))
	assert.Fail(time.Now(), should.HappenOn, time.Now())
}

func TestShouldNOTHappenOn(t *testing.T) {
	assert := NewAssertion(t)

	assert.TypeMismatch(1, should.NOT.HappenOn, time.Now())
	assert.TypeMismatch(time.Now(), should.NOT.HappenOn, 1)

	assert.ExpectedCountInvalid(time.Now(), should.NOT.HappenOn)
	assert.ExpectedCountInvalid(time.Now(), should.NOT.HappenOn, time.Now(), time.Now())

	now := time.Now()
	assert.Fail(now.UTC(), should.NOT.HappenOn, now.In(time.Local))
	assert.Pass(time.Now(), should.NOT.HappenOn, time.Now())
}
