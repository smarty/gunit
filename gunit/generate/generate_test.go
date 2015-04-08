package generate

import "github.com/smartystreets/gunit/gunit/parse"

// receive a list of fixtures
// produce a single string template that can be written to a file
// What about checksumming all the files in the directory? (Is that a separate piece?)

//////////////////////////////////////////////////////////////////////////////

var (
	inputs = []*parse.Fixture{
		{
			StructName: "A",
		},
		{
			StructName:    "B",
			TestCaseNames: []string{"TestB1"},
		},
		{
			StructName:    "C",
			TestSetupName: "SetupC_",
			TestCaseNames: []string{"TestC1", "TestC2"},
		},
		{
			StructName:       "D",
			TestTeardownName: "TeardownD_",
			TestCaseNames:    []string{"TestD1", "TestD2"},
		},
		{
			StructName:       "E",
			TestSetupName:    "SetupE_",
			TestTeardownName: "TeardownE_",
			TestCaseNames:    []string{"TestE1", "TestE2"},
		},
		{
			StructName:       "F",
			FixtureSetupName: "SetupF",
			TestCaseNames:    []string{"TestF1", "TestF2"},
		},
		{
			StructName:          "G",
			FixtureTeardownName: "TeardownG",
			TestCaseNames:       []string{"TestG1", "TestG2"},
		},
		{
			StructName:          "H",
			FixtureTeardownName: "TeardownH",
			TestCaseNames:       []string{"TestH1", "TestH2"},
		},
		// TODO: Skipped fixture (mixed with separate inputs)
		// TODO: Focused fixture (mixed with separate inputs)
		// TODO: Skipped and Focused together
		// TODO: Skipped test cases
		// TODO: Focused test cases
	}
)
