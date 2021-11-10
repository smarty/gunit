package examples

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/smartystreets/gunit"
)

func TestExampleFixture(t *testing.T) {
	t.Name()
	gunit.Run(new(ExampleFixture), t, gunit.Options.SequentialTestCases())
}

type ExampleFixture struct {
	*gunit.Fixture // Required: Embedding this type is what allows gunit.Run to run the tests in this fixture.
	name           string
	age            int
	// Declare useful state here (probably the stuff being testing, any fakes, etc...).
}

func (this *ExampleFixture) SetupStuff() {
	// This optional method will be executed before each "Test"
	// method because it starts with "Setup".
	this.name = "rong.xu"
	this.T().Log(this.name)
}
func (this *ExampleFixture) TeardownStuff() {
	// This optional method will be executed after each "Test"
	// method (because it starts with "Teardown"), even if the test method panics.
}

func (this *ExampleFixture) FixtureSetupStuff() {
	// This optional method will be executed before each "Test"
	// method because it starts with "Setup".
	this.name = "yu.he"
	this.age = 30
	log.Println("in FixtureSetupStuff...")
	this.T().Log("in setup")
	fmt.Println("test name in setup", this.Name())
}
func (this *ExampleFixture) FixtureTeardownStuff() {
	// This optional method will be executed after each "Test"
	// method (because it starts with "Teardown"), even if the test method panics.
	this.name = "gannicus"
	log.Println("in FixtureTeardownStuff...")
	fmt.Println(this.name, time.Now(), this.Name())
}

func (this *ExampleFixture) SkipTestWithError() {
	this.Error("hi")
}

func (this *ExampleFixture) SkipTestWithErrorf() {
	this.Errorf("hi")
}

func (this *ExampleFixture) TestWithPrint() {
	this.Print("hi")
	log.Println(this.name, time.Now())
	time.Sleep(2 * time.Second)
	panic("panic one")
	this.T().Log("my name is ", this.name)
}

// This is an actual test case:
func (this *ExampleFixture) TestWithAssertions() {
	// Built-in assertion functions:
	this.Assert(1 == 1, "One should equal one")
	this.AssertEqual(1, 1)
	this.AssertDeepEqual(1, 1)
	this.AssertSprintEqual(1, 1.0)
	this.AssertSprintfEqual(uint(1), int64(1), "%d")

	// External assertion functions from the `should` package:
	// import "github.com/smartystreets/assertions/should"
	// ...
	// this.So(42, should.Equal, 42)
	// this.So("Hello, World!", should.ContainSubstring, "World")
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
