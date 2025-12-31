package should

import "time"

// HappenWithin ensures that the first time value happens within
// a specified duration of the other time value.
// The actual value should be a time.Time.
// The first expected value should be a time.Duration.
// The second expected value should be a time.Time.
func HappenWithin(actual any, expected ...any) error {
	err := validateExpected(2, expected)
	if err != nil {
		return err
	}
	err = validateType(actual, time.Time{})
	if err != nil {
		return err
	}
	err = validateType(expected[0], time.Nanosecond)
	if err != nil {
		return err
	}
	err = validateType(expected[1], time.Time{})
	if err != nil {
		return err
	}
	a := actual.(time.Time)
	b := expected[1].(time.Time)
	diff := a.Sub(b).Abs()
	EXPECTED := expected[0].(time.Duration)
	if diff > EXPECTED {
		return failure("\n"+
			"Actual: %s\n"+
			"Target: %s\n"+
			"Max:    %s\n"+
			"Diff:   %s",
			a, b, EXPECTED, diff,
		)
	}
	return nil
}
