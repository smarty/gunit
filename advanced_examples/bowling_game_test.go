package examples

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestBowlingGameScoringFixture(t *testing.T) {
	gunit.Run(new(BowlingGameScoringFixture), t)
}

type BowlingGameScoringFixture struct {
	*gunit.Fixture
	game *Game
}

func (this *BowlingGameScoringFixture) Setup() {
	this.game = NewGame()
}

func (this *BowlingGameScoringFixture) TestAfterAllGutterBallsTheScoreShouldBeZero() {
	this.rollMany(20, 0)
	this.So(this.game.Score(), should.Equal, 0)
}

func (this *BowlingGameScoringFixture) TestAfterAllOnesTheScoreShouldBeTwenty() {
	this.rollMany(20, 1)
	this.So(this.game.Score(), should.Equal, 20)
}

func (this *BowlingGameScoringFixture) TestSpareReceivesSingleRollBonus() {
	this.rollSpare()
	this.game.Roll(4)
	this.game.Roll(3)
	this.rollMany(16, 0)
	this.So(this.game.Score(), should.Equal, 21)
}

func (this *BowlingGameScoringFixture) TestStrikeReceivesDoubleRollBonus() {
	this.rollStrike()
	this.game.Roll(4)
	this.game.Roll(3)
	this.rollMany(16, 0)
	this.So(this.game.Score(), should.Equal, 24)
}

func (this *BowlingGameScoringFixture) TestPerfectGame() {
	this.rollMany(12, 10)
	this.So(this.game.Score(), should.Equal, 300)
}

func (this *BowlingGameScoringFixture) rollMany(times, pins int) {
	for x := 0; x < times; x++ {
		this.game.Roll(pins)
	}
}
func (this *BowlingGameScoringFixture) rollSpare() {
	this.game.Roll(5)
	this.game.Roll(5)
}
func (this *BowlingGameScoringFixture) rollStrike() {
	this.game.Roll(10)
}
