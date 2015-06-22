# gunit

## Installation

```
$ go get github.com/smartystreets/gunit/gunit
```

-------------------------

We now present `gunit`, yet another testing tool for Go.

> Not again... ([GoConvey](http://goconvey.co) was crazy enough...but sort of cool, ok I'll pay attention...)

No wait, this tool has some very interesting properties. It's a mix of good things provided by the built-in testing package, the [assertions](https://github.com/smartystreets/assertions) you know and love from the [GoConvey](http://goconvey.co) project, the [xUnit](https://en.wikipedia.org/wiki/XUnit) testing style (the first real unit testing framework), and it's all glued together with `go test`.

> Blah, blah, yeah, yeah. Ok, so what's wrong with just using the standard "testing" package? What's better about this `gunit` thing?

The convention established by the "testing" package and the `go test` tool only allows for local function scope:

```
func TestSomething(t *testing.T) {
	// blah blah blah
}
```

This limited scope makes extracting functions or structs inconvenient as state will have to be passed to such extractions or state returned from them. It can get messy to keep a test nice and short. Here's the basic idea of what the test author using `gunit` would implement in a `*_test.go` file:

```go

package examples

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
	// Because this method's name starts with 'Skip', it will be skipped.
}

func (self *ExampleFixture) LongTestSlowOperation() {
	// Because this method's name starts with 'Long', it will be skipped if `go test` is run with the `short` flag.
	time.Sleep(time.Hour)
	So(true, should.BeTrue)
}
```

-------------------------

> So, this doesn't import the standard go testing package at all, right? Do I have to run some other command to run my tests?

You're correct, the code you see above doesn't have anything to do with the `"testing"` package. But you still run `go test` to execute those tests...

> Wait, without any test functions (you know, `func TestSomething(t *testing.T) {...}`) and without any reference to `*testing.T` anywhere, how do you mark a test as failed? You're saying I can still run `go test`? I don't get it.

Astute observations. `gunit` allows the test author to use a _struct_ as the scope for a group of related test cases, in the style of [xUnit](https://en.wikipedia.org/wiki/XUnit) fixtures. This makes extraction of setup/teardown behavior (as well as invoking the system under test) much simpler because all state for the test can be declared as fields on a struct which embeds the `Fixture` type from the `gunit` package.

Your question about the missing `func Test...` and the non-existent `*testing.T` is relevant. The missing link is a [command](https://github.com/smartystreets/gunit/gunit) that comes with the `gunit` project that scans your test fixtures and generates test functions that call all the appropriate methods for you! `*testing.T` is wrapped up by that generated code and you don't need to worry about calling any methods on it.

> Wow, that sounds strangely cool and border-line wrong at the same time.

In either case, here's what the generated code looks like:

```go

//////////////////////////////////////////////////////////////////////////////
// Generated Code ////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

package examples

import (
	"testing"

	"github.com/smartystreets/gunit"
)

///////////////////////////////////////////////////////////////////////////////

func TestExampleFixture(t *testing.T) {
	defer TeardownExampleFixture()
	SetupExampleFixture()

	fixture := gunit.NewFixture(t)
	defer fixture.Finalize()

	test0 := &ExampleFixture{Fixture: fixture}
	test0.RunTestCase__(test0.TestWithAssertions, "Test with assertions")

	fixture.Skip("Skipping test case: 'Skip test with nothing'")
}

func (self *ExampleFixture) RunTestCase__(test func(), description string) {
	self.Describe(description)
	defer self.TeardownStuff()
	self.SetupStuff()
	test()
}

///////////////////////////////////////////////////////////////////////////////

func init() {
	gunit.Validate("98884d1f827ddcee8a923e672c3cf2ba")
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////// Generated Code //
///////////////////////////////////////////////////////////////////////////////

```

> What's with the `init` function at the bottom of the generated code?

Ah, that. The call to `gunit.Validate` passes an md5 checksum of the contents of all `*.go` files in the package when the code was generated. Any change to those files between initial generation and test execution will necessitate regenerating the code, by running the [`gunit` command](https://github.com/smartystreets/gunit/gunit).

> Oh, so that prevents your struct-based tests from getting out of sync with the test functions `go test` expects to run.

Exactly. And you can invoke the `gunit` command be calling `go generate` if you put the following comment somewhere in your package (even in a `*_test.go` file):

```
//go:generate gunit
```

We use a script that runs our tests automatically whenever a `*.go` file changes (and it also runs `go generate`). Depending on the number of test fixtures in a package it generally takes a hundredth of a second to run. Your mileage may vary. Enjoy.

[Advanced Examples](https://github.com/smartystreets/gunit/tree/master/advanced_examples)

----------------------------------------------------------------------------
