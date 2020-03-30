package examples

import (
	"testing"

	"github.com/smartystreets/gunit"
)

func TestHowToRegisterAllTestCasesAsSequential(t *testing.T) {
	gunit.Run(new(HowToRegisterAllTestCasesAsSequential), t,
		gunit.Options.SequentialTestCases(), // <-- Just pass the SequentialTestCases option to gunit.Run(...)!
	)
}

type HowToRegisterAllTestCasesAsSequential struct {
	*gunit.Fixture
}

func (this *HowToRegisterAllTestCasesAsSequential) TestTestsOnThisFixtureWillNOTBeRunInParallelWithEachOther() {
}

func (this *HowToRegisterAllTestCasesAsSequential) TestA() {}
func (this *HowToRegisterAllTestCasesAsSequential) TestB() {}
func (this *HowToRegisterAllTestCasesAsSequential) TestC() {}
