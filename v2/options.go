package gunit

type config struct {
	freshFixture    bool
	parallelFixture bool
	parallelTests   bool
}

// Option is a function that modifies a config.
// See Options for provided behaviors.
type Option func(*config)

type singleton struct{}

// Options provides the sole entrypoint
// to the option functions provided by
// this package.
var Options singleton

// FreshFixture signals to Run that the
// new instances of the provided fixture
// are to be instantiated for each and
// every test case. The Setup and Teardown
// methods are also executed on the
// specifically instantiated fixtures.
// NOTE: the SetupSuite and TeardownSuite
// methods are always run on the provided
// fixture instance, regardless of this
// options having been provided.
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
		c.freshFixture = true
		Options.FreshFixture()(c)
	}
}

// UnitTests is a composite option that
// signals to Run that the test suite can
// be treated as a unit-test suite by
// employing parallelism and fresh fixtures
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
// allow reuse of potentially expensive resources.
func (singleton) IntegrationTests() Option {
	return func(c *config) {
		Options.SharedFixture()(c)
	}
}

var defaultOptions = []Option{
	Options.UnitTests(),
}
