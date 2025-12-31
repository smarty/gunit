package should

import "reflect"

// HaveLength uses reflection to verify that len(actual) == 0.
func HaveLength(actual any, expected ...any) error {
	err := validateExpected(1, expected)
	if err != nil {
		return err
	}

	err = validateKind(actual, kindsWithLength...)
	if err != nil {
		return err
	}

	err = validateKind(expected[0], kindSlice(signedIntegerKinds)...)
	if err != nil {
		return err
	}

	expectedLength := reflect.ValueOf(expected[0]).Int()
	actualLength := int64(reflect.ValueOf(actual).Len())
	if actualLength == expectedLength {
		return nil
	}

	return failure("got length of %d, want %d", actualLength, expectedLength)
}
