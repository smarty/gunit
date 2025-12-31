package should

import (
	"errors"
	"fmt"
	"reflect"
)

// WrapError uses errors.Is to verify that actual is an error value
// that wraps expected[0] (also an error value).
func WrapError(actual any, expected ...any) error {
	err := validateExpected(1, expected)
	if err != nil {
		return err
	}

	inner, ok := expected[0].(error)
	if !ok {
		return errTypeMismatch(expected[0])
	}

	outer, ok := actual.(error)
	if !ok {
		return errTypeMismatch(actual)
	}

	if errors.Is(outer, inner) {
		return nil
	}

	return fmt.Errorf("%w:\n"+
		"\t            outer err: (%s)\n"+
		"\tshould wrap inner err: (%s)",
		ErrAssertionFailure,
		outer,
		inner,
	)
}

func errTypeMismatch(v any) error {
	return fmt.Errorf(
		"%w: got %s, want error",
		ErrTypeMismatch,
		reflect.TypeOf(v),
	)
}
