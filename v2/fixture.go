package gunit

type Fixture struct{ TestingT }

// Write implements io.Writer, which is convenient when using a fixture
// as a log target.
func (this *Fixture) Write(p []byte) (int, error) {
	this.Helper()
	this.Log(string(p))
	return len(p), nil
}

// So is a convenience method for reporting assertion failure messages
// with the many assertion functions found in github.com/smarty/gunit/v2/should.
// Example: this.So(actual, should.Equal, expected)
func (this *Fixture) So(actual any, assert assertion, expected ...any) bool {
	err := assert(actual, expected...)
	if err != nil {
		this.Helper()
		this.Error(err)
	}
	return err == nil
}

type assertion func(actual any, expected ...any) error

type TestingT interface {
	Cleanup(func())
	Error(args ...any)
	Errorf(format string, args ...any)
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...any)
	Fatalf(format string, args ...any)
	Helper()
	Log(args ...any)
	Logf(format string, args ...any)
	Name() string
	Setenv(key, value string)
	Skip(args ...any)
	SkipNow()
	Skipf(format string, args ...any)
	Skipped() bool
	TempDir() string
}
