package should_test

import (
	"math"
	"testing"

	"github.com/smarty/gunit/v2/assert/should"
)

func TestShouldBeGreaterThan(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual-but-missing-expected", should.BeGreaterThan)
	assert.ExpectedCountInvalid("actual", should.BeGreaterThan, "expected", "required")
	assert.TypeMismatch(true, should.BeGreaterThan, 1)
	assert.TypeMismatch(1, should.BeGreaterThan, true)

	assert.Fail("a", should.BeGreaterThan, "b") // both strings
	assert.Pass("b", should.BeGreaterThan, "a")

	assert.Fail(1, should.BeGreaterThan, 1) // both ints
	assert.Pass(2, should.BeGreaterThan, 1)

	assert.Pass(float32(2.0), should.BeGreaterThan, float64(1)) // both floats
	assert.Fail(1.0, should.BeGreaterThan, 2.0)

	assert.Pass(int32(2), should.BeGreaterThan, int64(1)) // both signed
	assert.Fail(int32(1), should.BeGreaterThan, int64(2))

	assert.Pass(uint32(2), should.BeGreaterThan, uint64(1)) // both unsigned
	assert.Fail(uint32(1), should.BeGreaterThan, uint64(2))

	assert.Pass(int32(2), should.BeGreaterThan, uint32(1)) // signed and unsigned
	assert.Fail(int32(1), should.BeGreaterThan, uint32(2))
	// if actual < 0: false
	// (because by definition the expected value, an unsigned value must be >= 0)
	const reallyBig uint64 = math.MaxUint64
	assert.Fail(-1, should.BeGreaterThan, reallyBig)

	assert.Pass(uint32(2), should.BeGreaterThan, int32(1)) // unsigned and signed
	assert.Fail(uint32(1), should.BeGreaterThan, int32(2))
	// if actual > math.MaxInt64: true
	// (because by definition the expected value, a signed value must be > math.MaxInt64)
	const tooBig uint64 = math.MaxInt64 + 1
	assert.Pass(tooBig, should.BeGreaterThan, 42)

	assert.Pass(2.0, should.BeGreaterThan, 1) // float and integer
	assert.Fail(1.0, should.BeGreaterThan, 2)

	assert.Pass(2, should.BeGreaterThan, 1.0) // integer and float
	assert.Fail(1, should.BeGreaterThan, 2.0)
}

func TestShouldNOTBeGreaterThan(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual-but-missing-expected", should.NOT.BeGreaterThan)
	assert.ExpectedCountInvalid("actual", should.NOT.BeGreaterThan, "expected", "required")
	assert.TypeMismatch(true, should.NOT.BeGreaterThan, 1)
	assert.TypeMismatch(1, should.NOT.BeGreaterThan, true)

	assert.Pass("a", should.NOT.BeGreaterThan, "b") // both strings
	assert.Fail("b", should.NOT.BeGreaterThan, "a")

	assert.Pass(1, should.NOT.BeGreaterThan, 1) // both ints
	assert.Fail(2, should.NOT.BeGreaterThan, 1)

	assert.Fail(float32(2.0), should.NOT.BeGreaterThan, float64(1)) // both floats
	assert.Pass(1.0, should.NOT.BeGreaterThan, 2.0)

	assert.Fail(int32(2), should.NOT.BeGreaterThan, int64(1)) // both signed
	assert.Pass(int32(1), should.NOT.BeGreaterThan, int64(2))

	assert.Fail(uint32(2), should.NOT.BeGreaterThan, uint64(1)) // both unsigned
	assert.Pass(uint32(1), should.NOT.BeGreaterThan, uint64(2))

	assert.Fail(int32(2), should.NOT.BeGreaterThan, uint32(1)) // signed and unsigned
	assert.Pass(int32(1), should.NOT.BeGreaterThan, uint32(2))
	// if actual < 0: true
	// (because by definition the expected value, an unsigned value must be >= 0)
	const reallyBig uint64 = math.MaxUint64
	assert.Pass(-1, should.NOT.BeGreaterThan, reallyBig)

	assert.Fail(uint32(2), should.NOT.BeGreaterThan, int32(1)) // unsigned and signed
	assert.Pass(uint32(1), should.NOT.BeGreaterThan, int32(2))
	// if actual > math.MaxInt64: false
	// (because by definition the expected value, a signed value can't be > math.MaxInt64)
	const tooBig uint64 = math.MaxInt64 + 1
	assert.Fail(tooBig, should.NOT.BeGreaterThan, 42)

	assert.Fail(2.0, should.NOT.BeGreaterThan, 1) // float and integer
	assert.Pass(1.0, should.NOT.BeGreaterThan, 2)

	assert.Fail(2, should.NOT.BeGreaterThan, 1.0) // integer and float
	assert.Pass(1, should.NOT.BeGreaterThan, 2.0)
}
