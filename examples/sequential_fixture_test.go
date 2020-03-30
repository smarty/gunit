package examples

import (
	"testing"

	"github.com/smartystreets/gunit"
)

func TestHowToRegisterASequentialFixture(t *testing.T) {
	gunit.Run(new(HowToRegisterASequentialFixture), t,
		gunit.Options.SequentialFixture(), // <-- Just pass the SequentialFixture option to gunit.Run(...)!
	)
}

type HowToRegisterASequentialFixture struct {
	*gunit.Fixture
}

func (this *HowToRegisterASequentialFixture) TestTestsOnThisFixtureWillNOTBeRunInParallelAnyOtherTests() {
}

func (this *HowToRegisterASequentialFixture) TestA() {}
func (this *HowToRegisterASequentialFixture) TestB() {}
func (this *HowToRegisterASequentialFixture) TestC() {}
