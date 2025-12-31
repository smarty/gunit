package should_test

import (
	"testing"

	"github.com/smarty/gunit/v2/assert/should"
)

func TestShouldPanic(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.Panic, "EXPECTED", "EXTRA")
	assert.TypeMismatch("wrong type", should.Panic)

	assert.Fail(func() {}, should.Panic)
	assert.Pass(func() { panic("yay") }, should.Panic)
	assert.Pass(func() { panic(nil) }, should.Panic) // tricky!
}

func TestShouldNotPanic(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.NOT.Panic, "EXPECTED", "EXTRA")
	assert.TypeMismatch("wrong type", should.NOT.Panic)

	assert.Fail(func() { panic("boo") }, should.NOT.Panic)
	assert.Pass(func() {}, should.NOT.Panic)
}
