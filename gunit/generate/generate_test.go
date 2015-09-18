package generate

import (
	"strings"
	"testing"

	"github.com/smartystreets/gunit/gunit/parse"
)

func TestGenerateTestFileFails(t *testing.T) {
	fixtures := []*parse.Fixture{{StructName: "Not a valid struct name (spaces and parens!)"}}
	file, err := TestFile("blah", fixtures, "", nil)
	if err == nil {
		t.Error("Expected a generate error, got nil instead.")
	}
	if len(file) > 0 {
		t.Error("Expected no generated content, got:", string(file))
	}
}

func TestGenerateWithoutPackageNameFails(t *testing.T) {
	fixtures := []*parse.Fixture{{StructName: "A"}}
	file, err := TestFile("", fixtures, "", nil)
	if err == nil {
		t.Error("Expected a generate error, got nil instead.")
	}
	if len(file) > 0 {
		t.Error("Expected no generated content, got:", file)
	}
}

func TestGenerateValidTestFile(t *testing.T) {
	fixtures := []*parse.Fixture{{StructName: "A"}}
	file, err := TestFile("blah", fixtures, "42", map[string]string{"hello": "world"})
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
	"testing"

	"github.com/smartystreets/gunit"
)

///////////////////////////////////////////////////////////////////////////////

func Test_A(t *testing.T) {
	t.Skip("Fixture 'A' has no test cases.")
}

///////////////////////////////////////////////////////////////////////////////

var __code__ = map[string]string{"hello": "world"}

