package should

import "time"

// HappenAfter ensures that the first time value happens after the second.
func HappenAfter(actual any, expected ...any) error {
	err := validateExpected(1, expected)
	if err != nil {
		return err
	}
	err = validateType(actual, time.Time{})
	if err != nil {
		return err
	}
	err = validateType(expected[0], time.Time{})
	if err != nil {
		return err
	}
	return BeGreaterThan(actual, expected[0])
}
