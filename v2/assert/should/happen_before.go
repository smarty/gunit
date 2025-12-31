package should

import "time"

// HappenBefore ensures that the first time value happens before the second.
func HappenBefore(actual any, expected ...any) error {
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
	return BeLessThan(actual, expected[0])
}
