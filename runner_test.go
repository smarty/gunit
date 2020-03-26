package gunit

import (
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

/**************************************************************************/
/**************************************************************************/

func TestRunnerEndsFatallyIfFixtureIsIncompatible(t *testing.T) {
	t.Parallel()

	test := Setup(false)
	ensureEmbeddedFixture(new(FixtureWithoutEmbeddedGunitFixture), test.fakeT)
	assertions.New(t).So(test.fixture.Failed(), should.BeTrue)
}

type FixtureWithoutEmbeddedGunitFixture struct {
	Fixture string /* should be: *gunit.Fixture */
}

/**************************************************************************/
/**************************************************************************/

func TestMarkedAsSkippedIfNoTestCases(t *testing.T) {
	Run(new(FixtureWithNoTestCases), t, Options.SequentialTestCases())
}

type FixtureWithNoTestCases struct{ *Fixture }

/**************************************************************************/
/**************************************************************************/

func TestRunnerFixtureWithSetupAndTeardown(t *testing.T) {
	Run(new(FixtureWithSetupTeardown), t, Options.SequentialTestCases())
	assertSetupTeardownInvocationsInCorrectOrder(t)
}
func assertSetupTeardownInvocationsInCorrectOrder(t *testing.T) {
	expectedInvocations := []string{
		"Setup", "Test3", "Teardown",
		"Setup", "Test1", "Teardown",
		// Test2 and Test4 are always skipped
	}
	if testing.Short() {
		expectedInvocations = expectedInvocations[3:]
	}
	assertions.New(t).So(invocations_A, should.Resemble, expectedInvocations)
}

var invocations_A []string

type FixtureWithSetupTeardown struct{ *Fixture }

func (this *FixtureWithSetupTeardown) Setup()         { invocations_A = append(invocations_A, "Setup") }
func (this *FixtureWithSetupTeardown) Teardown()      { invocations_A = append(invocations_A, "Teardown") }
func (this *FixtureWithSetupTeardown) Test1()         { invocations_A = append(invocations_A, "Test1") }
func (this *FixtureWithSetupTeardown) SkipTest2()     { invocations_A = append(invocations_A, "Test2") }
func (this *FixtureWithSetupTeardown) LongTest3()     { invocations_A = append(invocations_A, "Test3") }
func (this *FixtureWithSetupTeardown) SkipLongTest4() { invocations_A = append(invocations_A, "Test4") }

/**************************************************************************/
/**************************************************************************/

func TestRunnerFixture(t *testing.T) {
	Run(new(PlainFixture), t, Options.SequentialTestCases())
	assertInvocationsInCorrectOrder(t)
}
func assertInvocationsInCorrectOrder(t *testing.T) {
	expectedInvocations := []string{"Test3", "Test1"} // Test2 and Test4 are always skipped
	if testing.Short() {
		expectedInvocations = expectedInvocations[1:]
	}
	assertions.New(t).So(invocations_B, should.Resemble, expectedInvocations)
}

var invocations_B []string

type PlainFixture struct{ *Fixture }

func (this *PlainFixture) Test1()         { invocations_B = append(invocations_B, "Test1") }
func (this *PlainFixture) SkipTest2()     { invocations_B = append(invocations_B, "Test2") }
func (this *PlainFixture) LongTest3()     { invocations_B = append(invocations_B, "Test3") }
func (this *PlainFixture) SkipLongTest4() { invocations_B = append(invocations_B, "Test4") }

/**************************************************************************/
/**************************************************************************/

func TestRunnerFixtureWithFocus(t *testing.T) {
	Run(new(FixtureWithFocus), t, Options.SequentialTestCases())
	assertFocusIsOnlyInvocation(t)
}
func assertFocusIsOnlyInvocation(t *testing.T) {
	assertions.New(t).So(invocations_C, should.Resemble, []string{"Test3"})
}

var invocations_C []string

type FixtureWithFocus struct{ *Fixture }

func (this *FixtureWithFocus) Test1()      { invocations_C = append(invocations_C, "Test1") }
func (this *FixtureWithFocus) Test2()      { invocations_C = append(invocations_C, "Test2") }
func (this *FixtureWithFocus) FocusTest3() { invocations_C = append(invocations_C, "Test3") }
func (this *FixtureWithFocus) Test4()      { invocations_C = append(invocations_C, "Test4") }

/**************************************************************************/
/**************************************************************************/

func TestRunnerFixtureWithFocusLong(t *testing.T) {
	Run(new(FixtureWithFocusLong), t, Options.SequentialTestCases())
	assertFocusLongIsOnlyInvocation(t)
}
func assertFocusLongIsOnlyInvocation(t *testing.T) {
	expected := []string{"Test3"}
	if testing.Short() {
		expected = nil
	}
	assertions.New(t).So(invocations_D, should.Resemble, expected)
}

var invocations_D []string

type FixtureWithFocusLong struct{ *Fixture }

func (this *FixtureWithFocusLong) Test1()          { invocations_D = append(invocations_D, "Test1") }
func (this *FixtureWithFocusLong) Test2()          { invocations_D = append(invocations_D, "Test2") }
func (this *FixtureWithFocusLong) FocusLongTest3() { invocations_D = append(invocations_D, "Test3") }
func (this *FixtureWithFocusLong) Test4()          { invocations_D = append(invocations_D, "Test4") }

/**************************************************************************/
/**************************************************************************/

func TestRunnerFixtureWithOnlyOneFocus(t *testing.T) {
	Run(new(RunnerFixtureWithOnlyOneFocus), t, Options.SequentialTestCases())
	assertSingleFocusIsOnlyInvocation(t)
}
func assertSingleFocusIsOnlyInvocation(t *testing.T) {
	assertions.New(t).So(invocations_E, should.Resemble, []string{"Test1"})
}

var invocations_E []string

type RunnerFixtureWithOnlyOneFocus struct{ *Fixture }

func (this *RunnerFixtureWithOnlyOneFocus) FocusTest1() {
	invocations_E = append(invocations_E, "Test1")
}

/**************************************************************************/
/**************************************************************************/
