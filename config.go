package gunit

type configuration struct {
	SequentialFixture    bool
	SequentialTestCases  bool
	SkippedTestCases     bool
	LongRunningTestCases bool
}

func newConfig(options ...option) configuration {
	config := new(configuration)
	Options.composite(options...)(config)
	return *config
}

func (this configuration) ParallelFixture() bool {
	return !this.SequentialFixture
}

func (this configuration) ParallelTestCases() bool {
	return !this.SequentialTestCases
}
