package gunit

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/smartystreets/assertions"
)

type TestCase struct {
	T *T
}

func NewTestCase(t *testing.T) *TestCase {
	return &TestCase{
		T: &T{T: t},
	}
}

// So is a convenience method for reporting assertion failure messages,
// say from the assertion functions found in github.com/smartystreets/assertions/should.
// Example: self.So(actual, should.Equal, expected)
func (self TestCase) So(actual interface{}, assert func(actual interface{}, expected ...interface{}) string, expected ...interface{}) {
	if ok, failure := assertions.So(actual, assert, expected...); !ok {
		self.T.Fail()
		fmt.Print(report(failure, 3)) // Log + test function + test case method.
	}
}

//////////////////////////////////////////////////////////////////////////////

type T struct{ *testing.T }

func (self *T) Error(args ...interface{}) {
	self.Log(args...)
	self.T.Fail()
}
func (self *T) Errorf(message string, args ...interface{}) {
	self.Logf(message, args...)
	self.T.Fail()
}

func (self *T) Fatal(args ...interface{}) {
	self.Log(args...)
	self.T.FailNow()
}
func (self *T) Fatalf(message string, args ...interface{}) {
	self.Logf(message, args...)
	self.T.FailNow()
}

func (self *T) Skip(args ...interface{}) {
	// TODO: this should mark the current test case as skipped, but should not call
	// self.T.SkipNow() until the end of the fixture-level teardown because it calls
	// runtime.Goexit().
	self.Log(args...)
	self.T.SkipNow()
}
func (self *T) Skipf(message string, args ...interface{}) {
	// TODO: see comment about Skip() (above).
	self.Logf(message, args...)
	self.T.SkipNow()
}

func (self *T) Logf(message string, args ...interface{}) {
	fmt.Print(report(fmt.Sprintf(message, args...), 4)) // report + Log + test function + test case method.

}
func (self *T) Log(args ...interface{}) {
	fmt.Print(report(fmt.Sprintln(args...), 4)) // report + Log + test function + test case method.
}

//////////////////////////////////////////////////////////////////////////////

func report(message string, depth int) string {
	file, line := localizeCaller(depth)
	return indent(message, file, line)
}
func localizeCaller(depth int) (file string, line int) {
	// TODO: this should traverse the stack until it finds the right line (??).
	_, file, line, ok := runtime.Caller(depth)
	if ok {
		file = file[strings.LastIndex(file, string(os.PathSeparator))+1:]
	} else {
		file, line = "<unknown_file>", 0
	}
	return file, line
}
func indent(message, file string, line int) string {
	localized := strings.TrimSpace(location(file, line) + message)
	indented := strings.Replace(localized, "\n", "\n\t\t", -1)
	return "\t" + indented + "\n"
}
func location(file string, line int) string {
	return fmt.Sprintf("%s:%d \n", file, line)
}

//////////////////////////////////////////////////////////////////////////////
