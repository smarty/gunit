package should

import (
	"reflect"
	"strings"
)

// StartWith verified that actual starts with expected[0].
// The actual value may be an array, slice, or string.
func StartWith(actual any, expected ...any) error {
	err := validateExpected(1, expected)
	if err != nil {
		return err
	}

	err = validateKind(actual, orderedContainerKinds...)
	if err != nil {
		return err
	}

	actualValue := reflect.ValueOf(actual)
	EXPECTED := expected[0]

	switch reflect.TypeOf(actual).Kind() {
	case reflect.Array, reflect.Slice:
		if actualValue.Len() == 0 {
			break
		}
		first := actualValue.Index(0).Interface()
		if Equal(EXPECTED, first) == nil {
			return nil
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
		prefix := EXPECTED.(string)
		if strings.HasPrefix(full, prefix) {
			return nil
		}
	}

	return failure("\n"+
		"   proposed prefix: %#v\n"+
		"   not a prefix of: %#v",
		EXPECTED,
		actual,
	)
}
