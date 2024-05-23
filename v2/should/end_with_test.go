package should_test

import (
	"testing"

	"github.com/smarty/gunit/v2/should"
)

func TestShouldEndWith(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.EndWith)
	assert.ExpectedCountInvalid("actual", should.EndWith, "EXPECTED", "EXTRA")

	assert.KindMismatch("string", should.EndWith, false)
	assert.KindMismatch(1, should.EndWith, "hi")

	// strings:
	assert.Fail("", should.EndWith, "no")
	assert.Pass("abc", should.EndWith, 'c')
	assert.Pass("integrate", should.EndWith, "ate")

	// slices:
	assert.Fail([]byte{}, should.EndWith, 'b')
	assert.Fail([]byte(nil), should.EndWith, 'b')
	assert.Fail([]byte("abc"), should.EndWith, 'b')
	assert.Pass([]byte("abc"), should.EndWith, 'c')
	assert.Pass([]byte("abc"), should.EndWith, 99)

	// arrays:
	assert.Fail([3]byte{'a', 'b', 'c'}, should.EndWith, 'b')
	assert.Pass([3]byte{'a', 'b', 'c'}, should.EndWith, 'c')
	assert.Pass([3]byte{'a', 'b', 'c'}, should.EndWith, 99)
}
