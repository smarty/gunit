package should

import (
	"errors"
	"reflect"
	"strings"
)

// Contain determines whether actual contains expected[0].
// The actual value may be a map, array, slice, or string:
//   - In the case of maps the expected value is assumed to be a map key.
//   - In the case of slices and arrays the expected value is assumed to be a member.
//   - In the case of strings the expected value may be a rune or substring.
func Contain(actual any, expected ...any) error {
	err := validateExpected(1, expected)
	if err != nil {
		return err
	}

	err = validateKind(actual, containerKinds...)
	if err != nil {
		return err
	}

	actualValue := reflect.ValueOf(actual)
	EXPECTED := expected[0]

	switch reflect.TypeOf(actual).Kind() {
	case reflect.Map:
		expectedValue := reflect.ValueOf(EXPECTED)
		value := actualValue.MapIndex(expectedValue)
		if value.IsValid() {
			return nil
		}
	case reflect.Array, reflect.Slice:
		for i := 0; i < actualValue.Len(); i++ {
			item := actualValue.Index(i).Interface()
			if Equal(EXPECTED, item) == nil {
				return nil
			}
		}
	case reflect.String:
		err = validateKind(EXPECTED, reflect.String, reflectRune)
		if err != nil {
			return err
		}

		expectedRune, ok := EXPECTED.(rune)
		if ok {
			EXPECTED = string(expectedRune)
		}

		full := actual.(string)
		sub := EXPECTED.(string)
		if strings.Contains(full, sub) {
			return nil
		}
	}

	return failure("\n"+
		"   item absent: %#v\n"+
		"   within:      %#v",
		EXPECTED,
		actual,
	)
}

// Contain (negated!)
func (negated) Contain(actual any, expected ...any) error {
	err := Contain(actual, expected...)
	if errors.Is(err, ErrAssertionFailure) {
		return nil
	}

	if err != nil {
		return err
	}

	return failure("\n"+
		"item found: %#v\n"+
		"within:     %#v",
		expected[0],
		actual,
	)
}

const reflectRune = reflect.Int32
