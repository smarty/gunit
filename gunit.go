// Package gunit provides "testing" package hooks and convenience
// functions for writing tests in an xUnit style.
// NOTE: Only some of the exported names in this package
// are meant to be referenced by users of this package:
//
// - Fixture              // (as an embedded field on your xUnit-style struct)
// - Fixture.So(...)      // (as a convenient assertion method: So(expected, should.Equal, actual))
// - Fixture.Ok(...)      // (as a convenient boolean assertion method: Ok(condition, optionalMessage))
// - Fixture.Error(...)   // (works just like *testing.T.Error)
// - Fixture.Errorf(...)  // (works just like *testing.T.Errorf)
// - Fixture.Print(...)   // (works just like fmt.Print)
// - Fixture.Printf(...)  // (works just like fmt.Printf)
// - Fixture.Println(...) // (works just like fmt.Println)
//
// The rest are called from code generated by the command at
// github.com/smartystreets/gunit/gunit.
// Please see the README file and the examples folder for examples.
package gunit

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"

	"github.com/smartystreets/assertions"
)

// tt represents the functional subset from *testing.T needed by Fixture.
type tt interface {
	Log(args ...interface{})
	Fail()
	Failed() bool
}

// Fixture keeps track of test status (failed, passed, skipped) and
// handles custom logging for xUnit style tests as an embedded field.
type Fixture struct {
	t       tt
	log     *bytes.Buffer
	verbose bool
}

// NewFixture is called by generated code.
func NewFixture(t tt, verbose bool) *Fixture { // FUTURE: un-export this along with deletion of 'gunit' command code.
	return &Fixture{t: t, verbose: verbose, log: &bytes.Buffer{}}
}

// Parallel is now deprecated. At one point it was analogous to *testing.T.Parallel.
func (this *Fixture) Parallel() {
	this.Println("[WARNING] Fixture.Parallel() is now deprecated and will soon be removed.")
}

// So is a convenience method for reporting assertion failure messages,
// say from the assertion functions found in github.com/smartystreets/assertions/should.
// Example: this.So(actual, should.Equal, expected)
func (this *Fixture) So(
	actual interface{},
	assert func(actual interface{}, expected ...interface{}) string,
	expected ...interface{}) bool {

	ok, failure := assertions.So(actual, assert, expected...)
	if !ok {
		this.t.Fail()
		this.reportFailure(failure)
	}
	return ok
}

func (this *Fixture) Ok(condition bool, messages ...string) {
	if !condition {
		if len(messages) == 0 {
			messages = append(messages, "Expected condition to be true, was false instead.")
		}
		this.t.Fail()
		this.reportFailure(strings.Join(messages, ", "))
	}
}

func (this *Fixture) Error(args ...interface{}) {
	this.t.Fail()
	this.reportFailure(fmt.Sprint(args...))
}

func (this *Fixture) Errorf(format string, args ...interface{}) {
	this.t.Fail()
	this.reportFailure(fmt.Sprintf(format, args...))
}

func (this *Fixture) reportFailure(failure string) {
	this.Print(newFailureReport(failure))
}

// Print is nearly analogous to fmt.Print and is ideal for printing in the middle of a test case.
func (this *Fixture) Print(a ...interface{}) {
	fmt.Fprint(this, a...)
}

// Printf is nearly analogous to fmt.Printf and is ideal for printing in the middle of a test case.
func (this *Fixture) Printf(format string, a ...interface{}) {
	fmt.Fprintf(this, format, a...)
}

// Println is nearly analogous to fmt.Println and is ideal for printing in the middle of a test case.
func (this *Fixture) Println(a ...interface{}) {
	fmt.Fprintln(this, a...)
}

// Write implements io.Writer.
func (this *Fixture) Write(p []byte) (int, error) {
	return this.log.Write(p)
}

// Finalize is called by generated code.
func (this *Fixture) Finalize() {
	if r := recover(); r != nil {
		this.recoverPanic(r)
	}

	if this.t.Failed() || (this.verbose && this.log.Len() > 0) {
		this.t.Log("\n" + strings.TrimSpace(this.log.String()) + "\n")
	}
}

// Failed is analogous to *testing.T.Failed().
func (this *Fixture) Failed() bool {
	return this.t.Failed()
}

func (this *Fixture) recoverPanic(r interface{}) {
	this.Println("PANIC:", r)
	buffer := make([]byte, 1024*16)
	runtime.Stack(buffer, false)
	this.Println(strings.TrimSpace(string(buffer)))
	this.Println("* (Additional tests may have been skipped as a result of the panic shown above.)")
	this.t.Fail()
}
