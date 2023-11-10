package gunit

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

/*
Run accepts a fixture with Test* methods and
optional setup/teardown methods and executes
the suite. Fixtures must be struct types which
embeds a *gunit.T. Assuming a fixture struct
with test methods 'Test1' and 'Test2' execution
would proceed as follows:

 1. fixture.SetupSuite()
 2. fixture.Setup()
 3. fixture.Test1()
 4. fixture.Teardown()
 5. fixture.Setup()
 6. fixture.Test2()
 7. fixture.Teardown()
 8. fixture.TeardownSuite()

The methods provided by Options may be supplied
to this function to tweak the execution.
*/
func Run(fixture any, options ...Option) {
	config := new(config)
	for _, option := range append(defaultOptions, options...) {
		option(config)
	}

	fixtureValue := reflect.ValueOf(fixture)
	fixtureType := reflect.TypeOf(fixture)
	t, ok := tryExtractT(fixtureValue)
	if !ok {
		panic("Failed to extract *testing.T via embedded *Fixture instance.")
	}

	var (
		testNames        []string
		skippedTestNames []string
		focusedTestNames []string
	)
	for x := 0; x < fixtureType.NumMethod(); x++ {
		name := fixtureType.Method(x).Name
		method := fixtureValue.MethodByName(name)
		_, isNiladic := method.Interface().(func())
		if !isNiladic {
			continue
		}

		if strings.HasPrefix(name, "Test") {
			testNames = append(testNames, name)
		} else if strings.HasPrefix(name, "SkipTest") {
			skippedTestNames = append(skippedTestNames, name)
		} else if strings.HasPrefix(name, "FocusTest") {
			focusedTestNames = append(focusedTestNames, name)
		}
	}

	if len(focusedTestNames) > 0 {
		testNames = focusedTestNames
	}

	if len(testNames) == 0 {
		t.Skip("NOT IMPLEMENTED (no test cases defined, or they are all marked as skipped)")
		return
	}

	if config.parallelFixture {
		t.Parallel()
	}

	setup, hasSetup := fixture.(setupSuite)
	if hasSetup {
		setup.SetupSuite()
	}

	teardown, hasTeardown := fixture.(teardownSuite)
	if hasTeardown {
		defer teardown.TeardownSuite()
	}

	for _, name := range skippedTestNames {
		testCase{t: t, manualSkip: true, name: name}.Run()
	}

	for _, name := range testNames {
		testCase{t, name, config, false, fixtureType, fixtureValue}.Run()
	}
}

func tryExtractT(fixtureValue reflect.Value) (t *testing.T, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			ok = false
		}
	}()
	return fixtureValue.Elem().FieldByName("T").Elem().FieldByName("TestingT").Interface().(*testing.T), true
}

type testCase struct {
	t            *testing.T
	name         string
	config       *config
	manualSkip   bool
	fixtureType  reflect.Type
	fixtureValue reflect.Value
}

func (this testCase) Run() {
	_ = this.t.Run(this.name, this.decideRun())
}
func (this testCase) decideRun() func(*testing.T) {
	if this.manualSkip {
		return this.skipFunc("Skipping: " + this.name)
	}
	return this.runTest
}
func (this testCase) skipFunc(message string) func(*testing.T) {
	return func(t *testing.T) { t.Skip(message) }
}
func (this testCase) runTest(t *testing.T) {
	if this.config.parallelTests {
		t.Parallel()
	}

	fixtureValue := this.fixtureValue
	if this.config.freshFixture {
		fixtureValue = reflect.New(this.fixtureType.Elem())
	}
	fixtureValue.Elem().FieldByName("T").Set(reflect.ValueOf(New(t)))

	setup, hasSetup := fixtureValue.Interface().(setupTest)
	if hasSetup {
		setup.Setup()
	}

	teardown, hasTeardown := fixtureValue.Interface().(teardownTest)
	if hasTeardown {
		defer teardown.Teardown()
	}

	fixtureValue.MethodByName(this.name).Call(nil)
}

type (
	setupSuite    interface{ SetupSuite() }
	setupTest     interface{ Setup() }
	teardownTest  interface{ Teardown() }
	teardownSuite interface{ TeardownSuite() }

	assertion func(actual any, expected ...any) string
)

type TestingT interface {
	Helper()
	Log(...any)
	Error(...any)
}
type T struct{ TestingT }

func New(t TestingT) *T { return &T{TestingT: t} }

func (this *T) Write(p []byte) (int, error) {
	this.Helper()
	this.Log(string(p))
	return len(p), nil
}

// So is a convenience method for reporting assertion failure messages
// with the many assertion functions found in github.com/smarty/assertions/should.
// Example: this.So(actual, should.Equal, expected)
func (this *T) So(actual any, assert assertion, expected ...any) bool {
	failure := assert(actual, expected...)
	failed := len(failure) > 0
	if failed {
		this.Error(failure)
	}
	return !failed
}
func (this *T) AssertNil(v any) {
	this.Helper()
	if v != nil {
		this.Error("Expected nil, got:", v)
	}
}
func (this *T) AssertNotNil(v any) {
	this.Helper()
	if v == nil {
		this.Error("Expected non-nil, but it was.")
	}
}
func (this *T) AssertTrue(b bool) {
	this.Helper()
	if !b {
		this.Error("Expected true, got false.")
	}
}
func (this *T) AssertFalse(b bool) {
	this.Helper()
	if b {
		this.Error("Expected false, got true.")
	}
}
func (this *T) AssertEqual(a, b any) {
	this.Helper()
	if err := equal(a, b); err != nil {
		this.Error(err)
	}
}
func (this *T) AssertNotEqual(a, b any) {
	this.Helper()
	if err := equal(a, b); err == nil {
		this.Error("Provided values were equal.")
	}
}
func (this *T) AssertError(a, b error) {
	this.Helper()
	if !errors.Is(a, b) && !errors.Is(b, a) {
		this.Error(fmt.Sprintf("Provided errors are unrelated:\n"+"a: %s\n"+"b: %s", a, b))
	}
}
