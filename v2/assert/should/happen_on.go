package should

import (
	"errors"
	"time"
)

// HappenOn ensures that two time values happen at the same instant.
// See the time.Time.Equal method for the details.
// This function defers to Equal to do the work.
func HappenOn(actual any, expected ...any) error {
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
	return Equal(actual, expected...)
}

// HappenOn negated!
func (negated) HappenOn(actual any, expected ...any) error {
	err := HappenOn(actual, expected...)
	if errors.Is(err, ErrAssertionFailure) {
		return nil
	}
	if err != nil {
		return err
	}

	return failure("\n"+
		"  expected:     %#v\n"+
		"  to not equal: %#v\n"+
		"  (but it did)",
		expected[0],
		actual,
	)
}
