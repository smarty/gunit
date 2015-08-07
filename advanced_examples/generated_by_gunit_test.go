//////////////////////////////////////////////////////////////////////////////
// Generated Code ////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

package examples

import (
	"os"
	"testing"

	"github.com/smartystreets/gunit"
)

///////////////////////////////////////////////////////////////////////////////

func TestBowlingGameScoringTests_TestAfterAllGutterBallsTheScoreShouldBeZero(t *testing.T) {
	fixture := gunit.NewFixture(t, os.Stdout, testing.Verbose())
	test := &BowlingGameScoringTests{Fixture: fixture}
	test.Setup()
	test.TestAfterAllGutterBallsTheScoreShouldBeZero()
}

func TestBowlingGameScoringTests_TestAfterAllOnesTheScoreShouldBeTwenty(t *testing.T) {
	fixture := gunit.NewFixture(t, os.Stdout, testing.Verbose())
	test := &BowlingGameScoringTests{Fixture: fixture}
	test.Setup()
	test.TestAfterAllOnesTheScoreShouldBeTwenty()
}

func TestBowlingGameScoringTests_TestSpareReceivesSingleRollBonus(t *testing.T) {
	fixture := gunit.NewFixture(t, os.Stdout, testing.Verbose())
	test := &BowlingGameScoringTests{Fixture: fixture}
	test.Setup()
	test.TestSpareReceivesSingleRollBonus()
}

func TestBowlingGameScoringTests_TestStrikeRecievesDoubleRollBonus(t *testing.T) {
	fixture := gunit.NewFixture(t, os.Stdout, testing.Verbose())
	test := &BowlingGameScoringTests{Fixture: fixture}
	test.Setup()
	test.TestStrikeRecievesDoubleRollBonus()
}

func TestBowlingGameScoringTests_TestPerfectGame(t *testing.T) {
	fixture := gunit.NewFixture(t, os.Stdout, testing.Verbose())
	test := &BowlingGameScoringTests{Fixture: fixture}
	test.Setup()
	test.TestPerfectGame()
}

///////////////////////////////////////////////////////////////////////////////

func TestEnvironmentControllerFixture_TestEverythingTurnedOffAtStartup(t *testing.T) {
	fixture := gunit.NewFixture(t, os.Stdout, testing.Verbose())
	test := &EnvironmentControllerFixture{Fixture: fixture}
	test.Setup()
	test.TestEverythingTurnedOffAtStartup()
}

func TestEnvironmentControllerFixture_TestEverythingOffWhenComfortable(t *testing.T) {
	fixture := gunit.NewFixture(t, os.Stdout, testing.Verbose())
	test := &EnvironmentControllerFixture{Fixture: fixture}
	test.Setup()
	test.TestEverythingOffWhenComfortable()
}

func TestEnvironmentControllerFixture_TestCoolerAndBlowerWhenHot(t *testing.T) {
	fixture := gunit.NewFixture(t, os.Stdout, testing.Verbose())
	test := &EnvironmentControllerFixture{Fixture: fixture}
	test.Setup()
	test.TestCoolerAndBlowerWhenHot()
}

func TestEnvironmentControllerFixture_TestHeaterAndBlowerWhenCold(t *testing.T) {
	fixture := gunit.NewFixture(t, os.Stdout, testing.Verbose())
	test := &EnvironmentControllerFixture{Fixture: fixture}
	test.Setup()
	test.TestHeaterAndBlowerWhenCold()
}

func TestEnvironmentControllerFixture_TestHighAlarmOnIfAtThreshold(t *testing.T) {
	fixture := gunit.NewFixture(t, os.Stdout, testing.Verbose())
	test := &EnvironmentControllerFixture{Fixture: fixture}
	test.Setup()
	test.TestHighAlarmOnIfAtThreshold()
}

func TestEnvironmentControllerFixture_TestLowAlarmOnIfAtThreshold(t *testing.T) {
	fixture := gunit.NewFixture(t, os.Stdout, testing.Verbose())
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
