package generate

import (
	"strings"
	"testing"

	"github.com/smartystreets/gunit/gunit/parse"
)

func TestGenerateSetupsAndCasesAndTeardowns(t *testing.T) {
	actual := singleLine(TestFile("blah", inputs))
	if actual != expected {
		t.Errorf("\nExpected: [%s]\nActual:   [%s]", expected, actual)
	}
}

func singleLine(value string) string {
	return strings.Replace(value, "\n", "|", -1)
}

//////////////////////////////////////////////////////////////////////////////

var (
	inputs = []*parse.Fixture{
		{
			StructName: "B",
			TestCases:  []parse.TestCase{{Index: 0, Name: "TestB1", StructName: "B"}},
		},
		{
			StructName:    "C",
			TestSetupName: "SetupC_",
			TestCases: []parse.TestCase{
				{Index: 0, Name: "TestC1", StructName: "C"},
				{Index: 1, Name: "TestC2", StructName: "C"},
			},
		},
		{
			StructName:       "D",
			TestTeardownName: "TeardownD_",
			TestCases: []parse.TestCase{
				{Index: 0, Name: "TestD1", StructName: "D"},
				{Index: 1, Name: "TestD2", StructName: "D"},
			},
		},
		{
			StructName:       "E",
			TestSetupName:    "SetupE_",
			TestTeardownName: "TeardownE_",
			TestCases: []parse.TestCase{
				{Index: 0, Name: "TestE1", StructName: "E"},
				{Index: 1, Name: "TestE2", StructName: "E"},
			},
		},
		{
			StructName:       "F",
			FixtureSetupName: "SetupF",
			TestCases: []parse.TestCase{
				{Index: 0, Name: "TestF1", StructName: "F"},
				{Index: 1, Name: "TestF2", StructName: "F"},
			},
		},
		{
			StructName:          "G",
			FixtureTeardownName: "TeardownG",
			TestCases: []parse.TestCase{
				{Index: 0, Name: "TestG1", StructName: "G"},
				{Index: 1, Name: "TestG2", StructName: "G"},
			},
		},
		{
			StructName:          "H",
			FixtureSetupName:    "SetupH",
			FixtureTeardownName: "TeardownH",
			TestCases: []parse.TestCase{
				{Index: 0, Name: "TestH1", StructName: "H"},
				{Index: 1, Name: "TestH2", StructName: "H"},
			},
		},
	}
)

var expected = singleLine(`//////////////////////////////////////////////////////////////////////////////
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
	test0.RunTestCase__(test0.TestB1, "Test b1")
}

func (self *B) RunTestCase__(test func(), description string) {
	self.T.Log(description)
	test()
}

//////////////////////////////////////////////////////////////////////////////

func TestC(t *testing.T) {
	fixture := gunit.NewFixture(t)
	defer fixture.Finalize()

	test0 := &C{Fixture: fixture}
	test0.RunTestCase__(test0.TestC1, "Test c1")

	test1 := &C{Fixture: fixture}
	test1.RunTestCase__(test1.TestC2, "Test c2")
}

func (self *C) RunTestCase__(test func(), description string) {
	self.T.Log(description)
	self.SetupC_()
	test()
}

//////////////////////////////////////////////////////////////////////////////

func TestD(t *testing.T) {
	fixture := gunit.NewFixture(t)
	defer fixture.Finalize()

	test0 := &D{Fixture: fixture}
	test0.RunTestCase__(test0.TestD1, "Test d1")

	test1 := &D{Fixture: fixture}
	test1.RunTestCase__(test1.TestD2, "Test d2")
}

func (self *D) RunTestCase__(test func(), description string) {
	self.T.Log(description)
	defer self.TeardownD_()
	test()
}

//////////////////////////////////////////////////////////////////////////////

func TestE(t *testing.T) {
	fixture := gunit.NewFixture(t)
	defer fixture.Finalize()

	test0 := &E{Fixture: fixture}
	test0.RunTestCase__(test0.TestE1, "Test e1")

	test1 := &E{Fixture: fixture}
	test1.RunTestCase__(test1.TestE2, "Test e2")
}

func (self *E) RunTestCase__(test func(), description string) {
	self.T.Log(description)
	defer self.TeardownE_()
	self.SetupE_()
	test()
}

//////////////////////////////////////////////////////////////////////////////

func TestF(t *testing.T) {
	SetupF()

	fixture := gunit.NewFixture(t)
	defer fixture.Finalize()

	test0 := &F{Fixture: fixture}
	test0.RunTestCase__(test0.TestF1, "Test f1")

	test1 := &F{Fixture: fixture}
	test1.RunTestCase__(test1.TestF2, "Test f2")
}

func (self *F) RunTestCase__(test func(), description string) {
	self.T.Log(description)
	test()
}

//////////////////////////////////////////////////////////////////////////////

func TestG(t *testing.T) {
	defer TeardownG()

	fixture := gunit.NewFixture(t)
	defer fixture.Finalize()

	test0 := &G{Fixture: fixture}
	test0.RunTestCase__(test0.TestG1, "Test g1")

	test1 := &G{Fixture: fixture}
	test1.RunTestCase__(test1.TestG2, "Test g2")
}

func (self *G) RunTestCase__(test func(), description string) {
	self.T.Log(description)
	test()
}

//////////////////////////////////////////////////////////////////////////////

func TestH(t *testing.T) {
	defer TeardownH()
	SetupH()

	fixture := gunit.NewFixture(t)
	defer fixture.Finalize()

	test0 := &H{Fixture: fixture}
	test0.RunTestCase__(test0.TestH1, "Test h1")

	test1 := &H{Fixture: fixture}
	test1.RunTestCase__(test1.TestH2, "Test h2")
}

func (self *H) RunTestCase__(test func(), description string) {
	self.T.Log(description)
	test()
}

//////////////////////////////////////////////////////////////////////////////
`)
