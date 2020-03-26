package examples

import (
	"testing"

	"github.com/smartystreets/gunit"
)

func TestHowToRegisterTheFixtureAndAllItsCasesAsSequential(t *testing.T) {
	gunit.Run(new(HowToRegisterTheFixtureAndAllItsCasesAsSequential), t,
		gunit.Options.AllSequential(), // <-- Just pass the AllSequential option to gunit.Run(...)!
	)
}

type HowToRegisterTheFixtureAndAllItsCasesAsSequential struct {
	*gunit.Fixture
}

func (this *HowToRegisterTheFixtureAndAllItsCasesAsSequential) TestThisFixtureAndItsTestsWillNotBeRunInParallelInAnyWay() {
}

func (this *HowToRegisterTheFixtureAndAllItsCasesAsSequential) TestA() {}
func (this *HowToRegisterTheFixtureAndAllItsCasesAsSequential) TestB() {}
func (this *HowToRegisterTheFixtureAndAllItsCasesAsSequential) TestC() {}
