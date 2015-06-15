package generate

import (
	"strings"
	"testing"

	"github.com/smartystreets/gunit/gunit/parse"
)

func TestGenerateTestFileFails(t *testing.T) {
	fixtures := []*parse.Fixture{{StructName: "Not a valid struct name (spaces and parens!)"}}
	file, err := TestFile("blah", fixtures, "")
	if err == nil {
		t.Error("Expected a generate error, got nil instead.")
	}
	if len(file) > 0 {
		t.Error("Expected no generated content, got:", file)
	}
}

func TestGenerateWithoutPackageNameFails(t *testing.T) {
	fixtures := []*parse.Fixture{{StructName: "A"}}
	file, err := TestFile("", fixtures, "")
	if err == nil {
		t.Error("Expected a generate error, got nil instead.")
	}
	if len(file) > 0 {
		t.Error("Expected no generated content, got:", file)
	}
}

func TestGenerateValidTestFile(t *testing.T) {
	fixtures := []*parse.Fixture{{StructName: "A"}}
	file, err := TestFile("blah", fixtures, "42")
	if err != nil {
		t.Error("Unexpected err:", err)
	}
	if string(file) != expectedFileOutput {
		t.Errorf("Expected:\n%s\n\nActual:\n%s", expectedFileOutput, file)
	}
}

const expectedFileOutput = `//////////////////////////////////////////////////////////////////////////////
// Generated Code ////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

package blah

import (
	"os"
	"testing"

	"github.com/smartystreets/gunit"
)

///////////////////////////////////////////////////////////////////////////////

func TestA(t *testing.T) {
	fixture := gunit.NewFixture(t, os.Stdout, testing.Verbose())
	defer fixture.Finalize()

	fixture.Skip("Fixture 'A' has no test cases.")
}

///////////////////////////////////////////////////////////////////////////////

func init() {
	gunit.Validate("42")
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////// Generated Code //
///////////////////////////////////////////////////////////////////////////////
`

//////////////////////////////////////////////////////////////////////////////

func TestGenerateTestFunction(t *testing.T) {
	for i, test := range testFunction_TestCases {
		function, err := TestFunction(test.input)
		if err == nil && test.err {
			t.Error("Expected a parse error but got nil.")
			continue
		} else if err != nil && test.err {
			t.Log("✔ " + test.description)
			continue
		}

		actual := strings.TrimSpace(string(function))
		expected := strings.TrimSpace(test.expected)
		if actual != expected {
			t.Errorf("FAILED: Case #%d\nExpected:\n%s\n\nActual:\n%s", i, expected, actual)
		} else {
			t.Log("✔ " + test.description)
		}
	}
}

type TestFunction_TestCase struct {
	input       *parse.Fixture
	expected    string
	err         bool
	description string
	SKIP        bool
}

