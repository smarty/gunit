package gunit

import (
	"testing"

	"github.com/smarty/gunit/v2/assert"
	"github.com/smarty/gunit/v2/assert/should"
)

func TestLongRunningFixtureSkippedInShortMode(t *testing.T) {
	fixture := &LongRunningSuite{}
	t.Run("run", func(t *testing.T) {
		Run(fixture, t, Options.SharedFixture(), stubShortMode(true), Options.LongRunning())
	})
	assert.So(t, fixture.events, should.BeEmpty)
	assert.So(t, t.Failed(), should.BeFalse)
}

func TestLongRunningFixtureRunsWhenNotShortMode(t *testing.T) {
	fixture := &LongRunningSuite{}
	t.Run("run", func(t *testing.T) {
		Run(fixture, t, Options.SharedFixture(), stubShortMode(false), Options.LongRunning())
	})
	assert.So(t, fixture.events, should.Equal, []string{
		"SetupSuite",
		"Setup",
		"Test1",
		"Teardown",
		"TeardownSuite",
	})
	assert.So(t, t.Failed(), should.BeFalse)
}

func TestIntegrationTestsFixtureSkippedInShortMode(t *testing.T) {
	fixture := &LongRunningSuite{}
	t.Run("run", func(t *testing.T) {
		Run(fixture, t, stubShortMode(true), Options.IntegrationTests())
	})
	assert.So(t, fixture.events, should.BeEmpty)
	assert.So(t, t.Failed(), should.BeFalse)
}

func TestLongRunningShortModeSkipReason(t *testing.T) {
	config := &config{}
	stubShortMode(true)(config)
	Options.LongRunning()(config)
	assert.So(t, config.skipAllTests, should.BeTrue)
	assert.So(t, config.skipReason, should.Equal, "Skipping long-running fixture (-short)")
}

func TestDefaultSkipReason(t *testing.T) {
	config := &config{}
	stubShortMode(true)(config)
	Options.skipReason(defaultSkipReason)(config)
	assert.So(t, config.skipAllTests, should.BeFalse)
	assert.So(t, config.skipReason, should.Equal, "NOT IMPLEMENTED (no test cases defined, or they are all marked as skipped)")
}

func stubShortMode(short bool) Option {
	return Options.shortFunc(func() bool { return short })
}

type LongRunningSuite struct {
	*Fixture
	events []string
}

func (this *LongRunningSuite) SetupSuite()         { this.record("SetupSuite") }
func (this *LongRunningSuite) TeardownSuite()      { this.record("TeardownSuite") }
func (this *LongRunningSuite) Setup()              { this.record("Setup") }
func (this *LongRunningSuite) Teardown()           { this.record("Teardown") }
func (this *LongRunningSuite) Test1()              { this.record("Test1") }
func (this *LongRunningSuite) record(event string) { this.events = append(this.events, event) }
