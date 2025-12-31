package gunit

import (
	"context"
	"fmt"
	"io"
	"runtime/debug"
	"testing"
)

type Fixture struct{ TestingT }

func (this *Fixture) Print(a ...any)            { _, _ = fmt.Fprint(this.Output(), a...) }
func (this *Fixture) Printf(f string, a ...any) { _, _ = fmt.Fprintf(this.Output(), f, a...) }
func (this *Fixture) Println(a ...any)          { _, _ = fmt.Fprintln(this.Output(), a...) }

// So is a convenience method for reporting assertion failure messages
// with the many assertion functions found in github.com/smarty/gunit/v2/should.
// Example: this.So(actual, should.Equal, expected)
func (this *Fixture) So(actual any, assert Assertion, expected ...any) {
	this.Helper()
	So(this, actual, assert, expected...)
}

// Run is analogous to *testing.T.Run and allows for running subtests from
// test fixture methods (such as for table-driven tests).
func (this *Fixture) Run(name string, test func(fixture *Fixture)) {
	this.TestingT.(*testing.T).Run(name, func(t *testing.T) {
		t.Helper()
		fixture := &Fixture{TestingT: t}
		defer func() {
			if r := recover(); r != nil {
				fixture.Fail()
				fixture.Log(panicReport(r, debug.Stack()))
			}
		}()
		test(fixture)
	})
}

type TestingT interface {
	Cleanup(func())
	Context() context.Context
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
	Output() io.Writer
	Setenv(key, value string)
	Skip(args ...any)
	SkipNow()
	Skipf(format string, args ...any)
	Skipped() bool
	TempDir() string
}
