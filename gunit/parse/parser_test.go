package parse

import (
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

//////////////////////////////////////////////////////////////////////////////

func TestParseFileWithValidFixturesAndConstructs(t *testing.T) {
	test := &FixtureParsingFixture{t: t, input: comprehensiveTestCode}
	test.ParseFixtures()
	test.AssertFixturesParsedAccuratelyAndCompletely()
}

func TestParseFileWithMalformedSourceCode(t *testing.T) {
	test1 := &FixtureParsingFixture{t: t, input: malformedTestCode}
	test1.ParseFixtures()
	test1.AssertErrorWasReturned()

	test2 := &FixtureParsingFixture{t: t, input: malformedMissingPointerOnEmbeddedStruct}
	test2.ParseFixtures()
	test2.AssertErrorWasReturned()

	test3 := &FixtureParsingFixture{t: t, input: malformedMissingPointerOnReceiver}
	test3.ParseFixtures()
	test3.AssertErrorWasReturned()
}

//////////////////////////////////////////////////////////////////////////////

type FixtureParsingFixture struct {
	t *testing.T

	input      string
	readError  error
	parseError error
	fixtures   []*Fixture
}

func (this *FixtureParsingFixture) ParseFixtures() {
	this.fixtures, this.parseError = Fixtures(this.input)
}

func (this *FixtureParsingFixture) AssertFixturesParsedAccuratelyAndCompletely() {
	this.assertFileWasReadWithoutError()
	this.assertFileWasParsedWithoutError()
	this.assertAllFixturesParsed()
	this.assertParsedFixturesAreCorrect()
}
func (this *FixtureParsingFixture) assertFileWasReadWithoutError() {
	if this.readError != nil {
		this.t.Error("Problem: cound't read the input file:", this.readError)
		this.t.FailNow()
	}
}
func (this *FixtureParsingFixture) assertFileWasParsedWithoutError() {
	if this.parseError != nil {
		this.t.Error("Problem: unexpected parsing error: ", this.parseError)
		this.t.FailNow()
	}
}
func (this *FixtureParsingFixture) assertAllFixturesParsed() {
	if len(this.fixtures) != len(expected) {
		this.t.Logf("Problem: Got back the wrong number of fixtures. Expected: %d Got: %d", len(expected), len(this.fixtures))
		this.t.FailNow()
	}
}
func (this *FixtureParsingFixture) assertParsedFixturesAreCorrect() {
	for x := 0; x < len(expected); x++ {
		key := this.fixtures[x].StructName
		if ok, message := assertions.So(this.fixtures[x], should.Resemble, expected[key]); !ok {
			this.t.Errorf("Comparison failure for record: %d\n%s", x, message)
		}
	}
}

func (this *FixtureParsingFixture) AssertErrorWasReturned() {
	if this.parseError == nil {
		this.t.Error("Expected an error, but got nil instead")
	}
}

//////////////////////////////////////////////////////////////////////////////

var (
	expected = map[string]*Fixture{
		"BowlingGameScoringTests": {
			StructName: "BowlingGameScoringTests",
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
					Index:       3,
					Name:        "LongTestPerfectGame",
					StructName:  "BowlingGameScoringTests",
					Skipped:     false,
					LongRunning: true,
				},
				TestCase{
					Index:       4,
					Name:        "SkipLongTestPerfectGame",
					StructName:  "BowlingGameScoringTests",
					Skipped:     true,
					LongRunning: true,
				},
			},
			TestSetupName:    "SetupTheGame",
			TestTeardownName: "TeardownTheGame",
		},
	}
)

//////////////////////////////////////////////////////////////////////////////
