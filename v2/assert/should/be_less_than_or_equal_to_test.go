package should_test

import (
	"math"
	"testing"

	"github.com/smarty/gunit/v2/assert/should"
)

func TestShouldBeLessThanOrEqualTo(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual-but-missing-expected", should.BeLessThanOrEqualTo)
	assert.ExpectedCountInvalid("actual", should.BeLessThanOrEqualTo, "expected", "required")
	assert.TypeMismatch(true, should.BeLessThanOrEqualTo, 1)
	assert.TypeMismatch(1, should.BeLessThanOrEqualTo, true)

	assert.Fail("b", should.BeLessThanOrEqualTo, "a") // both strings
	assert.Pass("a", should.BeLessThanOrEqualTo, "b")
	assert.Pass("a", should.BeLessThanOrEqualTo, "a")

	assert.Fail(2, should.BeLessThanOrEqualTo, 1) // both ints
	assert.Pass(1, should.BeLessThanOrEqualTo, 2)
	assert.Pass(1, should.BeLessThanOrEqualTo, 1)

	assert.Pass(float32(1.0), should.BeLessThanOrEqualTo, float64(2)) // both floats
	assert.Fail(2.0, should.BeLessThanOrEqualTo, 1.0)
	assert.Pass(2.0, should.BeLessThanOrEqualTo, 2.0)

	assert.Pass(int32(1), should.BeLessThanOrEqualTo, int64(2)) // both signed
	assert.Fail(int32(2), should.BeLessThanOrEqualTo, int64(1))
	assert.Pass(int32(2), should.BeLessThanOrEqualTo, int64(2))

	assert.Pass(uint32(1), should.BeLessThanOrEqualTo, uint64(2)) // both unsigned
	assert.Fail(uint32(2), should.BeLessThanOrEqualTo, uint64(1))
	assert.Pass(uint32(2), should.BeLessThanOrEqualTo, uint64(2))

	assert.Pass(int32(1), should.BeLessThanOrEqualTo, uint32(2)) // signed and unsigned
	assert.Fail(int32(2), should.BeLessThanOrEqualTo, uint32(1))
	assert.Pass(int32(2), should.BeLessThanOrEqualTo, uint32(2))
	// if actual < 0: true
	// (because by definition the expected value, an unsigned value can't be < 0)
	const reallyBig uint64 = math.MaxUint64
	assert.Pass(-1, should.BeLessThanOrEqualTo, reallyBig)

	assert.Pass(uint32(1), should.BeLessThanOrEqualTo, int32(2)) // unsigned and signed
	assert.Fail(uint32(2), should.BeLessThanOrEqualTo, int32(1))
	assert.Pass(uint32(2), should.BeLessThanOrEqualTo, int32(2))
	// if actual > math.MaxInt64: false
	// (because by definition the expected value, a signed value can't be > math.MaxInt64)
	const tooBig uint64 = math.MaxInt64 + 1
	assert.Fail(tooBig, should.BeLessThanOrEqualTo, 42)

	assert.Pass(1.0, should.BeLessThanOrEqualTo, 2) // float and integer
	assert.Fail(2.0, should.BeLessThanOrEqualTo, 1)
	assert.Pass(2.0, should.BeLessThanOrEqualTo, 2)

	assert.Pass(1, should.BeLessThanOrEqualTo, 2.0) // integer and float
	assert.Fail(2, should.BeLessThanOrEqualTo, 1.0)
	assert.Pass(2, should.BeLessThanOrEqualTo, 2.0)

}

func TestShouldNotBeLessThanOrEqualTo(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual-but-missing-expected", should.NOT.BeLessThanOrEqualTo)
	assert.ExpectedCountInvalid("actual", should.NOT.BeLessThanOrEqualTo, "expected", "required")
	assert.TypeMismatch(true, should.NOT.BeLessThanOrEqualTo, 1)
	assert.TypeMismatch(1, should.NOT.BeLessThanOrEqualTo, true)

	assert.Pass("b", should.NOT.BeLessThanOrEqualTo, "a") // both strings
	assert.Fail("a", should.NOT.BeLessThanOrEqualTo, "b")
	assert.Fail("a", should.NOT.BeLessThanOrEqualTo, "a")

	assert.Pass(2, should.NOT.BeLessThanOrEqualTo, 1) // both ints
	assert.Fail(1, should.NOT.BeLessThanOrEqualTo, 2)
	assert.Fail(1, should.NOT.BeLessThanOrEqualTo, 1)

	assert.Fail(float32(1.0), should.NOT.BeLessThanOrEqualTo, float64(2)) // both floats
	assert.Pass(2.0, should.NOT.BeLessThanOrEqualTo, 1.0)
	assert.Fail(2.0, should.NOT.BeLessThanOrEqualTo, 2.0)

	assert.Fail(int32(1), should.NOT.BeLessThanOrEqualTo, int64(2)) // both signed
	assert.Pass(int32(2), should.NOT.BeLessThanOrEqualTo, int64(1))
	assert.Fail(int32(2), should.NOT.BeLessThanOrEqualTo, int64(2))

	assert.Fail(uint32(1), should.NOT.BeLessThanOrEqualTo, uint64(2)) // both unsigned
	assert.Pass(uint32(2), should.NOT.BeLessThanOrEqualTo, uint64(1))
	assert.Fail(uint32(2), should.NOT.BeLessThanOrEqualTo, uint64(2))

	assert.Fail(int32(1), should.NOT.BeLessThanOrEqualTo, uint32(2)) // signed and unsigned
	assert.Pass(int32(2), should.NOT.BeLessThanOrEqualTo, uint32(1))
	assert.Fail(int32(2), should.NOT.BeLessThanOrEqualTo, uint32(2))
	// if actual < 0: true
	// (because by definition the expected value, an unsigned value can't be < 0)
	const reallyBig uint64 = math.MaxUint64
	assert.Fail(-1, should.NOT.BeLessThanOrEqualTo, reallyBig)

	assert.Fail(uint32(1), should.NOT.BeLessThanOrEqualTo, int32(2)) // unsigned and signed
	assert.Pass(uint32(2), should.NOT.BeLessThanOrEqualTo, int32(1))
	assert.Fail(uint32(2), should.NOT.BeLessThanOrEqualTo, int32(2))
	// if actual > math.MaxInt64: false
	// (because by definition the expected value, a signed value can't be > math.MaxInt64)
	const tooBig uint64 = math.MaxInt64 + 1
	assert.Pass(tooBig, should.NOT.BeLessThanOrEqualTo, 42)

	assert.Fail(1.0, should.NOT.BeLessThanOrEqualTo, 2) // float and integer
	assert.Pass(2.0, should.NOT.BeLessThanOrEqualTo, 1)
	assert.Fail(2.0, should.NOT.BeLessThanOrEqualTo, 2)

	assert.Fail(1, should.NOT.BeLessThanOrEqualTo, 2.0) // integer and float
	assert.Pass(2, should.NOT.BeLessThanOrEqualTo, 1.0)
	assert.Fail(2, should.NOT.BeLessThanOrEqualTo, 2.0)

}
