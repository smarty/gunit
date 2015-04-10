package parse

import (
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/assertions"
)

//////////////////////////////////////////////////////////////////////////////

func TestParseFileWithValidFixturesAndConstructs(t *testing.T) {
	test := &FixtureParsingFixture{t: t, input: "example_input_test.go.txt"}
	test.ParseFixtures()
	test.AssertFixturesParsedAccuratelyAndCompletely()
}

func TestParseFileWithMalformedSourceCode(t *testing.T) {
	test := &FixtureParsingFixture{t: t, input: "example_input_failure_test.go.txt"}
	test.ParseFixtures()
	test.AssertErrorWasReturned()
}

//////////////////////////////////////////////////////////////////////////////

type FixtureParsingFixture struct {
	t *testing.T

	input      string
	readError  error
	parseError error
	fixtures   []*Fixture
}

func (self *FixtureParsingFixture) ParseFixtures() {
	var source []byte
	source, self.readError = ioutil.ReadFile(self.input)
	self.fixtures, self.parseError = ParseFixtures(string(source))
}

func (self *FixtureParsingFixture) AssertFixturesParsedAccuratelyAndCompletely() {
	self.assertFileWasReadWithoutError()
	self.assertFileWasParsedWithoutError()
	self.assertAllFixturesParsed()
	self.assertParsedFixturesAreCorrect()
}
func (self *FixtureParsingFixture) assertFileWasReadWithoutError() {
	if self.readError != nil {
		self.t.Error("Problem: cound't read the input file:", self.readError)
		self.t.FailNow()
	}
}
func (self *FixtureParsingFixture) assertFileWasParsedWithoutError() {
	if self.parseError != nil {
		self.t.Error("Problem: unexpected parsing error: ", self.parseError)
		self.t.FailNow()
	}
}
func (self *FixtureParsingFixture) assertAllFixturesParsed() {
	if len(self.fixtures) != len(expected) {
		self.t.Logf("Problem: Got back the wrong number of fixtures. Expected: %d Got: %d", len(expected), len(self.fixtures))
		self.t.FailNow()
	}
}
func (self *FixtureParsingFixture) assertParsedFixturesAreCorrect() {
	for x := 0; x < len(expected); x++ {
		key := self.fixtures[x].StructName
		if ok, message := So(self.fixtures[x], ShouldResemble, expected[key]); !ok {
			self.t.Errorf("Comparison failure for record: %d\n%s", x, message)
		}
	}
}

func (self *FixtureParsingFixture) AssertErrorWasReturned() {
	if self.parseError == nil {
		self.t.Error("Expected an error, but got nil instead")
	}
}

//////////////////////////////////////////////////////////////////////////////

var (
	expected = map[string]*Fixture{
		"BowlingGameScoringTests": {
			StructName:          "BowlingGameScoringTests",
			FixtureSetupName:    "SetupBowlingGameScoringTests",
			FixtureTeardownName: "TeardownBowlingGameScoringTests",
			HasFocusedTestCases: true,
			TestCases: []TestCase{
				TestCase{
					Index:      0,
					Name:       "TestAfterAllGutterBallsTheScoreShouldBeZero",
					StructName: "BowlingGameScoringTests",
				},
				TestCase{
					Index:      1,
					Name:       "TestAfterAllOnesTheScoreShouldBeTwenty",
					StructName: "BowlingGameScoringTests",
				},
				TestCase{
					Index:      2,
					Name:       "SkipTestASpareDeservesABonus",
					StructName: "BowlingGameScoringTests",
					Skipped:    true,
				},
				TestCase{
					Index:      3,
					Name:       "FocusTestAStrikeDeservesABigBonus",
					StructName: "BowlingGameScoringTests",
					Focused:    true,
				},
			},
			TestSetupName:    "SetupTheGame",
			TestTeardownName: "TeardownTheGame",
		},
		"SkipFixture": {
			StructName: "SkipFixture",
			Skipped:    true,
		},
		"FocusFixture": {
			StructName: "FocusFixture",
			Focused:    true,
		},
	}
)

//////////////////////////////////////////////////////////////////////////////
