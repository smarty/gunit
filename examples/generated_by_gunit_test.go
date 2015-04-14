//////////////////////////////////////////////////////////////////////////////
// Generated Code ////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

package examples

import (
	"testing"

	"github.com/smartystreets/gunit"
)

///////////////////////////////////////////////////////////////////////////////

func TestBowlingGameScoringTests(t *testing.T) {
	fixture := gunit.NewFixture(t)
	defer fixture.Finalize()

	test0 := &BowlingGameScoringTests{Fixture: fixture}
	test0.RunTestCase__(test0.TestAfterAllGutterBallsTheScoreShouldBeZero, "Test after all gutter balls the score should be zero")

	test1 := &BowlingGameScoringTests{Fixture: fixture}
	test1.RunTestCase__(test1.TestAfterAllOnesTheScoreShouldBeTwenty, "Test after all ones the score should be twenty")

	test2 := &BowlingGameScoringTests{Fixture: fixture}
	test2.RunTestCase__(test2.TestSpareReceivesSingleRollBonus, "Test spare receives single roll bonus")

	test3 := &BowlingGameScoringTests{Fixture: fixture}
	test3.RunTestCase__(test3.TestStrikeRecievesDoubleRollBonus, "Test strike recieves double roll bonus")

	test4 := &BowlingGameScoringTests{Fixture: fixture}
	test4.RunTestCase__(test4.TestPerfectGame, "Test perfect game")
}

func (self *BowlingGameScoringTests) RunTestCase__(test func(), description string) {
	self.T.Log(description)
	self.Setup()
	test()
}

///////////////////////////////////////////////////////////////////////////////

func TestEnvironmentControllerFixture(t *testing.T) {
	fixture := gunit.NewFixture(t)
	defer fixture.Finalize()

	test0 := &EnvironmentControllerFixture{Fixture: fixture}
	test0.RunTestCase__(test0.TestEverythingTurnedOffAtStartup, "Test everything turned off at startup")

	test1 := &EnvironmentControllerFixture{Fixture: fixture}
	test1.RunTestCase__(test1.TestEverythingOffWhenComfortable, "Test everything off when comfortable")

	test2 := &EnvironmentControllerFixture{Fixture: fixture}
	test2.RunTestCase__(test2.TestCoolerAndBlowerWhenHot, "Test cooler and blower when hot")

	test3 := &EnvironmentControllerFixture{Fixture: fixture}
	test3.RunTestCase__(test3.TestHeaterAndBlowerWhenCold, "Test heater and blower when cold")

	test4 := &EnvironmentControllerFixture{Fixture: fixture}
	test4.RunTestCase__(test4.TestHighAlarmOnIfAtThreshold, "Test high alarm on if at threshold")

	test5 := &EnvironmentControllerFixture{Fixture: fixture}
	test5.RunTestCase__(test5.TestLowAlarmOnIfAtThreshold, "Test low alarm on if at threshold")
}

func (self *EnvironmentControllerFixture) RunTestCase__(test func(), description string) {
	self.T.Log(description)
	self.Setup()
	test()
}

///////////////////////////////////////////////////////////////////////////////

func init() {
	gunit.Validate("9771380f5c38c24e440d11756ffb8aa5")
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////// Generated Code //
///////////////////////////////////////////////////////////////////////////////
