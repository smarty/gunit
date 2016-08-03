//////////////////////////////////////////////////////////////////////////////
// Generated Code ////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

package examples

import (
	"testing"

	"github.com/smartystreets/gunit"
)

///////////////////////////////////////////////////////////////////////////////

func Test_BowlingGameScoringTests__TestAfterAllGutterBallsTheScoreShouldBeZero(t *testing.T) {
	fixture := gunit.NewFixture(t, testing.Verbose())
	defer fixture.Finalize()
	test := &BowlingGameScoringTests{Fixture: fixture}
	test.Setup()
	test.TestAfterAllGutterBallsTheScoreShouldBeZero()
}

func Test_BowlingGameScoringTests__TestAfterAllOnesTheScoreShouldBeTwenty(t *testing.T) {
	fixture := gunit.NewFixture(t, testing.Verbose())
	defer fixture.Finalize()
	test := &BowlingGameScoringTests{Fixture: fixture}
	test.Setup()
	test.TestAfterAllOnesTheScoreShouldBeTwenty()
}

func Test_BowlingGameScoringTests__TestSpareReceivesSingleRollBonus(t *testing.T) {
	fixture := gunit.NewFixture(t, testing.Verbose())
	defer fixture.Finalize()
	test := &BowlingGameScoringTests{Fixture: fixture}
	test.Setup()
	test.TestSpareReceivesSingleRollBonus()
}

func Test_BowlingGameScoringTests__TestStrikeRecievesDoubleRollBonus(t *testing.T) {
	fixture := gunit.NewFixture(t, testing.Verbose())
	defer fixture.Finalize()
	test := &BowlingGameScoringTests{Fixture: fixture}
	test.Setup()
	test.TestStrikeRecievesDoubleRollBonus()
}

func Test_BowlingGameScoringTests__TestPerfectGame(t *testing.T) {
	fixture := gunit.NewFixture(t, testing.Verbose())
	defer fixture.Finalize()
	test := &BowlingGameScoringTests{Fixture: fixture}
	test.Setup()
	test.TestPerfectGame()
}

///////////////////////////////////////////////////////////////////////////////

func Test_EnvironmentControllerFixture__TestEverythingTurnedOffAtStartup(t *testing.T) {
	fixture := gunit.NewFixture(t, testing.Verbose())
	defer fixture.Finalize()
	test := &EnvironmentControllerFixture{Fixture: fixture}
	test.Setup()
	test.TestEverythingTurnedOffAtStartup()
}

func Test_EnvironmentControllerFixture__TestEverythingOffWhenComfortable(t *testing.T) {
	fixture := gunit.NewFixture(t, testing.Verbose())
	defer fixture.Finalize()
	test := &EnvironmentControllerFixture{Fixture: fixture}
	test.Setup()
	test.TestEverythingOffWhenComfortable()
}

func Test_EnvironmentControllerFixture__TestCoolerAndBlowerWhenHot(t *testing.T) {
	fixture := gunit.NewFixture(t, testing.Verbose())
	defer fixture.Finalize()
	test := &EnvironmentControllerFixture{Fixture: fixture}
	test.Setup()
	test.TestCoolerAndBlowerWhenHot()
}

func Test_EnvironmentControllerFixture__TestHeaterAndBlowerWhenCold(t *testing.T) {
	fixture := gunit.NewFixture(t, testing.Verbose())
	defer fixture.Finalize()
	test := &EnvironmentControllerFixture{Fixture: fixture}
	test.Setup()
	test.TestHeaterAndBlowerWhenCold()
}

func Test_EnvironmentControllerFixture__TestHighAlarmOnIfAtThreshold(t *testing.T) {
	fixture := gunit.NewFixture(t, testing.Verbose())
	defer fixture.Finalize()
	test := &EnvironmentControllerFixture{Fixture: fixture}
	test.Setup()
	test.TestHighAlarmOnIfAtThreshold()
}

func Test_EnvironmentControllerFixture__TestLowAlarmOnIfAtThreshold(t *testing.T) {
	fixture := gunit.NewFixture(t, testing.Verbose())
	defer fixture.Finalize()
	test := &EnvironmentControllerFixture{Fixture: fixture}
	test.Setup()
	test.TestLowAlarmOnIfAtThreshold()
}

///////////////////////////////////////////////////////////////////////////////

func init() {
	gunit.Validate("0a085d3a3201a50b682d18a273aa68de")
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////// Generated Code //
///////////////////////////////////////////////////////////////////////////////
