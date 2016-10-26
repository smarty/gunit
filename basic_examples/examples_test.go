package examples

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestExample(t *testing.T) {
	gunit.Run(new(ExampleFixture), t)
}

type ExampleFixture struct {
	*gunit.Fixture // Required: Embedding this type is what makes the magic happen.

	// Declare useful state here (probably the stuff being testing, any fakes, etc...).
}

func (self *ExampleFixture) SetupStuff() {
	// This optional method will be executed before each "Test"
	// method because it starts with "Setup".
}
func (self *ExampleFixture) TeardownStuff() {
	// This optional method will be executed after each "Test"
	// method (because it starts with "Teardown"), even if the test method panics.
}

// This is an actual test case:
func (self *ExampleFixture) TestWithAssertions() {
	// Here's how to use the functions from the `should`
	// package at github.com/smartystreets/assertions/should
	// to perform assertions:
	self.So(42, should.Equal, 42)
	self.So("Hello, World!", should.ContainSubstring, "World")
	self.Ok(1 == 1, "One should equal one")
}

func (self *ExampleFixture) SkipTestWithNothing() {
	// Because this method's name starts with 'Skip', this will be skipped in the generated code.
}

func (self *ExampleFixture) LongTest() {
	// Because this method's name starts with 'Long', it will be skipped if the -short flag is passed to `go test`.
}
