package examples

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
