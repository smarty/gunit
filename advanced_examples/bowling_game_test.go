package examples

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestBowlingGameScoring(t *testing.T) {
	gunit.Run(new(BowlingGameScoringTests), t)
}

type BowlingGameScoringTests struct {
	*gunit.Fixture

	game *Game
}

func (this *BowlingGameScoringTests) Setup() {
	this.game = NewGame()
}

func (this *BowlingGameScoringTests) TestAfterAllGutterBallsTheScoreShouldBeZero() {
	this.rollMany(20, 0)
	this.So(this.game.Score(), should.Equal, 0)
}

func (this *BowlingGameScoringTests) TestAfterAllOnesTheScoreShouldBeTwenty() {
	this.rollMany(20, 1)
	this.So(this.game.Score(), should.Equal, 20)
}

func (this *BowlingGameScoringTests) TestSpareReceivesSingleRollBonus() {
	this.rollSpare()
	this.game.Roll(4)
	this.game.Roll(3)
	this.rollMany(16, 0)
	this.So(this.game.Score(), should.Equal, 21)
}

func (this *BowlingGameScoringTests) TestStrikeReceivesDoubleRollBonus() {
	this.rollStrike()
	this.game.Roll(4)
	this.game.Roll(3)
	this.rollMany(16, 0)
	this.So(this.game.Score(), should.Equal, 24)
}

func (this *BowlingGameScoringTests) TestPerfectGame() {
	this.rollMany(12, 10)
	this.So(this.game.Score(), should.Equal, 300)
}

func (this *BowlingGameScoringTests) rollMany(times, pins int) {
	for x := 0; x < times; x++ {
		this.game.Roll(pins)
	}
}
func (this *BowlingGameScoringTests) rollSpare() {
	this.game.Roll(5)
	this.game.Roll(5)
}
func (this *BowlingGameScoringTests) rollStrike() {
	this.game.Roll(10)
}
