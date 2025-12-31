package gunit

import (
	"errors"

	"github.com/smarty/gunit/v2/should"
)

func So(t testingT, actual any, assertion Assertion, expected ...any) {
	t.Helper()
	err := assertion(actual, expected...)
	if errors.Is(err, should.ErrFatalAssertionFailure) {
		t.Fatal(err)
	}
	if err != nil {
		t.Error(err)
	}
}

type testingT interface {
	Helper()
	Error(...any)
	Fatal(...any)
}
type Assertion func(actual any, expected ...any) error
