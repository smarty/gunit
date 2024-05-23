package must_test

import (
	"testing"

	"github.com/smarty/gunit/v2/must"
	"github.com/smarty/gunit/v2/should"
)

func TestWrapFatalSuccess(t *testing.T) {
	err := must.Equal(1, 1)
	should.So(t, err, should.BeNil)
}
func TestWrapFatalFailure(t *testing.T) {
	err := must.Equal(1, 2)
	should.So(t, err, should.WrapError, should.ErrFatalAssertionFailure)
	should.So(t, err, should.WrapError, should.ErrAssertionFailure)
}
func TestWrapFatalSuccess_NOT(t *testing.T) {
	err := must.NOT.Equal(1, 2)
	should.So(t, err, should.BeNil)
}
func TestWrapFatalFailure_NOT(t *testing.T) {
	err := must.NOT.Equal(1, 1)
	should.So(t, err, should.WrapError, should.ErrFatalAssertionFailure)
	should.So(t, err, should.WrapError, should.ErrAssertionFailure)
}
