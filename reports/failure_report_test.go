package reports

import (
	"log"
	"strings"
	"testing"
)

func TestFailureReport(t *testing.T) {
	report := newFailureReport("Failure", readExampleFile)
	report.scanStack(exampleStackFrames)
	actual := report.composeReport()

	actual = strings.TrimSpace(actual)
	if actual != expectedFailureReport {
		t.Errorf("Incorrect failure report.\nGot:\n%s\n\nWant:\n%s", actual, expectedFailureReport)
	}
}

func readExampleFile(path string) string {
	content, found := exampleFiles[path]
	if !found {
		log.Panicln("file not found:", path)
	}
	return content
}

var exampleFiles = map[string]string{
	"/Users/mike/src/github.com/smartystreets/gunit/advanced_examples/bowling_game_test.go":  populateFile("this.AssertEqual(expected, this.game.Score())", 54),
	"/Users/mike/src/github.com/smartystreets/gunit/advanced_examples/bowling_game2_test.go": populateFile("this.assertScore(0)", 24),
}

func populateFile(content string, line int) (final string) {
	return strings.Repeat("\n", line-1) + content + strings.Repeat("\n", 10)
}

var expectedFailureReport = strings.TrimSpace(`
(1): this.AssertEqual(expected, this.game.Score()) // bowling_game_test.go:54
(0): this.assertScore(0) // bowling_game2_test.go:24
Failure`)

var exampleStackFrames = []Frame{
	{Line: 211, Function: "runtime.Callers", File: "/usr/local/go/src/runtime/extern.go"},
	{Line: 11, Function: "github.com/smartystreets/gunit/reports.StackTrace", File: "/Users/mike/src/github.com/smartystreets/gunit/reports/runtime.go"},
	{Line: 93, Function: "github.com/smartystreets/gunit.(*Fixture).fail", File: "/Users/mike/src/github.com/smartystreets/gunit/fixture.go"},
	{Line: 61, Function: "github.com/smartystreets/gunit.(*Fixture).Assert", File: "/Users/mike/src/github.com/smartystreets/gunit/fixture.go"},
	{Line: 66, Function: "github.com/smartystreets/gunit.(*Fixture).AssertEqual", File: "/Users/mike/src/github.com/smartystreets/gunit/fixture.go"},
	{Line: 54, Function: "github.com/smartystreets/gunit/advanced_examples.(*BowlingGameScoringFixture).assertScore", File: "/Users/mike/src/github.com/smartystreets/gunit/advanced_examples/bowling_game_test.go"},
	{Line: 24, Function: "github.com/smartystreets/gunit/advanced_examples.(*BowlingGameScoringFixture).TestAfterAllGutterBallsTheScoreShouldBeZero", File: "/Users/mike/src/github.com/smartystreets/gunit/advanced_examples/bowling_game2_test.go"},
	{Line: 460, Function: "reflect.Value.call", File: "/usr/local/go/src/reflect/value.go"},
	{Line: 321, Function: "reflect.Value.Call", File: "/usr/local/go/src/reflect/value.go"},
	{Line: 86, Function: "github.com/smartystreets/gunit.(*testCase).runTest", File: "/Users/mike/src/github.com/smartystreets/gunit/test_case.go"},
	{Line: 76, Function: "github.com/smartystreets/gunit.(*testCase).runWithSetupAndTeardown", File: "/Users/mike/src/github.com/smartystreets/gunit/test_case.go"},
	{Line: 64, Function: "github.com/smartystreets/gunit.(*testCase).run", File: "/Users/mike/src/github.com/smartystreets/gunit/test_case.go"},
	{Line: 909, Function: "testing.tRunner", File: "/usr/local/go/src/testing/testing.go"},
}
