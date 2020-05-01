package gunit

import (
	"testing"

	assertions2 "github.com/smartystreets/gunit/assertions"
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
			SequentialFixture: true,
		},
		newConfig(Options.SequentialFixture()),
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
			SequentialFixture:   true,
		},
		newConfig(Options.AllSequential()),
	)

	assert.AssertDeepEqual(
		configuration{
			SequentialFixture:    true,
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
	assert.AssertTrue(parallel.ParallelFixture())
	assert.AssertTrue(parallel.ParallelTestCases())

	sequentialFixture := newConfig(Options.SequentialFixture())
	assert.AssertFalse(sequentialFixture.ParallelFixture())
	assert.AssertTrue(sequentialFixture.ParallelTestCases())

	sequentialTestCases := newConfig(Options.SequentialTestCases())
	assert.AssertTrue(sequentialTestCases.ParallelFixture())
	assert.AssertFalse(sequentialTestCases.ParallelTestCases())

	sequential := newConfig(Options.AllSequential())
	assert.AssertFalse(sequential.ParallelFixture())
	assert.AssertFalse(sequential.ParallelTestCases())
}
