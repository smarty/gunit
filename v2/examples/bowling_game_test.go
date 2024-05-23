package examples

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/smarty/gunit/v2"
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

func (this *BowlingGameScoringFixture) TestGutterGame() {
	this.rollMany(20, 0)
	this.assertScore(0)
}
func (this *BowlingGameScoringFixture) TestAllOnes() {
	this.rollMany(20, 1)
	this.assertScore(20)
}
func (this *BowlingGameScoringFixture) TestSpare() {
	this.rollSeveral(5, 5, 4, 3)
	this.assertScore(21)
}
func (this *BowlingGameScoringFixture) TestStrike() {
	this.rollSeveral(10, 3, 4)
	this.assertScore(24)
}
func (this *BowlingGameScoringFixture) Test21Throws() {
	this.rollMany(18, 0)
	this.rollSeveral(5, 5, 5)
	this.assertScore(15)
}
func (this *BowlingGameScoringFixture) TestPerfection() {
	this.rollMany(12, 10)
	this.assertScore(300)
}

func (this *BowlingGameScoringFixture) assertScore(expected int) {
	this.So(expected, shouldEqual, this.game.CalculateScore())
}

func (this *BowlingGameScoringFixture) rollMany(times, pins int) {
	for x := 0; x < times; x++ {
		this.game.RecordRoll(pins)
	}
}
func (this *BowlingGameScoringFixture) rollSeveral(rolls ...int) {
	for _, roll := range rolls {
		this.game.RecordRoll(roll)
	}
}

// TODO: use should.Equal (once it is defined)
func shouldEqual(actual any, expected ...any) error {
	if reflect.DeepEqual(actual, expected[0]) {
		return nil
	}
	return fmt.Errorf("shouldEqual failed: %v vs %v", actual, expected[0])
}
