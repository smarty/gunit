package gunit

type option func(*configuration)

var Options singleton

type singleton struct{}

// SkipAll is an option meant to be passed to gunit.Run(...)
// and causes each and every "Test" method in the corresponding
// fixture to be skipped (as if each had been prefixed with
// "Skip"). Even "Test" methods marked with the "Focus" prefix
// will be skipped.
func (singleton) SkipAll() option {
	return func(this *configuration) {
		this.SkippedTestCases = true
	}
}

// LongRunning is an option meant to be passed to
// gunit.Run(...) and, in the case that the -short testing
// flag has been passed at the command line, it causes each
// and every "Test" method in the corresponding fixture to
// be skipped (as if each had been prefixed with "Skip").
func (singleton) LongRunning() option {
	return func(this *configuration) {
		this.LongRunningTestCases = true
	}
}

// SequentialFixture is an option meant to be passed to
// gunit.Run(...) and signals that the corresponding fixture
// is not to be run in parallel with any tests (by not calling
// t.Parallel() on the provided *testing.T).
func (singleton) SequentialFixture() option {
	return func(this *configuration) {
		this.SequentialFixture = true
	}
}

// SequentialTestCases is an option meant to be passed to
// gunit.Run(...) and prevents gunit from calling t.Parallel()
// on the inner instances of *testing.T passed to the 'subtests'
// corresponding to "Test" methods which are created during
// the natural course of the corresponding invocation of
// gunit.Run(...).
func (singleton) SequentialTestCases() option {
	return func(this *configuration) {
		this.SequentialTestCases = true
	}
}

// AllSequential() has the combined effect of passing the
// following options to gunit.Run(...):
// 1. SequentialFixture
// 2. SequentialTestCases
func (singleton) AllSequential() option {
	return Options.composite(
		Options.SequentialFixture(),
		Options.SequentialTestCases(),
	)
}

// composite allows graceful chaining of options.
func (singleton) composite(options ...option) option {
	return func(this *configuration) {
		for _, option := range options {
			option(this)
		}
	}
}
