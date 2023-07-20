package examples

import (
	"testing"

	"github.com/smartystreets/gunit"
)

func TestHowToSkipAnEntireFixture(t *testing.T) {
	gunit.Run(new(HowToSkipAnEntireFixture), t,
		gunit.Options.SkipAll(), // <-- Just pass the SkipAll option to gunit.Run(...)!
	)
}

type HowToSkipAnEntireFixture struct {
	*gunit.Fixture
}

func (this *HowToSkipAnEntireFixture) TestAllTestMethodsWillBeSkipped_ThankGoodness() {
	this.Assert(false)
}

func (this *HowToSkipAnEntireFixture) TestA() { this.Assert(false) }
func (this *HowToSkipAnEntireFixture) TestB() { this.Assert(false) }
func (this *HowToSkipAnEntireFixture) TestC() { this.Assert(false) }
