package examples

import (
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

type ExampleFixture struct {
	*gunit.Fixture

	// Declare useful state here (probably the stuff being testing, any fakes, etc...).
}

func SetupExampleFixture() {
	// This function will be run once, before any "Test" methods on the ExampleFixture are executed.
}
func TeardownExampleFixture() {
	// This function will be run once, after all "Test" methods on the ExampleFixture have been executed.
}

func (self *ExampleFixture) SetupStuff() {
	// This method will be executed before each "Test" method becuase it starts with "Setup".
}
func (self *ExampleFixture) TeardownStuff() {
	// This method will be executed after each "Test" method (because it starts with "Teardown"), even if the test method panics.
}

// This function demonstrates how to use the should package from
// SmartyStreets to perform assertions:
func (self *ExampleFixture) TestWithAssertions() {
	self.So(42, should.Equal, 42)
	self.So("Hello, World!", should.ContainSubstring, "o, W")
}

func (self *ExampleFixture) SkipTestWithNothing() {
	// Because this method's name starts with 'Skip', this will be skipped in the generated code.
}
