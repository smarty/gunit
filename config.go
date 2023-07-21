package gunit

type configuration struct {
	SequentialTestCases  bool
	SkippedTestCases     bool
	LongRunningTestCases bool
}

func newConfig(options ...option) configuration {
	config := new(configuration)
	Options.composite(options...)(config)
	return *config
}

func (this configuration) ParallelTestCases() bool {
	return !this.SequentialTestCases
}
