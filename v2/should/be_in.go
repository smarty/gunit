package should

// BeIn determines whether actual is a member of expected[0].
// It defers to Contain.
func BeIn(actual any, expected ...any) error {
	err := validateExpected(1, expected)
	if err != nil {
		return err
	}

	err = Contain(expected[0], actual)
	if err != nil {
		return err
	}

	return nil
}

// BeIn (negated!)
func (negated) BeIn(actual any, expected ...any) error {
	err := validateExpected(1, expected)
	if err != nil {
		return err
	}

	return NOT.Contain(expected[0], actual)
}
