package examples

//go:generate gunit

import (
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

type ExampleFixture struct {
	*gunit.Fixture // Required: Embedding this type is what makes the magic happen.

	// Declare useful state here (probably the stuff being testing, any fakes, etc...).
}

func SetupExampleFixture() {
	// This optional, fixture-level setup function will be run once,
	// before any "Test" methods on the ExampleFixture are executed.
}
func TeardownExampleFixture() {
	// This optional, fixture-level function will be run once,
	// after all "Test" methods on the ExampleFixture have been executed.
}

func (self *ExampleFixture) SetupStuff() {
	// This optional method will be executed before each "Test"
	// method becuase it starts with "Setup".
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
}

func (self *ExampleFixture) SkipTestWithNothing() {
	// Because this method's name starts with 'Skip', this will be skipped in the generated code.
}
