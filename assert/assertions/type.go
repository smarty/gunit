package assertions

import (
	"errors"
	"fmt"
	"reflect"
)

// ShouldHaveSameTypeAs receives exactly two parameters and compares their underlying types for equality.
func ShouldHaveSameTypeAs(actual any, expected ...any) string {
	if fail := need(1, expected); fail != success {
		return fail
	}

	first := reflect.TypeOf(actual)
	second := reflect.TypeOf(expected[0])

	if first != second {
		return fmt.Sprintf(shouldHaveBeenA, actual, second, first)
	}

	return success
}

// ShouldNotHaveSameTypeAs receives exactly two parameters and compares their underlying types for inequality.
func ShouldNotHaveSameTypeAs(actual any, expected ...any) string {
	if fail := need(1, expected); fail != success {
		return fail
	}

	first := reflect.TypeOf(actual)
	second := reflect.TypeOf(expected[0])

	if (actual == nil && expected[0] == nil) || first == second {
		return fmt.Sprintf(shouldNotHaveBeenA, actual, second)
	}
	return success
}

// ShouldWrap asserts that the first argument (which must be an error value)
// 'wraps' the second/final argument (which must also be an error value).
// It relies on errors.Is to make the determination (https://golang.org/pkg/errors/#Is).
func ShouldWrap(actual any, expected ...any) string {
	if fail := need(1, expected); fail != success {
		return fail
	}

	if !isError(actual) || !isError(expected[0]) {
		return fmt.Sprintf(shouldWrapInvalidTypes, reflect.TypeOf(actual), reflect.TypeOf(expected[0]))
	}

	if !errors.Is(actual.(error), expected[0].(error)) {
		return fmt.Sprintf(`Expected error("%s") to wrap error("%s") but it didn't.`, actual, expected[0])
	}

	return success
}

func isError(value any) bool { _, ok := value.(error); return ok }
