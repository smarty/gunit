// Package gunit provides "testing" package hooks and convenience
// functions for writing tests in an xUnit style.
// See the README file and the examples folder for examples.
package gunit

import (
	"bytes"
	"fmt"
	"reflect"
	"runtime/debug"
	"strings"
	"testing"

	"github.com/smarty/gunit/reports"
)

// Fixture keeps track of test status (failed, passed, skipped) and
// handles custom logging for xUnit style tests as an embedded field.
// The Fixture manages an instance of *testing.T. Certain methods
// defined herein merely forward to calls on the *testing.T:
//
//   - Fixture.Error(...) ----> *testing.T.Error
//   - Fixture.Errorf(...) ---> *testing.T.Errorf
//   - Fixture.Print(...) ----> *testing.T.Log or fmt.Print
//   - Fixture.Printf(...) ---> *testing.T.Logf or fmt.Printf
//   - Fixture.Println(...) --> *testing.T.Log or fmt.Println
//   - Fixture.Failed() ------> *testing.T.Failed()
//   - Fixture.fail() --------> *testing.T.Fail()
//
// We don't use these methods much, preferring instead to lean heavily
// on Fixture.So and the rich set of should-style assertions provided at
// github.com/smarty/assertions/should
type Fixture struct {
	t       TestingT
	log     *bytes.Buffer
	verbose bool
}

func newFixture(t TestingT, verbose bool) *Fixture {
	t.Helper()

	return &Fixture{t: t, verbose: verbose, log: &bytes.Buffer{}}
}

// T exposes the TestingT (*testing.T) instance.
func (this *Fixture) T() TestingT { return this.t }

// Run is analogous to *testing.T.Run and allows for running subtests from
// test fixture methods (such as for table-driven tests).
func (this *Fixture) Run(name string, test func(fixture *Fixture)) {
	this.t.(*testing.T).Run(name, func(t *testing.T) {
		t.Helper()

		fixture := newFixture(t, this.verbose)
		defer fixture.finalize()
		test(fixture)
	})
}

// So is a convenience method for reporting assertion failure messages,
// from the many assertion functions found in github.com/smarty/assertions/should.
// Example: this.So(actual, should.Equal, expected)
func (this *Fixture) So(actual any, assert assertion, expected ...any) bool {
	failure := assert(actual, expected...)
	failed := len(failure) > 0
	if failed {
		this.fail(failure)
	}
	return !failed
}

// Assert tests a boolean which, if not true, marks the current test case as failed and
// prints the provided message.
func (this *Fixture) Assert(condition bool, messages ...string) bool {
	if !condition {
		if len(messages) == 0 {
			messages = append(messages, "Expected condition to be true, was false instead.")
		}
		this.fail(strings.Join(messages, ", "))
	}
	return condition
}
func (this *Fixture) AssertEqual(expected, actual any) bool {
	return this.Assert(expected == actual, fmt.Sprintf(comparisonFormat, fmt.Sprint(expected), fmt.Sprint(actual)))
}
func (this *Fixture) AssertSprintEqual(expected, actual any) bool {
	return this.AssertEqual(fmt.Sprint(expected), fmt.Sprint(actual))
}
func (this *Fixture) AssertSprintfEqual(expected, actual any, format string) bool {
	return this.AssertEqual(fmt.Sprintf(format, expected), fmt.Sprintf(format, actual))
}
func (this *Fixture) AssertDeepEqual(expected, actual any) bool {
	return this.Assert(reflect.DeepEqual(expected, actual),
		fmt.Sprintf(comparisonFormat, fmt.Sprintf("%#v", expected), fmt.Sprintf("%#v", actual)))
}

func (this *Fixture) Error(args ...any)            { this.fail(fmt.Sprint(args...)) }
func (this *Fixture) Errorf(f string, args ...any) { this.fail(fmt.Sprintf(f, args...)) }

func (this *Fixture) Print(a ...any)            { fmt.Fprintln(this.log, a...) }
func (this *Fixture) Printf(f string, a ...any) { fmt.Fprintln(this.log, fmt.Sprintf(f, a...)) }
func (this *Fixture) Println(a ...any)          { fmt.Fprintln(this.log, a...) }

// Write implements io.Writer. There are rare times when this is convenient (debugging via `log.SetOutput(fixture)`).
func (this *Fixture) Write(p []byte) (int, error) { return this.log.Write(p) }
func (this *Fixture) Failed() bool                { return this.t.Failed() }
func (this *Fixture) Name() string                { return this.t.Name() }

func (this *Fixture) fail(failure string) {
	this.t.Fail()
	this.Print(reports.FailureReport(failure, reports.StackTrace()))
}

func (this *Fixture) finalize() {
	this.t.Helper()

	if r := recover(); r != nil {
		this.recoverPanic(r)
	}

	if this.t.Failed() || (this.verbose && this.log.Len() > 0) {
		this.t.Log("\n" + strings.TrimSpace(this.log.String()) + "\n")
	}
}
func (this *Fixture) recoverPanic(r any) {
	this.t.Fail()
	this.Print(reports.PanicReport(r, debug.Stack()))
}

const comparisonFormat = "Expected: [%s]\nActual:   [%s]"

// assertion is a copy of github.com/smarty/assertions.assertion.
type assertion func(actual any, expected ...any) string
