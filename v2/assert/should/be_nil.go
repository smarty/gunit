package should

import (
	"errors"
	"reflect"
)

// BeNil verifies that actual is the nil value.
func BeNil(actual any, expected ...any) error {
	err := validateExpected(0, expected)
	if err != nil {
		return err
	}

	if actual == nil || interfaceHasNilValue(actual) {
		return nil
	}

	return failure("got %#v, want <nil>", actual)
}
func interfaceHasNilValue(actual any) bool {
	value := reflect.ValueOf(actual)
	kind := value.Kind()
	nillable := kind == reflect.Slice ||
		kind == reflect.Chan ||
		kind == reflect.Func ||
		kind == reflect.Ptr ||
		kind == reflect.Map

	// Careful: reflect.Value.IsNil() will panic unless it's
	// an interface, chan, map, func, slice, or ptr
	// Reference: http://golang.org/pkg/reflect/#Value.IsNil
	return nillable && value.IsNil()
}

// BeNil negated!
func (negated) BeNil(actual any, expected ...any) error {
	err := BeNil(actual, expected...)
	if errors.Is(err, ErrAssertionFailure) {
		return nil
	}

	if err != nil {
		return err
	}

	return failure("got nil, want non-<nil>")
}
