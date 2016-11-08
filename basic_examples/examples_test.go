package examples

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestExampleFixture(t *testing.T) {
	gunit.Run(new(ExampleFixture), t)
}

type ExampleFixture struct {
	*gunit.Fixture // Required: Embedding this type is what allows gunit.Run to run the tests in this fixture.

	// Declare useful state here (probably the stuff being testing, any fakes, etc...).
}

func (this *ExampleFixture) SetupStuff() {
	// This optional method will be executed before each "Test"
	// method because it starts with "Setup".
}
func (this *ExampleFixture) TeardownStuff() {
	// This optional method will be executed after each "Test"
	// method (because it starts with "Teardown"), even if the test method panics.
}

// This is an actual test case:
func (this *ExampleFixture) TestWithAssertions() {
	// Here's how to use the functions from the `should`
	// package at github.com/smartystreets/assertions/should
	// to perform assertions:
	this.So(42, should.Equal, 42)
	this.So("Hello, World!", should.ContainSubstring, "World")
	this.Ok(1 == 1, "One should equal one")
}

func (this *ExampleFixture) SkipTestWithNothing() {
	// Because this method's name starts with 'Skip', this will be skipped when `go test` is run.
}

func (this *ExampleFixture) LongTest() {
	// Because this method's name starts with 'Long', it will be skipped if the -short flag is passed to `go test`.
}

func (this *ExampleFixture) SkipLongTest() {
	// Because this method's name starts with 'Skip', it will be skipped when `go test` is run. Removing the 'Skip'
	// prefix would reveal the 'Long' prefix, explained above.
}
