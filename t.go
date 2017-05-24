package gunit

// tt represents the functional subset from *testing.T needed by Fixture.
type tt interface {
	Log(args ...interface{})
	Fail()
	Failed() bool
}
