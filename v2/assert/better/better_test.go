package better_test

import (
	"testing"

	"github.com/smarty/gunit/v2/assert"
	"github.com/smarty/gunit/v2/assert/better"
	"github.com/smarty/gunit/v2/assert/should"
)

func TestWrapFatalSuccess(t *testing.T) {
	err := better.Equal(1, 1)
	assert.So(t, err, should.BeNil)
}
func TestWrapFatalFailure(t *testing.T) {
	err := better.Equal(1, 2)
	assert.So(t, err, should.WrapError, should.ErrFatalAssertionFailure)
	assert.So(t, err, should.WrapError, should.ErrAssertionFailure)
}
func TestWrapFatalSuccess_NOT(t *testing.T) {
	err := better.NOT.Equal(1, 2)
	assert.So(t, err, should.BeNil)
}
func TestWrapFatalFailure_NOT(t *testing.T) {
	err := better.NOT.Equal(1, 1)
	assert.So(t, err, should.WrapError, should.ErrFatalAssertionFailure)
	assert.So(t, err, should.WrapError, should.ErrAssertionFailure)
}
