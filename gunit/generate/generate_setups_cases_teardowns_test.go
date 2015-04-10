package generate

import (
	"strings"
	"testing"

	"github.com/smartystreets/gunit/gunit/parse"
)

func TestGenerateSetupsAndCasesAndTeardowns(t *testing.T) {
	actual := TestFile("blah", fixtures)
	if actual != expected {
		t.Errorf("\nExpected: [%s]\nActual:   [%s]", singleLine(expected), singleLine(actual))
	}
}

func singleLine(value string) string {
	return strings.Replace(value, "\n", "\\n", -1)
}

//////////////////////////////////////////////////////////////////////////////

var (
	fixtures = []*parse.Fixture{
		{
			StructName: "B",
			TestCases:  []parse.TestCase{{Name: "TestB1", StructName: "B"}},
		},
		// {
		// 	StructName:    "C",
		// 	TestSetupName: "SetupC_",
		// 	TestCaseNames: []string{"TestC1", "TestC2"},
		// },
		// {
		// 	StructName:       "D",
		// 	TestTeardownName: "TeardownD_",
		// 	TestCaseNames:    []string{"TestD1", "TestD2"},
		// },
		// {
		// 	StructName:       "E",
		// 	TestSetupName:    "SetupE_",
		// 	TestTeardownName: "TeardownE_",
		// 	TestCaseNames:    []string{"TestE1", "TestE2"},
		// },
		// {
		// 	StructName:       "F",
		// 	FixtureSetupName: "SetupF",
		// 	TestCaseNames:    []string{"TestF1", "TestF2"},
		// },
		// {
		// 	StructName:          "G",
		// 	FixtureTeardownName: "TeardownG",
		// 	TestCaseNames:       []string{"TestG1", "TestG2"},
		// },
		// {
		// 	StructName:          "H",
		// 	FixtureTeardownName: "TeardownH",
		// 	TestCaseNames:       []string{"TestH1", "TestH2"},
		// },
		// TODO: empty fixture (no defined test cases)
		// TODO: Skipped fixture (mixed with separate inputs)
		// TODO: Focused fixture (mixed with separate inputs)
		// TODO: Skipped and Focused together
		// TODO: Skipped test cases
		// TODO: Focused test cases
	}
)

const expected = `//////////////////////////////////////////////////////////////////////////////
// Generated Code ////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

package blah

import (
	"testing"

	"github.com/smartystreets/gunit"
)

//////////////////////////////////////////////////////////////////////////////

func TestB(t *testing.T) {
	fixture := gunit.NewFixture(t)
	defer fixture.Finalize()

	test0 := &B{Fixture: fixture}
	test0.RunTestCase(test0.TestB1, "Test b1")
}

func (self *B) RunTestCase(test func(), description string) {
	self.T.Log(description)
	test()
}

//////////////////////////////////////////////////////////////////////////////
`
