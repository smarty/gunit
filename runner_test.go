package gunit

import (
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func TestRunnerPanicsIfFixtureIsIncompatible(t *testing.T) {
	type FixtureWithoutEmbeddedGunitFixture struct { Fixture string /* should be: *gunit.Fixture */}
	defer assertPanic(t)
	Run(new(FixtureWithoutEmbeddedGunitFixture), t)
}
func assertPanic(t *testing.T) {
	assertions.New(t).So(recover(), should.NotBeNil)
}

func TestMarkedAsSkippedIfNoTestCases(t *testing.T) {
	type FixtureWithNoTestCases struct { *Fixture }
	Run(new(FixtureWithNoTestCases), t)
}

func TestRunnerFixtureWithSetupAndTeardown(t *testing.T) {
	invocations = []string{}

	defer assertSetupTeardownInvocationsInCorrectOrder(t)
	Run(new(RunnerFixtureSetupTeardown), t)
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
	assertions.New(t).So(invocations, should.Resemble, expectedInvocations)
}

/**************************************************************************/

func TestRunnerFixture(t *testing.T) {
	invocations = []string{}

	defer assertInvocationsInCorrectOrder(t)
	Run(new(RunnerFixturePlain), t)
}
func assertInvocationsInCorrectOrder(t *testing.T) {
	expectedInvocations := []string{"Test3", "Test1"} // Test2 and Test4 are always skipped
	if testing.Short() {
		expectedInvocations = expectedInvocations[1:]
	}
	assertions.New(t).So(invocations, should.Resemble, expectedInvocations)
}

/**************************************************************************/

var invocations []string

type RunnerFixtureSetupTeardown struct{ *Fixture }

func (this *RunnerFixtureSetupTeardown) Setup()         { invocations = append(invocations, "Setup") }
func (this *RunnerFixtureSetupTeardown) Teardown()      { invocations = append(invocations, "Teardown") }
func (this *RunnerFixtureSetupTeardown) Test1()         { invocations = append(invocations, "Test1") }
func (this *RunnerFixtureSetupTeardown) SkipTest2()     { invocations = append(invocations, "Test2") }
func (this *RunnerFixtureSetupTeardown) LongTest3()     { invocations = append(invocations, "Test3") }
func (this *RunnerFixtureSetupTeardown) SkipLongTest4() { invocations = append(invocations, "Test4") }

type RunnerFixturePlain struct{ *Fixture }

func (this *RunnerFixturePlain) Test1()         { invocations = append(invocations, "Test1") }
func (this *RunnerFixturePlain) SkipTest2()     { invocations = append(invocations, "Test2") }
func (this *RunnerFixturePlain) LongTest3()     { invocations = append(invocations, "Test3") }
func (this *RunnerFixturePlain) SkipLongTest4() { invocations = append(invocations, "Test4") }
