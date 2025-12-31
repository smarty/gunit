package should_test

import (
	"testing"

	"github.com/smarty/gunit/v2/assert/should"
)

func TestShouldStartWith(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.StartWith)
	assert.ExpectedCountInvalid("actual", should.StartWith, "EXPECTED", "EXTRA")

	assert.KindMismatch("string", should.StartWith, false)
	assert.KindMismatch(1, should.StartWith, "hi")

	// strings:
	assert.Fail("", should.StartWith, "no")
	assert.Pass("abc", should.StartWith, 'a')
	assert.Pass("integrate", should.StartWith, "in")

	// slices:
	assert.Fail([]byte{}, should.StartWith, 'b')
	assert.Fail([]byte(nil), should.StartWith, 'b')
	assert.Fail([]byte("abc"), should.StartWith, 'b')
	assert.Pass([]byte("abc"), should.StartWith, 'a')
	assert.Pass([]byte("abc"), should.StartWith, 97)

	// arrays:
	assert.Fail([3]byte{'a', 'b', 'c'}, should.StartWith, 'b')
	assert.Pass([3]byte{'a', 'b', 'c'}, should.StartWith, 'a')
	assert.Pass([3]byte{'a', 'b', 'c'}, should.StartWith, 97)
}
