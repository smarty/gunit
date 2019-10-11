package gunit

// TestingT represents the functional subset from *testing.T needed by Fixture.
type TestingT interface {
	Name() string
	Log(args ...interface{})
	Fail()
	Failed() bool
	Fatalf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}
