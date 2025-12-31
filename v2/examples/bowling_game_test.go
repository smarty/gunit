package examples

import (
	"slices"
	"testing"

	"github.com/smarty/gunit/v2"
	"github.com/smarty/gunit/v2/assert/should"
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
	this.So(this.game.CalculateScore(), should.Equal, expected)
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

func (this *BowlingGameScoringFixture) TestTable() {
	subTests := []struct {
		name     string
		rolls    []int
		expected int
	}{
		{name: "gutter-game", rolls: slices.Repeat([]int{0}, 20), expected: 0},
		{name: "all-ones", rolls: slices.Repeat([]int{1}, 20), expected: 20},
		{name: "spare", rolls: []int{5, 5, 4, 3}, expected: 21},
		{name: "strike", rolls: []int{10, 3, 4}, expected: 24},
		{name: "perfection", rolls: slices.Repeat([]int{10}, 12), expected: 300},
	}
	for _, sub := range subTests {
		this.Run(sub.name, func(fixture *gunit.Fixture) {
			game := NewGame()
			for _, roll := range sub.rolls {
				game.RecordRoll(roll)
			}
			fixture.So(game.CalculateScore(), should.Equal, sub.expected)
		})
	}
}
