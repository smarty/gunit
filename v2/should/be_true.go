package should

// BeTrue verifies that actual is the boolean true value.
func BeTrue(actual any, expected ...any) error {
	err := validateExpected(0, expected)
	if err != nil {
		return err
	}

	err = validateType(actual, *new(bool))
	if err != nil {
		return err
	}

	boolean := actual.(bool)
	if !boolean {
		return failure("got <false>, want <true>")
	}
	return nil
}
