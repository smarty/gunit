package should

import "errors"

// BeGreaterThan verifies that actual is greater than expected.
// Both actual and expected must be strings or numeric in type.
func BeGreaterThan(actual any, EXPECTED ...any) error {
	lessThanErr := BeLessThan(actual, EXPECTED...)
	if errors.Is(lessThanErr, ErrTypeMismatch) || errors.Is(lessThanErr, ErrExpectedCountInvalid) {
		return lessThanErr
	}
	equalityErr := Equal(actual, EXPECTED...)
	if lessThanErr == nil || equalityErr == nil {
		return failure("%v was not greater than %v", actual, EXPECTED[0])
	}
	return nil
}

// BeGreaterThan negated!
func (negated) BeGreaterThan(actual any, expected ...any) error {
	err := BeGreaterThan(actual, expected...)
	if errors.Is(err, ErrAssertionFailure) {
		return nil
	}

	if err != nil {
		return err
	}

	return failure("\n"+
		"  expected:               %#v\n"+
		"  to not be greater than: %#v\n"+
		"  (but it was)",
		expected[0],
		actual,
	)
}
