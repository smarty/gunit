package should

import (
	"errors"
	"sort"
	"time"
)

// BeChronological asserts whether actual is a []time.Time and
// whether the values are in chronological order.
func BeChronological(actual any, expected ...any) error {
	err := validateExpected(0, expected)
	if err != nil {
		return err
	}

	var t []time.Time
	err = validateType(actual, t)
	if err != nil {
		return err
	}

	times := actual.([]time.Time)
	if sort.SliceIsSorted(times, func(i, j int) bool { return times[i].Before(times[j]) }) {
		return nil
	}
	return failure("expected to be chronological: %v", times)
}

// BeChronological (negated!)
func (negated) BeChronological(actual any, expected ...any) error {
	err := BeChronological(actual, expected...)
	if errors.Is(err, ErrAssertionFailure) {
		return nil
	}
	if err != nil {
		return err
	}
	return failure("want non-chronological times, got chronological times:", actual)
}
