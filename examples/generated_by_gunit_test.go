///////////////////////////////////////////////////////////////////////////////
// Generated Code /////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

package examples

import (
	"testing"

	"github.com/smartystreets/gunit"
)

func TestBowlingGameScoring(t *testing.T) {
	fixture := gunit.NewFixture(t)
	defer fixture.Finalize()

	test0 := &BowlingGameScoringTests{Fixture: fixture}
	test0.RunTestCase(
		test0.TestAfterAllGutterBallsTheScoreShouldBeZero,
		"After All Gutter Balls The Score Should Be Zero")

	test1 := &BowlingGameScoringTests{Fixture: fixture}
	test1.RunTestCase(
		test1.TestAfterAllOnesTheScoreShouldBeTwenty,
		"After All Ones The Score Should Be Twenty")
}

func (self *BowlingGameScoringTests) RunTestCase(test func(), description string) {
	self.T.Log(description)
	self.Setup()
	test()
}

///////////////////////////////////////////////////////////////////////////////

func TestEnvironmentController(t *testing.T) {
	test0 := &EnvironmentControllerFixture{Fixture: gunit.NewFixture(t)}
	test0.RunTestCase(
		test0.TestShouldStartWithEverythingDeactivated,
		"Should start with everything deactivated")

	test1 := &EnvironmentControllerFixture{Fixture: gunit.NewFixture(t)}
	test1.RunTestCase(
		test1.TestNoHardwareComponentsAreActivatedWhenTemperatureIsJustRight,
		"No Hardware Components Are Activated When Temperature Is Just Right")

	test2 := &EnvironmentControllerFixture{Fixture: gunit.NewFixture(t)}
	test2.RunTestCase(
		test2.TestCoolerAndBlowerActivatedWhenTemperatureIsTooHot,
		"Test Cooler And Blower Activated When Temperature Is Too Hot")

	test3 := &EnvironmentControllerFixture{Fixture: gunit.NewFixture(t)}
	test3.RunTestCase(
		test2.TestHeaterAndBlowerActivatedWhenTemperatureIsTooCold,
		"Test Heater And Blower Activated When Temperature Is Too Cold")

}

func (self *EnvironmentControllerFixture) RunTestCase(test func(), description string) {
	self.T.Log(description)
	self.Setup()
	test()
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////// Generated Code //
///////////////////////////////////////////////////////////////////////////////
