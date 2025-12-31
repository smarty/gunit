package should

// BeFalse verifies that actual is the boolean false value.
func BeFalse(actual any, expected ...any) error {
	err := validateExpected(0, expected)
	if err != nil {
		return err
	}

	err = validateType(actual, *new(bool))
	if err != nil {
		return err
	}

	boolean := actual.(bool)
	if boolean {
		return failure("got <true>, want <false>")
	}

	return nil
}
