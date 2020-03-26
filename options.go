package gunit

var Options options

type options struct{}

// SkipFixture is an option meant to be passed to gunit.Run(...)
// and causes each and every "Test" method in the corresponding
// fixture to be skipped (as if each had been prefixed with
// "Skip"). Even "Test" methods marked with the "Focus" prefix
// will be skipped.
//
// TODO
//
func (options) SkipFixture() option {
	return optionSkipFixture
}

// LongRunningFixture is an option meant to be passed to
// gunit.Run(...) and, in the case that the -short testing
// flag has been passed at the command line, it causes each
// and every "Test" method in the corresponding fixture to
// be skipped (as if each had been prefixed with "Skip").
//
// TODO
//
func (options) LongRunningFixture() option {
	return optionLongRunningFixture
}

// SequentialFixture is an option meant to be passed to
// gunit.Run(...) and signals that the corresponding fixture
// is not to be run in parallel with any tests (by not calling
// t.Parallel() on the provided *testing.T).
func (options) SequentialFixture() option {
	return optionSequentialFixture
}

// SequentialTestCases is an option meant to be passed to
// gunit.Run(...) and prevents gunit from calling t.Parallel()
// on the inner instances of *testing.T passed to the 'subtests'
// corresponding to "Test" methods which are created during
// the natural course of the corresponding invocation of
// gunit.Run(...).
//
// TODO
//
func (options) SequentialTestCases() option {
	return optionSequentialTestCases
}

type option string

const (
	optionSequentialFixture   option = "option:SequentialFixture"
	optionSequentialTestCases option = "option:SequentialTestCases"
	optionSkipFixture         option = "option:SkipFixture"
	optionLongRunningFixture  option = "option:LongRunningFixture"
)

func contains(options []option, value option) bool {
	for _, option := range options {
		if option == value {
			return true
		}
	}
	return false
}
