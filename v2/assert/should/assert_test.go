package should_test

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/smarty/gunit/v2/assert"
	"github.com/smarty/gunit/v2/assert/should"
)

type Assertion struct{ *testing.T }

func NewAssertion(t *testing.T) *Assertion {
	return &Assertion{T: t}
}
func (this *Assertion) ExpectedCountInvalid(actual any, assertion assert.Assertion, expected ...any) {
	this.Helper()
	this.err(actual, assertion, expected, should.ErrExpectedCountInvalid)
}
func (this *Assertion) TypeMismatch(actual any, assertion assert.Assertion, expected ...any) {
	this.Helper()
	this.err(actual, assertion, expected, should.ErrTypeMismatch)
}
func (this *Assertion) KindMismatch(actual any, assertion assert.Assertion, expected ...any) {
	this.Helper()
	this.err(actual, assertion, expected, should.ErrKindMismatch)
}
func (this *Assertion) Fail(actual any, assertion assert.Assertion, expected ...any) {
	this.Helper()
	this.err(actual, assertion, expected, should.ErrAssertionFailure)
}
func (this *Assertion) Pass(actual any, assertion assert.Assertion, expected ...any) {
	this.Helper()
	this.err(actual, assertion, expected, nil)
}
func (this *Assertion) err(actual any, assertion assert.Assertion, expected []any, expectedErr error) {
	this.Helper()
	_, file, line, _ := runtime.Caller(2)
	subTest := fmt.Sprintf("%s:%d", filepath.Base(file), line)
	this.Run(subTest, func(t *testing.T) {
		t.Helper()
		err := assertion(actual, expected...)
		if !errors.Is(err, expectedErr) {
			t.Errorf("[FAIL]\n"+
				"expected: %v\n"+
				"actual:   %v",
				expected,
				actual,
			)
		} else if testing.Verbose() {
			t.Log(
				"\n", err, "\n",
				"(above error report printed for visual inspection)",
			)
		}
	})
}