var testFunction_TestCases = []TestFunction_TestCase{

	/////////////////////////////////////////////////////////////////////////////////////////////

	{
		input: &parse.Fixture{
			StructName: "Not a valid struct name (see the spaces and parens?)",
		},
		expected:    ``,
		err:         true,
		description: "Wonky fixture should cause error (but I'm not sure how this could ever happen)",
	},

	/////////////////////////////////////////////////////////////////////////////////////////////

	{
		input: &parse.Fixture{
			StructName: "A",
		},
		expected: `

func TestA(t *testing.T) {
	fixture := gunit.NewFixture(t, os.Stdout, testing.Verbose())
	defer fixture.Finalize()

	fixture.Skip("Fixture 'A' has no test cases.")
}

`,
		description: "single fixture with no test cases",
	},

	/////////////////////////////////////////////////////////////////////////////////////////////

	{
		input: &parse.Fixture{
			StructName: "B",
			TestCases:  []parse.TestCase{{Index: 0, Name: "TestB1", StructName: "B"}},
		},
		expected: `

func TestB(t *testing.T) {
	fixture := gunit.NewFixture(t, os.Stdout, testing.Verbose())
	defer fixture.Finalize()

	test0 := &B{Fixture: fixture}
	test0.RunTestCase__(test0.TestB1, "B 1")
}

func (self *B) RunTestCase__(test func(), description string) {
	self.Describe(description)
	test()
}

`,
		description: "A fixture with a single test case, no setups, no teardowns, no skips",
	},

	/////////////////////////////////////////////////////////////////////////////////////////////

	{
		input: &parse.Fixture{
			StructName:    "C",
			TestSetupName: "SetupC_",
			TestCases: []parse.TestCase{
				{Index: 0, Name: "TestC1", StructName: "C"},
				{Index: 1, Name: "TestC2", StructName: "C"},
			},
		},
		expected: `

func TestC(t *testing.T) {
	fixture := gunit.NewFixture(t, os.Stdout, testing.Verbose())
	defer fixture.Finalize()

	test0 := &C{Fixture: fixture}
	test0.RunTestCase__(test0.TestC1, "C 1")

	test1 := &C{Fixture: fixture}
	test1.RunTestCase__(test1.TestC2, "C 2")
}

func (self *C) RunTestCase__(test func(), description string) {
	self.Describe(description)
	self.SetupC_()
	test()
}
`,
		description: "A fixture with two test cases and a setup",
	},
	/////////////////////////////////////////////////////////////////////////////////////////////

	{
		input: &parse.Fixture{
			StructName:       "D",
			TestTeardownName: "TeardownD_",
			TestCases: []parse.TestCase{
				{Index: 0, Name: "TestD1", StructName: "D"},
				{Index: 1, Name: "TestD2", StructName: "D"},
			},
		},
		expected: `

func TestD(t *testing.T) {
	fixture := gunit.NewFixture(t, os.Stdout, testing.Verbose())
	defer fixture.Finalize()

	test0 := &D{Fixture: fixture}
	test0.RunTestCase__(test0.TestD1, "D 1")

	test1 := &D{Fixture: fixture}
	test1.RunTestCase__(test1.TestD2, "D 2")
}

func (self *D) RunTestCase__(test func(), description string) {
	self.Describe(description)
	defer self.TeardownD_()
	test()
}
`,
		description: "A fixture with a two test cases and a teardown",
	},

	/////////////////////////////////////////////////////////////////////////////////////////////

	{
		input: &parse.Fixture{
			StructName:       "E",
			TestSetupName:    "SetupE_",
			TestTeardownName: "TeardownE_",
			TestCases: []parse.TestCase{
				{Index: 0, Name: "TestE1", StructName: "E"},
				{Index: 1, Name: "TestE2", StructName: "E"},
			},
		},
		expected: `

func TestE(t *testing.T) {
	fixture := gunit.NewFixture(t, os.Stdout, testing.Verbose())
	defer fixture.Finalize()

	test0 := &E{Fixture: fixture}
	test0.RunTestCase__(test0.TestE1, "E 1")

	test1 := &E{Fixture: fixture}
	test1.RunTestCase__(test1.TestE2, "E 2")
}

func (self *E) RunTestCase__(test func(), description string) {
	self.Describe(description)
	defer self.TeardownE_()
	self.SetupE_()
	test()
}
`,
		description: "A fixture with two test cases, a setup and a teardown",
	},

	/////////////////////////////////////////////////////////////////////////////////////////////

	{
		input: &parse.Fixture{
			StructName:       "F",
			FixtureSetupName: "SetupF",
			TestCases: []parse.TestCase{
				{Index: 0, Name: "TestF1", StructName: "F"},
				{Index: 1, Name: "TestF2", StructName: "F"},
			},
		},
		expected: `

func TestF(t *testing.T) {
	SetupF()

	fixture := gunit.NewFixture(t, os.Stdout, testing.Verbose())
	defer fixture.Finalize()

	test0 := &F{Fixture: fixture}
	test0.RunTestCase__(test0.TestF1, "F 1")

	test1 := &F{Fixture: fixture}
	test1.RunTestCase__(test1.TestF2, "F 2")
}

func (self *F) RunTestCase__(test func(), description string) {
	self.Describe(description)
	test()
}
`,
		description: "One-time fixture setup",
	},

	/////////////////////////////////////////////////////////////////////////////////////////////

	{
		input: &parse.Fixture{
			StructName:          "G",
			FixtureTeardownName: "TeardownG",
			TestCases: []parse.TestCase{
				{Index: 0, Name: "TestG1", StructName: "G"},
				{Index: 1, Name: "TestG2", StructName: "G"},
			},
		},
		expected: `

func TestG(t *testing.T) {
	defer TeardownG()

	fixture := gunit.NewFixture(t, os.Stdout, testing.Verbose())
	defer fixture.Finalize()

	test0 := &G{Fixture: fixture}
	test0.RunTestCase__(test0.TestG1, "G 1")

	test1 := &G{Fixture: fixture}
	test1.RunTestCase__(test1.TestG2, "G 2")
}

func (self *G) RunTestCase__(test func(), description string) {
	self.Describe(description)
	test()
}
`,
		description: "One-time fixture teardown",
	},

	/////////////////////////////////////////////////////////////////////////////////////////////

	{
		input: &parse.Fixture{
			StructName:          "H",
			FixtureSetupName:    "SetupH",
			FixtureTeardownName: "TeardownH",
			TestCases: []parse.TestCase{
				{Index: 0, Name: "TestH1", StructName: "H"},
				{Index: 1, Name: "TestH2", StructName: "H"},
			},
		},
		expected: `

func TestH(t *testing.T) {
	defer TeardownH()
	SetupH()

	fixture := gunit.NewFixture(t, os.Stdout, testing.Verbose())
	defer fixture.Finalize()

	test0 := &H{Fixture: fixture}
	test0.RunTestCase__(test0.TestH1, "H 1")

	test1 := &H{Fixture: fixture}
	test1.RunTestCase__(test1.TestH2, "H 2")
}

func (self *H) RunTestCase__(test func(), description string) {
	self.Describe(description)
	test()
}
`,
		description: "One-time fixture setup and teardown",
	},

	/////////////////////////////////////////////////////////////////////////////////////////////

	{
		input: &parse.Fixture{
			StructName: "I",
			Skipped:    true,
			TestCases: []parse.TestCase{
				{Index: 0, Name: "TestI1", StructName: "I"},
				{Index: 1, Name: "TestI2", StructName: "I"},
			},
		},
		expected: `

func TestI(t *testing.T) {
	t.Skip("('I') Skipping test case: 'I 1'")
	t.Skip("('I') Skipping test case: 'I 2'")
}
`,
		description: "Skipping a fixture marks all test cases as skipped",
	},

	/////////////////////////////////////////////////////////////////////////////////////////////

	{
		input: &parse.Fixture{
			StructName: "J",
			TestCases: []parse.TestCase{
				{Index: 0, Name: "TestJ1", StructName: "J", Skipped: true},
				{Index: 1, Name: "TestJ2", StructName: "J"},
			},
		},
		expected: `

func TestJ(t *testing.T) {
	fixture := gunit.NewFixture(t, os.Stdout, testing.Verbose())
	defer fixture.Finalize()

	fixture.Skip("Skipping test case: 'J 1'")

	test1 := &J{Fixture: fixture}
	test1.RunTestCase__(test1.TestJ2, "J 2")
}

func (self *J) RunTestCase__(test func(), description string) {
	self.Describe(description)
	test()
}
`,
		description: "Skipped test case alongside non-skipped test case",
	},

	/////////////////////////////////////////////////////////////////////////////////////////////
}
