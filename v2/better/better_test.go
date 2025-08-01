package better_test

import (
	"testing"

	"github.com/smarty/gunit/v2"
	"github.com/smarty/gunit/v2/better"
	"github.com/smarty/gunit/v2/should"
)

func TestWrapFatalSuccess(t *testing.T) {
	err := better.Equal(1, 1)
	gunit.So(t, err, should.BeNil)
}
func TestWrapFatalFailure(t *testing.T) {
	err := better.Equal(1, 2)
	gunit.So(t, err, should.WrapError, should.ErrFatalAssertionFailure)
	gunit.So(t, err, should.WrapError, should.ErrAssertionFailure)
}
func TestWrapFatalSuccess_NOT(t *testing.T) {
	err := better.NOT.Equal(1, 2)
	gunit.So(t, err, should.BeNil)
}
func TestWrapFatalFailure_NOT(t *testing.T) {
	err := better.NOT.Equal(1, 1)
	gunit.So(t, err, should.WrapError, should.ErrFatalAssertionFailure)
	gunit.So(t, err, should.WrapError, should.ErrAssertionFailure)
}
