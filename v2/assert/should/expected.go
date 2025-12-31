package should

import "reflect"

func validateExpected(count int, expected []any) error {
	length := len(expected)
	if length == count {
		return nil
	}

	s := pluralize(length)
	return wrap(ErrExpectedCountInvalid, "got %d value%s, want %d", length, s, count)
}

func pluralize(count int) string {
	if count == 1 {
		return ""
	}
	return "s"
}

func validateType(actual, expected any) error {
	ACTUAL := reflect.TypeOf(actual)
	EXPECTED := reflect.TypeOf(expected)
	if ACTUAL == EXPECTED {
		return nil
	}
	return wrap(ErrTypeMismatch, "got %s, want %s", ACTUAL, EXPECTED)
}

func validateKind(actual any, kinds ...reflect.Kind) error {
	value := reflect.ValueOf(actual)
	kind := value.Kind()
	for _, k := range kinds {
		if k == kind {
			return nil
		}
	}
	return wrap(ErrKindMismatch, "got %s, want one of %v", kind, kinds)
}
