package assertions

import (
	"fmt"
	"reflect"
	"strings"
)

// ShouldStartWith receives exactly 2 string parameters and ensures that the first starts with the second.
func ShouldStartWith(actual any, expected ...any) string {
	if fail := need(1, expected); fail != success {
		return fail
	}

	value, valueIsString := actual.(string)
	prefix, prefixIsString := expected[0].(string)

	if !valueIsString || !prefixIsString {
		return fmt.Sprintf(shouldBothBeStrings, reflect.TypeOf(actual), reflect.TypeOf(expected[0]))
	}

	return shouldStartWith(value, prefix)
}
func shouldStartWith(value, prefix string) string {
	if !strings.HasPrefix(value, prefix) {
		shortval := value
		if len(shortval) > len(prefix) {
			shortval = shortval[:len(prefix)] + "..."
		}
		return fmt.Sprintf(shouldHaveStartedWith, value, prefix)
	}
	return success
}

// ShouldNotStartWith receives exactly 2 string parameters and ensures that the first does not start with the second.
func ShouldNotStartWith(actual any, expected ...any) string {
	if fail := need(1, expected); fail != success {
		return fail
	}

	value, valueIsString := actual.(string)
	prefix, prefixIsString := expected[0].(string)

	if !valueIsString || !prefixIsString {
		return fmt.Sprintf(shouldBothBeStrings, reflect.TypeOf(actual), reflect.TypeOf(expected[0]))
	}

	return shouldNotStartWith(value, prefix)
}
func shouldNotStartWith(value, prefix string) string {
	if strings.HasPrefix(value, prefix) {
		if value == "" {
			value = "<empty>"
		}
		if prefix == "" {
			prefix = "<empty>"
		}
		return fmt.Sprintf(shouldNotHaveStartedWith, value, prefix)
	}
	return success
}

// ShouldEndWith receives exactly 2 string parameters and ensures that the first ends with the second.
func ShouldEndWith(actual any, expected ...any) string {
	if fail := need(1, expected); fail != success {
		return fail
	}

	value, valueIsString := actual.(string)
	suffix, suffixIsString := expected[0].(string)

	if !valueIsString || !suffixIsString {
		return fmt.Sprintf(shouldBothBeStrings, reflect.TypeOf(actual), reflect.TypeOf(expected[0]))
	}

	return shouldEndWith(value, suffix)
}
func shouldEndWith(value, suffix string) string {
	if !strings.HasSuffix(value, suffix) {
		shortval := value
		if len(shortval) > len(suffix) {
			shortval = "..." + shortval[len(shortval)-len(suffix):]
		}
		return fmt.Sprintf(shouldHaveEndedWith, value, suffix)
	}
	return success
}

// ShouldNotEndWith receives exactly 2 string parameters and ensures that the first does not end with the second.
func ShouldNotEndWith(actual any, expected ...any) string {
	if fail := need(1, expected); fail != success {
		return fail
	}

	value, valueIsString := actual.(string)
	suffix, suffixIsString := expected[0].(string)

	if !valueIsString || !suffixIsString {
		return fmt.Sprintf(shouldBothBeStrings, reflect.TypeOf(actual), reflect.TypeOf(expected[0]))
	}

	return shouldNotEndWith(value, suffix)
}
func shouldNotEndWith(value, suffix string) string {
	if strings.HasSuffix(value, suffix) {
		if value == "" {
			value = "<empty>"
		}
		if suffix == "" {
			suffix = "<empty>"
		}
		return fmt.Sprintf(shouldNotHaveEndedWith, value, suffix)
	}
	return success
}

// ShouldContainSubstring receives exactly 2 string parameters and ensures that the first contains the second as a substring.
func ShouldContainSubstring(actual any, expected ...any) string {
	if fail := need(1, expected); fail != success {
		return fail
	}

	long, longOk := actual.(string)
	short, shortOk := expected[0].(string)

	if !longOk || !shortOk {
		return fmt.Sprintf(shouldBothBeStrings, reflect.TypeOf(actual), reflect.TypeOf(expected[0]))
	}

	if !strings.Contains(long, short) {
		return fmt.Sprintf(shouldHaveContainedSubstring, long, short)
	}
	return success
}

// ShouldNotContainSubstring receives exactly 2 string parameters and ensures that the first does NOT contain the second as a substring.
func ShouldNotContainSubstring(actual any, expected ...any) string {
	if fail := need(1, expected); fail != success {
		return fail
	}

	long, longOk := actual.(string)
	short, shortOk := expected[0].(string)

	if !longOk || !shortOk {
		return fmt.Sprintf(shouldBothBeStrings, reflect.TypeOf(actual), reflect.TypeOf(expected[0]))
	}

	if strings.Contains(long, short) {
		return fmt.Sprintf(shouldNotHaveContainedSubstring, long, short)
	}
	return success
}

// ShouldBeBlank receives exactly 1 string parameter and ensures that it is equal to "".
func ShouldBeBlank(actual any, expected ...any) string {
	if fail := need(0, expected); fail != success {
		return fail
	}
	value, ok := actual.(string)
	if !ok {
		return fmt.Sprintf(shouldBeString, reflect.TypeOf(actual))
	}
	if value != "" {
		return fmt.Sprintf(shouldHaveBeenBlank, value)
	}
	return success
}

// ShouldNotBeBlank receives exactly 1 string parameter and ensures that it is not equal to "".
func ShouldNotBeBlank(actual any, expected ...any) string {
	if fail := need(0, expected); fail != success {
		return fail
	}
	value, ok := actual.(string)
	if !ok {
		return fmt.Sprintf(shouldBeString, reflect.TypeOf(actual))
	}
	if value == "" {
		return shouldNotHaveBeenBlank
	}
	return success
}
