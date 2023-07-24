package gunit

import (
	"testing"

	assertions2 "github.com/bugVanisher/gunit/assertions"
)

func TestConfigOptions(t *testing.T) {
	assert := assertions2.New(t)

	assert.AssertDeepEqual(configuration{}, newConfig())

	assert.AssertDeepEqual(
		configuration{
			LongRunningTestCases: true,
		},
		newConfig(Options.LongRunning()),
	)

	assert.AssertDeepEqual(
		configuration{
			SkippedTestCases: true,
		},
		newConfig(Options.SkipAll()),
	)

	assert.AssertDeepEqual(
		configuration{
			SequentialTestCases: true,
		},
		newConfig(Options.SequentialTestCases()),
	)

	assert.AssertDeepEqual(
		configuration{
			SequentialTestCases: true,
		},
		newConfig(Options.AllSequential()),
	)

	assert.AssertDeepEqual(
		configuration{
			SequentialTestCases:  true,
			SkippedTestCases:     true,
			LongRunningTestCases: true,
		},
		newConfig(Options.AllSequential(), Options.SkipAll(), Options.LongRunning()),
	)
}

func TestConfigMethods(t *testing.T) {
	assert := assertions2.New(t)

	parallel := newConfig()
	assert.AssertTrue(parallel.ParallelTestCases())

	sequentialTestCases := newConfig(Options.SequentialTestCases())
	assert.AssertFalse(sequentialTestCases.ParallelTestCases())

}
