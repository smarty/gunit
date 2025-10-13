package assert

import "strings"

func So(t testingT, actual any, assertion Func, expected ...any) {
	t.Helper()
	result := assertion(actual, expected...)
	if strings.HasPrefix(result, "<<<FATAL>>>\n") {
		t.Fatal(result)
	} else if result != "" {
		t.Error(result)
	}
}

type testingT interface {
	Helper()
	Error(...any)
	Fatal(...any)
}

type Func func(actual any, expected ...any) string
