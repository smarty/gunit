package gunit

import "testing"

type config struct {
	skipAllTests    bool
	skipReason      string
	freshFixture    bool
	parallelFixture bool
	parallelTests   bool
	short           func() bool
}

// Option is a function that modifies a config.
// See Options for provided behaviors.
type Option func(*config)

type singleton struct{}

// Options provides the sole entrypoint
// to the option functions provided by
// this package.
var Options singleton

// SkipAll causes each and every "Test" method in the corresponding
// fixture to be skipped (as if each had been prefixed with
// "Skip"). Even "Test" methods marked with the "Focus" prefix
// will be skipped.
func (singleton) SkipAll() Option {
	return func(this *config) {
		this.skipAllTests = true
	}
}

// LongRunning signals to Run that, when `go test` was invoked
// with the -short flag, each and every "Test" method in the
// corresponding fixture is to be skipped (as if each had been
// prefixed with "Skip").
func (singleton) LongRunning() Option {
	return func(this *config) {
		if this.short() {
			Options.SkipAll()(this)
			this.skipReason = "Skipping long-running fixture (-short)"
		}
	}
}

// FreshFixture signals to Run that the
// new instances of the provided fixture
// are to be instantiated for each and
// every test case. The Setup and Teardown
// methods are also executed on the
// specifically instantiated fixtures.
// NOTE: the SetupSuite and TeardownSuite
// methods are always run on the provided
// fixture instance, regardless of this
// option having been provided.
func (singleton) FreshFixture() Option {
	return func(c *config) {
		c.freshFixture = true
	}
}

// SharedFixture signals to Run that the
// provided fixture instance is to be used
// to run all test methods. This mode is
// not compatible with ParallelFixture or
// ParallelTests and disables them.
func (singleton) SharedFixture() Option {
	return func(c *config) {
		c.freshFixture = false
		c.parallelTests = false
		c.parallelFixture = false
	}
}

// ParallelFixture signals to Run that the
// provided fixture instance can be executed
// in parallel with other go test functions.
// This option assumes that `go test` was
// invoked with the -parallel flag.
func (singleton) ParallelFixture() Option {
	return func(c *config) {
		c.parallelFixture = true
	}
}

// ParallelTests signals to Run that the
// test methods on the provided fixture
// instance can be executed in parallel
// with each other. This option assumes
// that `go test` was invoked with the
// -parallel flag.
func (singleton) ParallelTests() Option {
	return func(c *config) {
		c.parallelTests = true
		Options.FreshFixture()(c)
	}
}

// UnitTests signals to Run that the test
// suite can be treated as a unit-test suite
// by employing parallelism and fresh fixtures
// to maximize the chances of exposing
// unwanted coupling between tests.
func (singleton) UnitTests() Option {
	return func(c *config) {
		Options.ParallelTests()(c)
		Options.ParallelFixture()(c)
	}
}

// IntegrationTests is a composite option that
// signals to Run that the test suite should be
// treated as an integration test suite, avoiding
// parallelism and utilizing shared fixtures to
// allow reuse of potentially expensive resources,
// and skipping the entire fixture when `go test`
// is invoked with the -short flag.
func (singleton) IntegrationTests() Option {
	return func(c *config) {
		Options.SharedFixture()(c)
		Options.LongRunning()(c)
	}
}

// shortFunc supplies a func of the same signature
// as testing.Short, which is the default.
func (singleton) shortFunc(short func() bool) Option {
	return func(this *config) {
		this.short = short
	}
}

// skipReason supplies text for the value passed to t.Skip
// in the event that no test methods are able to be executed.
func (singleton) skipReason(reason string) Option {
	return func(c *config) {
		c.skipReason = reason
	}
}

var defaultOptions = []Option{
	Options.UnitTests(),
	Options.shortFunc(testing.Short),
	Options.skipReason(defaultSkipReason),
}

const defaultSkipReason = "NOT IMPLEMENTED (no test cases defined, or they are all marked as skipped)"