func init() {
	gunit.Validate("42")
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////// Generated Code //
///////////////////////////////////////////////////////////////////////////////
`

//////////////////////////////////////////////////////////////////////////////

func TestGenerateTestCases(t *testing.T) {
	for _, test := range testFunction_TestCases {
		if test.SKIP {
			t.Log("Skipping:", test.description)
			continue
		}
		function, err := TestCases(test.input)
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
			t.Errorf("FAILED: '%s'\nExpected:\n%s\n\nActual:\n%s", test.description, expected, actual)
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

func Test_A(t *testing.T) {
	t.Skip("Fixture 'A' has no test cases.")
}

`,
		description: "single fixture with no test cases",
	},

	/////////////////////////////////////////////////////////////////////////////////////////////

	{
		input: &parse.Fixture{
			Filename:   "file_test.go",
			StructName: "B",
			TestCases:  []parse.TestCase{{Index: 0, Name: "TestB1", StructName: "B"}},
		},
		expected: `

func Test_B__b_1(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose(), __code__["file_test.go"])
	defer fixture.Finalize()
	test := &B{Fixture: fixture}
	test.TestB1()
}

`,
		description: "A fixture with a single test case, no setups, no teardowns, no skips",
	},

	/////////////////////////////////////////////////////////////////////////////////////////////

	{
		input: &parse.Fixture{
			Filename:      "file_test.go",
			StructName:    "C",
			TestSetupName: "SetupC_",
			TestCases: []parse.TestCase{
				{Index: 0, Name: "TestC1", StructName: "C"},
				{Index: 1, Name: "TestC2", StructName: "C"},
			},
		},
		expected: `

func Test_C__c_1(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose(), __code__["file_test.go"])
	defer fixture.Finalize()
	test := &C{Fixture: fixture}
	test.SetupC_()
	test.TestC1()
}

func Test_C__c_2(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose(), __code__["file_test.go"])
	defer fixture.Finalize()
	test := &C{Fixture: fixture}
	test.SetupC_()
	test.TestC2()
}

`,
		description: "A fixture with two test cases and a setup",
	},
	/////////////////////////////////////////////////////////////////////////////////////////////

	{
		input: &parse.Fixture{
			Filename:         "file_test.go",
			StructName:       "D",
			TestTeardownName: "TeardownD_",
			TestCases: []parse.TestCase{
				{Index: 0, Name: "TestD1", StructName: "D"},
				{Index: 1, Name: "TestD2", StructName: "D"},
			},
		},
		expected: `

func Test_D__d_1(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose(), __code__["file_test.go"])
	defer fixture.Finalize()
	test := &D{Fixture: fixture}
	defer test.TeardownD_()
	test.TestD1()
}

func Test_D__d_2(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose(), __code__["file_test.go"])
	defer fixture.Finalize()
	test := &D{Fixture: fixture}
	defer test.TeardownD_()
	test.TestD2()
}

`,
		description: "A fixture with a two test cases and a teardown",
	},

	/////////////////////////////////////////////////////////////////////////////////////////////

	{
		input: &parse.Fixture{
			Filename:         "file_test.go",
			StructName:       "E",
			TestSetupName:    "SetupE_",
			TestTeardownName: "TeardownE_",
			TestCases: []parse.TestCase{
				{Index: 0, Name: "TestE1", StructName: "E"},
				{Index: 1, Name: "TestE2", StructName: "E"},
			},
		},
		expected: `

func Test_E__e_1(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose(), __code__["file_test.go"])
	defer fixture.Finalize()
	test := &E{Fixture: fixture}
	defer test.TeardownE_()
	test.SetupE_()
	test.TestE1()
}

func Test_E__e_2(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose(), __code__["file_test.go"])
	defer fixture.Finalize()
	test := &E{Fixture: fixture}
	defer test.TeardownE_()
	test.SetupE_()
	test.TestE2()
}

`,
		description: "A fixture with two test cases, a setup and a teardown",
	},

	/////////////////////////////////////////////////////////////////////////////////////////////

	{
		input: &parse.Fixture{
			Filename:   "file_test.go",
			StructName: "I",
			TestCases: []parse.TestCase{
				{Index: 0, Name: "TestI1", StructName: "I", Skipped: true},
				{Index: 1, Name: "TestI2", StructName: "I", Skipped: true},
			},
		},
		expected: `
func Test_I__i_1(t *testing.T) {
	t.Skip("Skipping test case: 'TestI1'")

	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose(), __code__["file_test.go"])
	defer fixture.Finalize()
	test := &I{Fixture: fixture}
	test.TestI1()
}

func Test_I__i_2(t *testing.T) {
	t.Skip("Skipping test case: 'TestI2'")

	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose(), __code__["file_test.go"])
	defer fixture.Finalize()
	test := &I{Fixture: fixture}
	test.TestI2()
}
`,
		description: "Skipping a fixture marks all test cases as skipped",
	},

	/////////////////////////////////////////////////////////////////////////////////////////////

	{
		input: &parse.Fixture{
			Filename:   "file_test.go",
			StructName: "J",
			TestCases: []parse.TestCase{
				{Index: 0, Name: "TestJ1", StructName: "J", Skipped: true},
				{Index: 1, Name: "TestJ2", StructName: "J"},
			},
		},
		expected: `

func Test_J__j_1(t *testing.T) {
	t.Skip("Skipping test case: 'TestJ1'")

	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose(), __code__["file_test.go"])
	defer fixture.Finalize()
	test := &J{Fixture: fixture}
	test.TestJ1()
}

func Test_J__j_2(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose(), __code__["file_test.go"])
	defer fixture.Finalize()
	test := &J{Fixture: fixture}
	test.TestJ2()
}

`,
		description: "Skipped test case alongside non-skipped test case",
	},

	/////////////////////////////////////////////////////////////////////////////////////////////

	{
		input: &parse.Fixture{
			Filename:   "file_test.go",
			StructName: "K",
			TestCases: []parse.TestCase{
				{Index: 0, Name: "TestK1", StructName: "K", LongRunning: true, Skipped: false},
				{Index: 1, Name: "TestK2", StructName: "K", LongRunning: true, Skipped: true},
			},
		},
		expected: `

func Test_K__k_1(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long-running test case.")
	}

	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose(), __code__["file_test.go"])
	defer fixture.Finalize()
	test := &K{Fixture: fixture}
	test.TestK1()
}

func Test_K__k_2(t *testing.T) {
	t.Skip("Skipping test case: 'TestK2'")

	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose(), __code__["file_test.go"])
	defer fixture.Finalize()
	test := &K{Fixture: fixture}
	test.TestK2()
}
`,
		description: "Skipped long-running test case alongside non-skipped long-running test case",
	},

	/////////////////////////////////////////////////////////////////////////////////////////////
}
