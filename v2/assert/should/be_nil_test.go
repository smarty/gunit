package should_test

import (
	"errors"
	"testing"

	"github.com/smarty/gunit/v2/assert/should"
)

func TestShouldBeNil(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.BeNil, "EXTRA")

	assert.Pass(nil, should.BeNil)
	assert.Pass([]string(nil), should.BeNil)
	assert.Pass((*string)(nil), should.BeNil)
	assert.Fail(notNil, should.BeNil)
}

func TestShouldNotBeNil(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.NOT.BeNil, "EXTRA")

	assert.Fail(nil, should.NOT.BeNil)
	assert.Fail([]string(nil), should.NOT.BeNil)
	assert.Pass(notNil, should.NOT.BeNil)
}

var notNil = errors.New("not nil")
