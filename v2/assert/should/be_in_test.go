package should_test

import (
	"testing"

	"github.com/smarty/gunit/v2/assert/should"
)

func TestShouldBeIn(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.BeIn)
	assert.ExpectedCountInvalid("actual", should.BeIn, "EXPECTED", "EXTRA")

	assert.KindMismatch(false, should.BeIn, "string")
	assert.KindMismatch("hi", should.BeIn, 1)

	// strings:
	assert.Fail("no", should.BeIn, "")
	assert.Pass("rat", should.BeIn, "integrate")
	assert.Pass('b', should.BeIn, "abc")

	// slices:
	assert.Fail('d', should.BeIn, []byte("abc"))
	assert.Pass('b', should.BeIn, []byte("abc"))
	assert.Pass(98, should.BeIn, []byte("abc"))

	// arrays:
	assert.Fail('d', should.BeIn, [3]byte{'a', 'b', 'c'})
	assert.Pass('b', should.BeIn, [3]byte{'a', 'b', 'c'})
	assert.Pass(98, should.BeIn, [3]byte{'a', 'b', 'c'})

	// maps:
	assert.Fail('b', should.BeIn, map[rune]int{'a': 1})
	assert.Pass('a', should.BeIn, map[rune]int{'a': 1})
}

func TestShouldNotBeIn(t *testing.T) {
	assert := NewAssertion(t)
	assert.ExpectedCountInvalid("actual", should.NOT.BeIn)
	assert.ExpectedCountInvalid("actual", should.NOT.BeIn, "EXPECTED", "EXTRA")
	assert.KindMismatch(false, should.NOT.BeIn, "string")
	assert.KindMismatch("hi", should.NOT.BeIn, 1)

	// strings:
	assert.Pass("no", should.NOT.BeIn, "yes")
	assert.Fail("rat", should.NOT.BeIn, "integrate")
	assert.Fail('b', should.NOT.BeIn, "abc")

	// slices:
	assert.Pass('d', should.NOT.BeIn, []byte("abc"))
	assert.Fail('b', should.NOT.BeIn, []byte("abc"))
	assert.Fail(98, should.NOT.BeIn, []byte("abc"))

	// arrays:
	assert.Pass('d', should.NOT.BeIn, [3]byte{'a', 'b', 'c'})
	assert.Fail('b', should.NOT.BeIn, [3]byte{'a', 'b', 'c'})
	assert.Fail(98, should.NOT.BeIn, [3]byte{'a', 'b', 'c'})

	// maps:
	assert.Pass('b', should.NOT.BeIn, map[rune]int{'a': 1})
	assert.Fail('a', should.NOT.BeIn, map[rune]int{'a': 1})
}
