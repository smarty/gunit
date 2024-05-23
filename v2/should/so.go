package should

func So(t testingT, actual any, assertion Assertion, expected ...any) {
	t.Helper()
	err := assertion(actual, expected...)
	if err != nil { // TODO: differentiate between should.* and must.*
		t.Error(err)
	}
}

type testingT interface {
	Helper()
	Error(...any)
}
type Assertion func(actual any, expected ...any) error
