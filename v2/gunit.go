package gunit

import (
	"fmt"
	"reflect"
	"runtime/debug"
	"strings"
	"testing"
)

/*
Run accepts a fixture with Test* methods and
optional setup/teardown methods and executes
the suite. Fixtures must be struct types which
embeds a *gunit.Fixture. Assuming a fixture struct
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
func Run(outerFixture any, t *testing.T, options ...Option) {
	config := new(config)
	for _, option := range append(defaultOptions, options...) {
		option(config)
	}

	fixtureValue := reflect.ValueOf(outerFixture)
	fixtureType := reflect.TypeOf(outerFixture)

	var (
		testNames        []string
		skippedTestNames []string
		focusedTestNames []string
	)
	for x := range fixtureType.NumMethod() {
		name := fixtureType.Method(x).Name
		method := fixtureValue.MethodByName(name)
		_, isNiladic := method.Interface().(func())
		if !isNiladic {
			continue
		}

		if strings.HasPrefix(name, "Test") {
			testNames = append(testNames, name)
		} else if config.skipAllTests || strings.HasPrefix(name, "SkipTest") {
			skippedTestNames = append(skippedTestNames, name)
		} else if strings.HasPrefix(name, "FocusTest") {
			focusedTestNames = append(focusedTestNames, name)
		}
	}

	for _, name := range skippedTestNames {
		testCase{t: t, manualSkip: true, name: name}.run()
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

	setInnerFixture(fixtureValue, t)

	setup, hasSetup := outerFixture.(setupSuite)
	if hasSetup {
		setup.SetupSuite()
	}

	teardown, hasTeardown := outerFixture.(teardownSuite)
	if hasTeardown {
		defer teardown.TeardownSuite()
	}

	for _, name := range testNames {
		testCase{
			t:            t,
			name:         name,
			config:       config,
			fixtureType:  fixtureType,
			fixtureValue: fixtureValue,
		}.run()
	}
}

func setInnerFixture(fixtureValue reflect.Value, t *testing.T) {
	defer func() {
		if recover() != nil {
			panic("must embed a *gunit.Fixture on the provided fixture")
		}
	}()
	fixtureValue.Elem().FieldByName("Fixture").Set(reflect.ValueOf(&Fixture{TestingT: t}))
}

type testCase struct {
	t            *testing.T
	name         string
	config       *config
	manualSkip   bool
	fixtureType  reflect.Type
	fixtureValue reflect.Value
}

func (this testCase) run() {
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
	defer func() {
		if r := recover(); r != nil {
			t.Fail()
			t.Log(panicReport(r, debug.Stack()))
		}
	}()
	if this.config.parallelTests {
		t.Parallel()
	}

	fixtureValue := this.fixtureValue
	if this.config.freshFixture {
		fixtureValue = reflect.New(this.fixtureType.Elem())
	}
	setInnerFixture(fixtureValue, t)

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
)

func panicReport(r any, stack []byte) string {
	var builder strings.Builder
	_, _ = fmt.Fprintln(&builder, "PANIC:", r)
	_, _ = fmt.Fprintln(&builder, "...")

	opened, closed := false, false
	for _, line := range strings.Split(string(stack), "\n") {
		if strings.Contains(line, "/runtime/panic.go:") {
			opened = true
			continue
		}
		if !opened || closed {
			continue
		}
		if strings.Contains(line, "reflect.Value.call({0x") {
			closed = true
			continue
		}
		_, _ = fmt.Fprintln(&builder, line)
	}
	return strings.TrimSpace(builder.String())
}
