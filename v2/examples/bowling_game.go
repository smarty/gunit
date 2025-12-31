package examples

type Game struct {
	rolls [maxThrowsPerGame]int
	roll  int
	score int
}

func NewGame() *Game {
	return new(Game)
}

func (this *Game) RecordRoll(pins int) {
	this.rolls[this.roll] = pins
	this.roll++
}

func (this *Game) CalculateScore() int {
	this.roll = 0
	for frame := 0; frame < framesPerGame; frame++ {
		this.score += this.calculateFrameScore()
		this.roll += this.advanceFrame()
	}
	return this.score
}

func (this *Game) calculateFrameScore() int {
	if this.isStrike() {
		return allPins + this.at(1) + this.at(2)
	}
	if this.isSpare() {
		return allPins + this.at(2)
	}
	return this.at(0) + this.at(1)
}
func (this *Game) advanceFrame() int {
	if this.isStrike() {
		return 1
	}
	return 2
}

func (this *Game) isStrike() bool    { return this.at(0) == allPins }
func (this *Game) isSpare() bool     { return this.at(0)+this.at(1) == allPins }
func (this *Game) at(offset int) int { return this.rolls[this.roll+offset] }

const (
	allPins          = 10
	framesPerGame    = 10
	maxThrowsPerGame = 21
)
