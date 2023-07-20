package gunit

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/smarty/gunit/scan"
)

// Run receives an instance of a struct that embeds *Fixture.
// The struct definition may include Setup*, Teardown*, and Test*
// methods which will be run as an xUnit-style test fixture.
func Run(fixture any, t *testing.T, options ...option) {
	t.Helper()
	run(fixture, t, newConfig(options...))
}

// RunSequential (like Run) receives an instance of a struct that embeds *Fixture.
// The fixture is run in much the same way, except that it will not be run in
// parallel with other fixtures in the same package, nor will test cases of the
// corresponding fixture be run in parallel with each other.
//
// # Deprecated
//
// Use Run(fixture, t, Options.AllSequential()) instead.
func RunSequential(fixture any, t *testing.T) {
	t.Helper()

	Run(fixture, t, Options.AllSequential())
}

func run(fixture any, t *testing.T, config configuration) {
	t.Helper()

	ensureEmbeddedFixture(fixture, t)

	_, filename, _, _ := runtime.Caller(2)
	positions := scan.LocateTestCases(filename)

	runner := newFixtureRunner(fixture, t, config, positions)
	runner.ScanFixtureForTestCases()
	runner.RunTestCases()
}

func ensureEmbeddedFixture(fixture any, t TestingT) {
	fixtureType := reflect.TypeOf(fixture)
	embedded, _ := fixtureType.Elem().FieldByName("Fixture")
	if embedded.Type != embeddedGoodExample.Type {
		t.Fatalf("Type (%v) lacks embedded *gunit.Fixture.", fixtureType)
	}
}

var embeddedGoodExample, _ = reflect.TypeOf(new(struct{ *Fixture })).Elem().FieldByName("Fixture")
