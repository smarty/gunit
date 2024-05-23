package should

import "errors"

func So(t testingT, actual any, assertion Assertion, expected ...any) {
	t.Helper()
	err := assertion(actual, expected...)
	if errors.Is(err, ErrFatalAssertionFailure) {
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
