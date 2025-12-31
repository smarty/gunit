package should_test

import (
	"math"
	"testing"

	"github.com/smarty/gunit/v2/assert/should"
)

func TestShouldBeLessThan(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual-but-missing-expected", should.BeLessThan)
	assert.ExpectedCountInvalid("actual", should.BeLessThan, "expected", "required")
	assert.TypeMismatch(true, should.BeLessThan, 1)
	assert.TypeMismatch(1, should.BeLessThan, true)

	assert.Fail("b", should.BeLessThan, "a") // both strings
	assert.Pass("a", should.BeLessThan, "b")

	assert.Fail(1, should.BeLessThan, 1) // both ints
	assert.Pass(1, should.BeLessThan, 2)

	assert.Pass(float32(1.0), should.BeLessThan, float64(2)) // both floats
	assert.Fail(2.0, should.BeLessThan, 1.0)

	assert.Pass(int32(1), should.BeLessThan, int64(2)) // both signed
	assert.Fail(int32(2), should.BeLessThan, int64(1))

	assert.Pass(uint32(1), should.BeLessThan, uint64(2)) // both unsigned
	assert.Fail(uint32(2), should.BeLessThan, uint64(1))

	assert.Pass(int32(1), should.BeLessThan, uint32(2)) // signed and unsigned
	assert.Fail(int32(2), should.BeLessThan, uint32(1))
	// if actual < 0: true
	// (because by definition the expected value, an unsigned value can't be < 0)
	const reallyBig uint64 = math.MaxUint64
	assert.Pass(-1, should.BeLessThan, reallyBig)

	assert.Pass(uint32(1), should.BeLessThan, int32(2)) // unsigned and signed
	assert.Fail(uint32(2), should.BeLessThan, int32(1))
	// if actual > math.MaxInt64: false
	// (because by definition the expected value, a signed value can't be > math.MaxInt64)
	const tooBig uint64 = math.MaxInt64 + 1
	assert.Fail(tooBig, should.BeLessThan, 42)

	assert.Pass(1.0, should.BeLessThan, 2) // float and integer
	assert.Fail(2.0, should.BeLessThan, 1)

	assert.Pass(1.0, should.BeLessThan, uint(2)) // float and unsigned integer
	assert.Fail(2.0, should.BeLessThan, uint(1))

	assert.Pass(1, should.BeLessThan, 2.0) // integer and float
	assert.Fail(2, should.BeLessThan, 1.0)

	assert.Pass(uint(1), should.BeLessThan, 2.0) // unsigned integer and float
	assert.Fail(uint(2), should.BeLessThan, 1.0)
}

func TestShouldNOTBeLessThan(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual-but-missing-expected", should.NOT.BeLessThan)
	assert.ExpectedCountInvalid("actual", should.NOT.BeLessThan, "expected", "required")
	assert.TypeMismatch(true, should.NOT.BeLessThan, 1)
	assert.TypeMismatch(1, should.NOT.BeLessThan, true)

	assert.Pass("b", should.NOT.BeLessThan, "a") // both strings
	assert.Fail("a", should.NOT.BeLessThan, "b")

	assert.Pass(1, should.NOT.BeLessThan, 1) // both ints
	assert.Fail(1, should.NOT.BeLessThan, 2)

	assert.Fail(float32(1.0), should.NOT.BeLessThan, float64(2)) // both floats
	assert.Pass(2.0, should.NOT.BeLessThan, 1.0)

	assert.Fail(int32(1), should.NOT.BeLessThan, int64(2)) // both signed
	assert.Pass(int32(2), should.NOT.BeLessThan, int64(1))

	assert.Fail(uint32(1), should.NOT.BeLessThan, uint64(2)) // both unsigned
	assert.Pass(uint32(2), should.NOT.BeLessThan, uint64(1))

	assert.Fail(int32(1), should.NOT.BeLessThan, uint32(2)) // signed and unsigned
	assert.Pass(int32(2), should.NOT.BeLessThan, uint32(1))
	// if actual < 0: false
	// (because by definition the expected value, an unsigned value can't be < 0)
	const reallyBig uint64 = math.MaxUint64
	assert.Fail(-1, should.NOT.BeLessThan, reallyBig)

	assert.Fail(uint32(1), should.NOT.BeLessThan, int32(2)) // unsigned and signed
	assert.Pass(uint32(2), should.NOT.BeLessThan, int32(1))
	// if actual > math.MaxInt64: true
	// (because by definition the expected value, a signed value can't be > math.MaxInt64)
	const tooBig uint64 = math.MaxInt64 + 1
	assert.Pass(tooBig, should.NOT.BeLessThan, 42)

	assert.Fail(1.0, should.NOT.BeLessThan, 2) // float and integer
	assert.Pass(2.0, should.NOT.BeLessThan, 1)

	assert.Fail(1, should.NOT.BeLessThan, 2.0) // integer and float
	assert.Pass(2, should.NOT.BeLessThan, 1.0)
}
