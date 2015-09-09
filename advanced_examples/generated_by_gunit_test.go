//////////////////////////////////////////////////////////////////////////////
// Generated Code ////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

package examples

import (
	"testing"

	"github.com/smartystreets/gunit"
)

///////////////////////////////////////////////////////////////////////////////

func Test_BowlingGameScoringTests__after_all_gutter_balls_the_score_should_be_zero(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose())
	defer fixture.Finalize()
	test := &BowlingGameScoringTests{Fixture: fixture}
	test.Setup()
	test.TestAfterAllGutterBallsTheScoreShouldBeZero()
}
func Test_BowlingGameScoringTests__after_all_ones_the_score_should_be_twenty(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose())
	defer fixture.Finalize()
	test := &BowlingGameScoringTests{Fixture: fixture}
	test.Setup()
	test.TestAfterAllOnesTheScoreShouldBeTwenty()
}
func Test_BowlingGameScoringTests__spare_receives_single_roll_bonus(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose())
	defer fixture.Finalize()
	test := &BowlingGameScoringTests{Fixture: fixture}
	test.Setup()
	test.TestSpareReceivesSingleRollBonus()
}
func Test_BowlingGameScoringTests__strike_recieves_double_roll_bonus(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose())
	defer fixture.Finalize()
	test := &BowlingGameScoringTests{Fixture: fixture}
	test.Setup()
	test.TestStrikeRecievesDoubleRollBonus()
}
func Test_BowlingGameScoringTests__perfect_game(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose())
	defer fixture.Finalize()
	test := &BowlingGameScoringTests{Fixture: fixture}
	test.Setup()
	test.TestPerfectGame()
}

///////////////////////////////////////////////////////////////////////////////

func Test_EnvironmentControllerFixture__everything_turned_off_at_startup(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose())
	defer fixture.Finalize()
	test := &EnvironmentControllerFixture{Fixture: fixture}
	test.Setup()
	test.TestEverythingTurnedOffAtStartup()
}
func Test_EnvironmentControllerFixture__everything_off_when_comfortable(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose())
	defer fixture.Finalize()
	test := &EnvironmentControllerFixture{Fixture: fixture}
	test.Setup()
	test.TestEverythingOffWhenComfortable()
}
func Test_EnvironmentControllerFixture__cooler_and_blower_when_hot(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose())
	defer fixture.Finalize()
	test := &EnvironmentControllerFixture{Fixture: fixture}
	test.Setup()
	test.TestCoolerAndBlowerWhenHot()
}
func Test_EnvironmentControllerFixture__heater_and_blower_when_cold(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose())
	defer fixture.Finalize()
	test := &EnvironmentControllerFixture{Fixture: fixture}
	test.Setup()
	test.TestHeaterAndBlowerWhenCold()
}
func Test_EnvironmentControllerFixture__high_alarm_on_if_at_threshold(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose())
	defer fixture.Finalize()
	test := &EnvironmentControllerFixture{Fixture: fixture}
	test.Setup()
	test.TestHighAlarmOnIfAtThreshold()
}
func Test_EnvironmentControllerFixture__low_alarm_on_if_at_threshold(t *testing.T) {
	t.Parallel()
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
