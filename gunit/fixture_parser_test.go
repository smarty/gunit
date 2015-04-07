package main

import (
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/assertions"
)

func TestParseFixtures(t *testing.T) {
	test := &FixtureParsingFixture{t: t, input: "example_input_test.go"}
	test.ParseFixtures()
	test.AssertFixturesParsedAccuratelyAndCompletely()
}

type FixtureParsingFixture struct {
	t          *testing.T
	input      string
	readError  error
	parseError error
	fixtures   []Fixture
}

func (self *FixtureParsingFixture) ParseFixtures() {
	var source []byte
	source, self.readError = ioutil.ReadFile(self.input)
	self.fixtures, self.parseError = ParseFixtures(string(source))
}

func (self *FixtureParsingFixture) AssertFixturesParsedAccuratelyAndCompletely() {
	if self.readError != nil {
		self.t.Error("Problem: cound't read the input file:", self.readError)
		self.t.FailNow()
	}
	if self.parseError != nil {
		self.t.Error("Problem: unexpected parsing error: ", self.parseError)
		self.t.FailNow()
	}
	if len(self.fixtures) == 0 {
		self.t.Log("Problem: No fixtures to examine.")
		self.t.FailNow()
	}

	if ok, message := So(self.fixtures, ShouldResemble, expected); !ok {
		self.t.Error("\n" + message)
	}
}

var (
	expected = []Fixture{
		{
			StructName: "BowlingGameScoringTests",
		},
	}
)
