package gunit

import (
	"testing"

	"github.com/smartystreets/gunit/assertions"
)

/**************************************************************************/
/**************************************************************************/

func TestRunnerEndsFatallyIfFixtureIsIncompatible(t *testing.T) {
	t.Parallel()

	test := Setup(false)
	ensureEmbeddedFixture(new(FixtureWithoutEmbeddedGunitFixture), test.fakeT)
	assertions.New(t).AssertTrue(test.fixture.Failed())
}

type FixtureWithoutEmbeddedGunitFixture struct {
	Fixture string /* should be: *gunit.Fixture */
}

/**************************************************************************/
/**************************************************************************/

func TestMarkedAsSkippedIfNoTestCases(t *testing.T) {
	Run(new(FixtureWithNoTestCases), t)
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
	assertions.New(t).AssertDeepEqual(expectedInvocations, invocations_A)
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
	assertions.New(t).AssertDeepEqual(expectedInvocations, invocations_B)
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
	assertions.New(t).AssertDeepEqual([]string{"Test3"}, invocations_C)
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
	assertions.New(t).AssertDeepEqual(expected, invocations_D)
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
	assertions.New(t).AssertDeepEqual([]string{"Test1"}, invocations_E)
}

var invocations_E []string

type RunnerFixtureWithOnlyOneFocus struct{ *Fixture }

func (this *RunnerFixtureWithOnlyOneFocus) FocusTest1() {
	invocations_E = append(invocations_E, "Test1")
}

/**************************************************************************/
/**************************************************************************/

func TestRunnerFixtureSkipAll(t *testing.T) {
	Run(new(FixtureSkipAll), t, Options.SequentialTestCases(), Options.SkipAll())
	assertions.New(t).AssertNil(invocations_F)
}

var invocations_F []string

type FixtureSkipAll struct{ *Fixture }

func (this *FixtureSkipAll) Setup()         { invocations_F = append(invocations_F, "Setup") }
func (this *FixtureSkipAll) Teardown()      { invocations_F = append(invocations_F, "Teardown") }
func (this *FixtureSkipAll) Test1()         { invocations_F = append(invocations_F, "Test1") }
func (this *FixtureSkipAll) SkipTest2()     { invocations_F = append(invocations_F, "Test2") }
func (this *FixtureSkipAll) LongTest3()     { invocations_F = append(invocations_F, "Test3") }
func (this *FixtureSkipAll) SkipLongTest4() { invocations_F = append(invocations_F, "Test4") }

/**************************************************************************/
/**************************************************************************/

func TestRunnerFixtureLongRunning(t *testing.T) {
	Run(new(PlainFixtureLongRunning), t, Options.SequentialTestCases(), Options.LongRunning())
	assertInvocationsInCorrectOrder_LongRunning(t)
}
func assertInvocationsInCorrectOrder_LongRunning(t *testing.T) {
	expectedInvocations := []string{"Test1", "Test3"} // Test2 and Test4 are always skipped
	if testing.Short() {
		expectedInvocations = nil
	}
	assertions.New(t).AssertDeepEqual(expectedInvocations, invocations_G)
}

var invocations_G []string

type PlainFixtureLongRunning struct{ *Fixture }

func (this *PlainFixtureLongRunning) Test1()     { invocations_G = append(invocations_G, "Test1") }
func (this *PlainFixtureLongRunning) SkipTest2() { invocations_G = append(invocations_G, "Test2") }
func (this *PlainFixtureLongRunning) Test3()     { invocations_G = append(invocations_G, "Test3") }
func (this *PlainFixtureLongRunning) SkipTest4() { invocations_G = append(invocations_G, "Test4") }

/**************************************************************************/
/**************************************************************************/
