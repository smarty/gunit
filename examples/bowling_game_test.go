package examples

import (
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

type BowlingGameScoringTests struct {
	*gunit.Fixture

	game *Game
}

func (self *BowlingGameScoringTests) Setup() {
	self.game = NewGame()
}

func (self *BowlingGameScoringTests) TestAfterAllGutterBallsTheScoreShouldBeZero() {
	self.rollZero(20)
	self.So(self.game.Score(), should.Equal, 0)
}

func (self *BowlingGameScoringTests) TestAfterAllOnesTheScoreShouldBeTwenty() {
	self.rollMany(20, 1)
	self.So(self.game.Score(), should.Equal, 20)
}

func (self *BowlingGameScoringTests) TestSpareReceivesSingleRollBonus() {
	self.rollSpare()
	self.game.Roll(4)
	self.game.Roll(3)
	self.rollZero(16)
	self.So(self.game.Score(), should.Equal, 21)
}

func (self *BowlingGameScoringTests) TestStrikeRecievesDoubleRollBonus() {
	self.rollStrike()
	self.game.Roll(4)
	self.game.Roll(3)
	self.rollZero(16)
	self.So(self.game.Score(), should.Equal, 24)
}

func (self *BowlingGameScoringTests) TestPerfectGame() {
	self.rollMany(12, 10)
	self.So(self.game.Score(), should.Equal, 300)
}

func (self *BowlingGameScoringTests) rollZero(times int) {
	self.rollMany(times, 0)
}

func (self *BowlingGameScoringTests) rollMany(times, pins int) {
	for x := 0; x < times; x++ {
		self.game.Roll(pins)
	}
}
func (self *BowlingGameScoringTests) rollSpare() {
	self.game.Roll(5)
	self.game.Roll(5)
}
func (self *BowlingGameScoringTests) rollStrike() {
	self.game.Roll(10)
}
