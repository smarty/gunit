package should_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/smarty/gunit/v2/assert/should"
)

func TestShouldWrapError(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.WrapError)
	assert.ExpectedCountInvalid("actual", should.WrapError, "EXPECTED", "EXTRA")

	assert.TypeMismatch(inner, should.WrapError, 42)
	assert.TypeMismatch(42, should.WrapError, inner)

	assert.Pass(outer, should.WrapError, inner)
	assert.Fail(inner, should.WrapError, outer)
}

var (
	inner = errors.New("inner")
	outer = fmt.Errorf("outer(%w)", inner)
)
