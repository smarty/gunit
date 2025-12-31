package should_test

import (
	"testing"

	"github.com/smarty/gunit/v2/should"
)

func TestShouldContain(t *testing.T) {
	assert := NewAssertion(t)
	assert.ExpectedCountInvalid("actual", should.Contain)
	assert.ExpectedCountInvalid("actual", should.Contain, "EXPECTED", "EXTRA")

	assert.KindMismatch("string", should.Contain, false)
	assert.KindMismatch(1, should.Contain, "hi")

	// strings:
	assert.Fail("", should.Contain, "no")
	assert.Pass("integrate", should.Contain, "rat")
	assert.Pass("abc", should.Contain, 'b')

	// slices:
	assert.Fail([]byte("abc"), should.Contain, 'd')
	assert.Pass([]byte("abc"), should.Contain, 'b')
	assert.Pass([]byte("abc"), should.Contain, 98)

	// arrays:
	assert.Fail([3]byte{'a', 'b', 'c'}, should.Contain, 'd')
	assert.Pass([3]byte{'a', 'b', 'c'}, should.Contain, 'b')
	assert.Pass([3]byte{'a', 'b', 'c'}, should.Contain, 98)

	// maps:
	assert.Fail(map[rune]int{'a': 1}, should.Contain, 'b')
	assert.Pass(map[rune]int{'a': 1}, should.Contain, 'a')
}

func TestShouldNotContain(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.NOT.Contain)
	assert.ExpectedCountInvalid("actual", should.NOT.Contain, "EXPECTED", "EXTRA")

	assert.KindMismatch(false, should.NOT.Contain, "string")
	assert.KindMismatch("hi", should.NOT.Contain, 1)

	// strings:
	assert.Pass("", should.NOT.Contain, "no")
	assert.Fail("integrate", should.NOT.Contain, "rat")
	assert.Fail("abc", should.NOT.Contain, 'b')

	// slices:
	assert.Pass([]byte("abc"), should.NOT.Contain, 'd')
	assert.Fail([]byte("abc"), should.NOT.Contain, 'b')
	assert.Fail([]byte("abc"), should.NOT.Contain, 98)

	// arrays:
	assert.Pass([3]byte{'a', 'b', 'c'}, should.NOT.Contain, 'd')
	assert.Fail([3]byte{'a', 'b', 'c'}, should.NOT.Contain, 'b')
	assert.Fail([3]byte{'a', 'b', 'c'}, should.NOT.Contain, 98)

	// maps:
	assert.Pass(map[rune]int{'a': 1}, should.NOT.Contain, 'b')
	assert.Fail(map[rune]int{'a': 1}, should.NOT.Contain, 'a')
}
