package should

import "errors"

// BeGreaterThanOrEqualTo verifies that actual is less than or equal to expected.
// Both actual and expected must be strings or numeric in type.
func BeGreaterThanOrEqualTo(actual any, expected ...any) error {
	err := Equal(actual, expected...)
	if err == nil {
		return nil
	}
	err = BeGreaterThan(actual, expected...)
	if errors.Is(err, ErrAssertionFailure) {
		return failure("%v was not greater than or equal to %v", actual, expected)
	}

	if err != nil {
		return err
	}
	return nil
}

// BeGreaterThanOrEqualTo negated!
func (negated) BeGreaterThanOrEqualTo(actual any, expected ...any) error {
	err := BeGreaterThanOrEqualTo(actual, expected...)
	if errors.Is(err, ErrAssertionFailure) {
		return nil
	}

	if err != nil {
		return err
	}

	return failure("\n"+
		"  expected:                           %#v\n"+
		"  to not be greater than or equal to: %#v\n"+
		"  (but it was)",
		expected[0],
		actual,
	)
}
