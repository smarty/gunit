package should_test

import (
	"testing"
	"time"

	"github.com/smarty/gunit/v2/should"
)

func TestShouldBeChronological(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.BeChronological, "EXTRA")

	assert.TypeMismatch(42, should.BeChronological)

	var (
		a = time.Now()
		b = a.Add(time.Nanosecond)
		c = b.Add(time.Nanosecond)
	)
	assert.Pass([]time.Time{}, should.BeChronological)
	assert.Pass([]time.Time{a, a, a}, should.BeChronological)
	assert.Pass([]time.Time{a, b, c}, should.BeChronological)
	assert.Fail([]time.Time{a, c, b}, should.BeChronological)
}

func TestShouldNOTBeChronological(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.NOT.BeChronological, "EXTRA")

	assert.TypeMismatch(42, should.NOT.BeChronological)

	var (
		a = time.Now()
		b = a.Add(time.Nanosecond)
		c = b.Add(time.Nanosecond)
	)
	assert.Fail([]time.Time{}, should.NOT.BeChronological)
	assert.Fail([]time.Time{a, a, a}, should.NOT.BeChronological)
	assert.Fail([]time.Time{a, b, c}, should.NOT.BeChronological)
	assert.Pass([]time.Time{a, c, b}, should.NOT.BeChronological)
}
