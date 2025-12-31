package should

import "errors"

// BeLessThanOrEqualTo verifies that actual is less than or equal to expected.
// Both actual and expected must be strings or numeric in type.
func BeLessThanOrEqualTo(actual any, expected ...any) error {
	err := Equal(actual, expected...)
	if err == nil {
		return nil
	}
	err = BeLessThan(actual, expected...)
	if errors.Is(err, ErrAssertionFailure) {
		return failure("%v was not less than or equal to %v", actual, expected)
	}

	if err != nil {
		return err
	}
	return nil
}

// BeLessThanOrEqualTo negated!
func (negated) BeLessThanOrEqualTo(actual any, expected ...any) error {
	err := BeLessThanOrEqualTo(actual, expected...)
	if errors.Is(err, ErrAssertionFailure) {
		return nil
	}

	if err != nil {
		return err
	}

	return failure("\n"+
		"  expected:                        %#v\n"+
		"  to not be less than or equal to: %#v\n"+
		"  (but it was)",
		expected[0],
		actual,
	)
}
