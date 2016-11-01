package parse

const malformedTestCode = "This isn't really a go file."

const malformedMissingPointerOnEmbeddedStruct = `package parse

import (
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

type BowlingGameScoringTests struct {
	gunit.Fixture // It's missing the pointer asterisk! It should be: *gunit.Fixture

	game *Game
}

func (this *BowlingGameScoringTests) TestAfterAllGutterBallsTheScoreShouldBeZero() {}
`

const malformedMissingPointerOnReceiver = `package parse

import (
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

type BowlingGameScoringTests struct {
	*gunit.Fixture // It's missing the pointer asterisk! It should be: '*gunit.Fixture'

	game *Game
}

func (this BowlingGameScoringTests) TestAfterAllGutterBallsTheScoreShouldBeZero() {
	// we are missing the pointer asterisk on the reciever type. Should be: '(this *BowlingGameScoringTests)'
}
`

const comprehensiveTestCode = `package parse

import (
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

type BowlingGameScoringTests struct {
	*gunit.Fixture

	game *Game
}

func (this *BowlingGameScoringTests) SetupTheGame() {
	this.game = NewGame()
}

func (this *BowlingGameScoringTests) TeardownTheGame() {
	this.game = nil
}

func (this *BowlingGameScoringTests) TestAfterAllGutterBallsTheScoreShouldBeZero() {
	this.rollMany(20, 0)
	this.So(this.game.Score(), should.Equal, 0)
}

func (this *BowlingGameScoringTests) TestAfterAllOnesTheScoreShouldBeTwenty() {
	this.rollMany(20, 1)
	this.So(this.game.Score(), should.Equal, 20)
}

func (this *BowlingGameScoringTests) SkipTestASpareDeservesABonus()      {}

func (this *BowlingGameScoringTests) LongTestPerfectGame() {
	this.rollMany(12, 10)
	this.So(this.game.Score(), should.Equal, 300)
}

func (this *BowlingGameScoringTests) SkipLongTestPerfectGame() {
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

func (this *BowlingGameScoringTests) TestNotNiladic_ShouldNotBeCollected(a int) {
	// This should not be collected (it's not niladic)
}
func (this *BowlingGameScoringTests) TestNotVoid_ShouldNOTBeCollected() int {
	return -1
	// This should not be collected (it's not void)
}

//////////////////////////////////////////////////////////////////////////////

// Game contains the state of a bowling game.
type Game struct {
	rolls   []int
	current int
}

// NewGame allocates and starts a new game of bowling.
func NewGame() *Game {
	game := new(Game)
	game.rolls = make([]int, maxThrowsPerGame)
	return game
}

// Roll rolls the ball and knocks down the number of pins specified by pins.
func (this *Game) Roll(pins int) {
	this.rolls[this.current] = pins
	this.current++
}

// Score calculates and returns the player's current score.
func (this *Game) Score() (sum int) {
	for throw, frame := 0, 0; frame < framesPerGame; frame++ {
		if this.isStrike(throw) {
			sum += this.strikeBonusFor(throw)
			throw += 1
		} else if this.isSpare(throw) {
			sum += this.spareBonusFor(throw)
			throw += 2
		} else {
			sum += this.framePointsAt(throw)
			throw += 2
		}
	}
	return sum
}

// isStrike determines if a given throw is a strike or not. A strike is knocking
// down all pins in one throw.
func (this *Game) isStrike(throw int) bool {
	return this.rolls[throw] == allPins
}

// strikeBonusFor calculates and returns the strike bonus for a throw.
func (this *Game) strikeBonusFor(throw int) int {
	return allPins + this.framePointsAt(throw+1)
}

// isSpare determines if a given frame is a spare or not. A spare is knocking
// down all pins in one frame with two throws.
func (this *Game) isSpare(throw int) bool {
	return this.framePointsAt(throw) == allPins
}

// spareBonusFor calculates and returns the spare bonus for a throw.
func (this *Game) spareBonusFor(throw int) int {
	return allPins + this.rolls[throw+2]
}

// framePointsAt computes and returns the score in a frame specified by throw.
func (this *Game) framePointsAt(throw int) int {
	return this.rolls[throw] + this.rolls[throw+1]
}

const (
	// allPins is the number of pins allocated per fresh throw.
	allPins = 10

	// framesPerGame is the number of frames per bowling game.
	framesPerGame = 10

	// maxThrowsPerGame is the maximum number of throws possible in a single game.
	maxThrowsPerGame = 21
)

//////////////////////////////////////////////////////////////////////////////
// These types shouldn't be parsed as fixtures:

type TestFixtureWrongTestCase struct {
	*blah.Fixture
}
type TestFixtureWrongPackage struct {
	*gunit.Fixture2
}

type Hah interface {
	Hi() string
}

type BlahFixture struct {
	blah int
}

//////////////////////////////////////////////////////////////////////////////
`
