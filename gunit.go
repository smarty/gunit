package gunit

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/smartystreets/assertions"
)

type Fixture struct{ *T }

func NewFixture(t *testing.T) *Fixture {
	return &Fixture{T: NewT(t)}
}

// So is a convenience method for reporting assertion failure messages,
// say from the assertion functions found in github.com/smartystreets/assertions/should.
// Example: self.So(actual, should.Equal, expected)
func (self *Fixture) So(actual interface{}, assert func(actual interface{}, expected ...interface{}) string, expected ...interface{}) {
	if ok, failure := assertions.So(actual, assert, expected...); !ok {
		self.T.Fail()
		self.T.Log("\t" + strings.Replace(failure, "\n", "\n\t\t", -1))
	}
}

func (self *Fixture) Finalize() {
	self.T.finalize()
}

//////////////////////////////////////////////////////////////////////////////

type T struct {
	*testing.T
	log     *bytes.Buffer
	skipped bool
}

func NewT(t *testing.T) *T {
	return &T{T: t, log: &bytes.Buffer{}}
}

func (self *T) finalize() {
	if self.log.Len() > 0 {
		if testing.Verbose() || self.T.Failed() {
			fmt.Println(self.log.String()) // Consider using self.T.Log(self.log.String()) // which will pollute the output with a line number from this file...
		}
	}
	if self.skipped {
		self.T.SkipNow()
	}
}

func (self *T) Error(args ...interface{}) {
	self.Log(args...)
	self.T.Fail()
}
func (self *T) Errorf(message string, args ...interface{}) {
	self.Logf(message, args...)
	self.T.Fail()
}

func (self *T) Skip(args ...interface{}) {
	self.Log(args...)
	self.skipped = true
}
func (self *T) Skipf(message string, args ...interface{}) {
	self.Logf(message, args...)
	self.skipped = true
}

func (self *T) Logf(message string, args ...interface{}) {
	self.log.WriteString("\t" + fmt.Sprintf(message, args...))

}
func (self *T) Log(args ...interface{}) {
	self.log.WriteString("\t" + fmt.Sprintln(args...))
}

//////////////////////////////////////////////////////////////////////////////

// TODO: implement the checksum validation function
func ValidateChecksums(checksums map[string]string) {
	// 1. For each file/checksum pair, verify that the file still exists and that the checksum matches the current contents
	// 2. Make sure there are not any additional files.
	// 3. If mismatch, fmt.Println("error!!!!") && os.Exit(1)
}
