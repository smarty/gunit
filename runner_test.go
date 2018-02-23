package gunit

import (
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

/**************************************************************************/
/**************************************************************************/

func TestRunnerPanicsIfFixtureIsIncompatible(t *testing.T) {
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
	RunSequential(new(FixtureWithNoTestCases), t)
}

type FixtureWithNoTestCases struct{ *Fixture }

/**************************************************************************/
/**************************************************************************/

func TestRunnerFixtureWithSetupAndTeardown(t *testing.T) {
	invocations_A = []string{}

	defer assertSetupTeardownInvocationsInCorrectOrder(t)
	RunSequential(new(RunnerFixtureSetupTeardown), t)
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

type RunnerFixtureSetupTeardown struct{ *Fixture }

func (this *RunnerFixtureSetupTeardown) Setup()         { invocations_A = append(invocations_A, "Setup") }
func (this *RunnerFixtureSetupTeardown) Teardown()      { invocations_A = append(invocations_A, "Teardown") }
func (this *RunnerFixtureSetupTeardown) Test1()         { invocations_A = append(invocations_A, "Test1") }
func (this *RunnerFixtureSetupTeardown) SkipTest2()     { invocations_A = append(invocations_A, "Test2") }
func (this *RunnerFixtureSetupTeardown) LongTest3()     { invocations_A = append(invocations_A, "Test3") }
func (this *RunnerFixtureSetupTeardown) SkipLongTest4() { invocations_A = append(invocations_A, "Test4") }

/**************************************************************************/
/**************************************************************************/

func TestRunnerFixture(t *testing.T) {
	invocations_B = []string{}

	defer assertInvocationsInCorrectOrder(t)
	RunSequential(new(RunnerFixturePlain), t)
}
func assertInvocationsInCorrectOrder(t *testing.T) {
	expectedInvocations := []string{"Test3", "Test1"} // Test2 and Test4 are always skipped
	if testing.Short() {
		expectedInvocations = expectedInvocations[1:]
	}
	assertions.New(t).So(invocations_B, should.Resemble, expectedInvocations)
}

var invocations_B []string

type RunnerFixturePlain struct{ *Fixture }

func (this *RunnerFixturePlain) Test1()         { invocations_B = append(invocations_B, "Test1") }
func (this *RunnerFixturePlain) SkipTest2()     { invocations_B = append(invocations_B, "Test2") }
func (this *RunnerFixturePlain) LongTest3()     { invocations_B = append(invocations_B, "Test3") }
func (this *RunnerFixturePlain) SkipLongTest4() { invocations_B = append(invocations_B, "Test4") }

/**************************************************************************/
/**************************************************************************/
