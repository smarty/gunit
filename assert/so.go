package assert

func So(t testingT, actual any, assertion assertion, expected ...any) {
	t.Helper()
	result := assertion(actual, expected...)
	if result != "" {
		t.Error(result)
	}
}

type testingT interface {
	Helper()
	Error(...any)
}

// assertion is a copy of github.com/smarty/gunit.assertion. TODO: export and make this the canonical definition.
type assertion func(actual any, expected ...any) string
