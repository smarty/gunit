package should_test

import (
	"math"
	"testing"

	"github.com/smarty/gunit/v2/assert/should"
)

func TestShouldBeGreaterThanOrEqualTo(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual-but-missing-expected", should.BeGreaterThanOrEqualTo)
	assert.ExpectedCountInvalid("actual", should.BeGreaterThanOrEqualTo, "expected", "required")
	assert.TypeMismatch(true, should.BeGreaterThanOrEqualTo, 1)
	assert.TypeMismatch(1, should.BeGreaterThanOrEqualTo, true)

	assert.Fail("a", should.BeGreaterThanOrEqualTo, "b") // both strings
	assert.Pass("b", should.BeGreaterThanOrEqualTo, "a")

	assert.Pass(2, should.BeGreaterThanOrEqualTo, 1) // both ints
	assert.Pass(1, should.BeGreaterThanOrEqualTo, 1)
	assert.Fail(1, should.BeGreaterThanOrEqualTo, 2)

	assert.Pass(float32(2.0), should.BeGreaterThanOrEqualTo, float64(1)) // both floats
	assert.Pass(float32(2.0), should.BeGreaterThanOrEqualTo, float64(2))
	assert.Fail(1.0, should.BeGreaterThanOrEqualTo, 2.0)

	assert.Pass(int32(2), should.BeGreaterThanOrEqualTo, int64(1)) // both signed
	assert.Pass(int32(2), should.BeGreaterThanOrEqualTo, int64(2))
	assert.Fail(int32(1), should.BeGreaterThanOrEqualTo, int64(2))

	assert.Pass(uint32(2), should.BeGreaterThanOrEqualTo, uint64(1)) // both unsigned
	assert.Pass(uint32(2), should.BeGreaterThanOrEqualTo, uint64(2))
	assert.Fail(uint32(1), should.BeGreaterThanOrEqualTo, uint64(2))

	assert.Pass(int32(2), should.BeGreaterThanOrEqualTo, uint32(1)) // signed and unsigned
	assert.Pass(int32(2), should.BeGreaterThanOrEqualTo, uint32(2))
	assert.Fail(int32(1), should.BeGreaterThanOrEqualTo, uint32(2))
	// if actual < 0: false
	// (because by definition the expected value, an unsigned value must be >= 0)
	const reallyBig uint64 = math.MaxUint64
	assert.Fail(-1, should.BeGreaterThanOrEqualTo, reallyBig)

	assert.Pass(uint32(2), should.BeGreaterThanOrEqualTo, int32(1)) // unsigned and signed
	assert.Pass(uint32(2), should.BeGreaterThanOrEqualTo, int32(2))
	assert.Fail(uint32(1), should.BeGreaterThanOrEqualTo, int32(2))
	// if actual > math.MaxInt64: true
	// (because by definition the expected value, a signed value must be > math.MaxInt64)
	const tooBig uint64 = math.MaxInt64 + 1
	assert.Pass(tooBig, should.BeGreaterThanOrEqualTo, 42)

	assert.Pass(2.0, should.BeGreaterThanOrEqualTo, 1) // float and integer
	assert.Pass(2.0, should.BeGreaterThanOrEqualTo, 2)
	assert.Fail(1.0, should.BeGreaterThanOrEqualTo, 2)

	assert.Pass(2, should.BeGreaterThanOrEqualTo, 1.0) // integer and float
	assert.Pass(2, should.BeGreaterThanOrEqualTo, 2.0)
	assert.Fail(1, should.BeGreaterThanOrEqualTo, 2.0)
}

func TestShouldNOTBeGreaterThanOrEqualTo(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual-but-missing-expected", should.NOT.BeGreaterThanOrEqualTo)
	assert.ExpectedCountInvalid("actual", should.NOT.BeGreaterThanOrEqualTo, "expected", "required")
	assert.TypeMismatch(true, should.NOT.BeGreaterThanOrEqualTo, 1)
	assert.TypeMismatch(1, should.NOT.BeGreaterThanOrEqualTo, true)

	assert.Pass("a", should.NOT.BeGreaterThanOrEqualTo, "b") // both strings
	assert.Fail("a", should.NOT.BeGreaterThanOrEqualTo, "a")
	assert.Fail("b", should.NOT.BeGreaterThanOrEqualTo, "a")

	assert.Pass(1, should.NOT.BeGreaterThanOrEqualTo, 2) // both ints
	assert.Fail(1, should.NOT.BeGreaterThanOrEqualTo, 1)
	assert.Fail(2, should.NOT.BeGreaterThanOrEqualTo, 1)

	assert.Fail(float32(2.0), should.NOT.BeGreaterThanOrEqualTo, float64(1)) // both floats
	assert.Fail(float32(2.0), should.NOT.BeGreaterThanOrEqualTo, float64(2))
	assert.Pass(1.0, should.NOT.BeGreaterThanOrEqualTo, 2.0)

	assert.Fail(int32(2), should.NOT.BeGreaterThanOrEqualTo, int64(1)) // both signed
	assert.Fail(int32(2), should.NOT.BeGreaterThanOrEqualTo, int64(2))
	assert.Pass(int32(1), should.NOT.BeGreaterThanOrEqualTo, int64(2))

	assert.Fail(uint32(2), should.NOT.BeGreaterThanOrEqualTo, uint64(1)) // both unsigned
	assert.Fail(uint32(2), should.NOT.BeGreaterThanOrEqualTo, uint64(2))
	assert.Pass(uint32(1), should.NOT.BeGreaterThanOrEqualTo, uint64(2))

	assert.Fail(int32(2), should.NOT.BeGreaterThanOrEqualTo, uint32(1)) // signed and unsigned
	assert.Fail(int32(2), should.NOT.BeGreaterThanOrEqualTo, uint32(2))
	assert.Pass(int32(1), should.NOT.BeGreaterThanOrEqualTo, uint32(2))
	// if actual < 0: true
	// (because by definition the expected value, an unsigned value must be >= 0)
	const reallyBig uint64 = math.MaxUint64
	assert.Pass(-1, should.NOT.BeGreaterThanOrEqualTo, reallyBig)

	assert.Fail(uint32(2), should.NOT.BeGreaterThanOrEqualTo, int32(1)) // unsigned and signed
	assert.Fail(uint32(2), should.NOT.BeGreaterThanOrEqualTo, int32(2))
	assert.Pass(uint32(1), should.NOT.BeGreaterThanOrEqualTo, int32(2))
	// if actual > math.MaxInt64: false
	// (because by definition the expected value, a signed value can't be > math.MaxInt64)
	const tooBig uint64 = math.MaxInt64 + 1
	assert.Fail(tooBig, should.NOT.BeGreaterThanOrEqualTo, 42)

	assert.Fail(2.0, should.NOT.BeGreaterThanOrEqualTo, 1) // float and integer
	assert.Fail(2.0, should.NOT.BeGreaterThanOrEqualTo, 2)
	assert.Pass(1.0, should.NOT.BeGreaterThanOrEqualTo, 2)

	assert.Fail(2, should.NOT.BeGreaterThanOrEqualTo, 1.0) // integer and float
	assert.Fail(2, should.NOT.BeGreaterThanOrEqualTo, 2.0)
	assert.Pass(1, should.NOT.BeGreaterThanOrEqualTo, 2.0)
}
