package assert

func So(t testingT, actual any, assertion Func, expected ...any) {
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

type Func func(actual any, expected ...any) string
