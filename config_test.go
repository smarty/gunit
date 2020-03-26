package gunit

import (
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func TestConfigOptions(t *testing.T) {
	assert := assertions.New(t)

	assert.So(newConfig(), should.Resemble, configuration{})

	assert.So(newConfig(Options.LongRunning()), should.Resemble, configuration{
		LongRunningTestCases: true,
	})

	assert.So(newConfig(Options.SkipAll()), should.Resemble, configuration{
		SkippedTestCases: true,
	})

	assert.So(newConfig(Options.SequentialFixture()), should.Resemble, configuration{
		SequentialFixture: true,
	})

	assert.So(newConfig(Options.SequentialTestCases()), should.Resemble, configuration{
		SequentialTestCases: true,
	})

	assert.So(newConfig(Options.AllSequential()), should.Resemble, configuration{
		SequentialTestCases: true,
		SequentialFixture:   true,
	})

	assert.So(newConfig(Options.AllSequential(), Options.SkipAll(), Options.LongRunning()), should.Resemble, configuration{
		SequentialFixture:    true,
		SequentialTestCases:  true,
		SkippedTestCases:     true,
		LongRunningTestCases: true,
	})
}

func TestConfigMethods(t *testing.T) {
	assert := assertions.New(t)

	parallel := newConfig()
	assert.So(parallel.ParallelFixture(), should.BeTrue)
	assert.So(parallel.ParallelTestCases(), should.BeTrue)

	sequentialFixture := newConfig(Options.SequentialFixture())
	assert.So(sequentialFixture.ParallelFixture(), should.BeFalse)
	assert.So(sequentialFixture.ParallelTestCases(), should.BeTrue)

	sequentialTestCases := newConfig(Options.SequentialTestCases())
	assert.So(sequentialTestCases.ParallelFixture(), should.BeTrue)
	assert.So(sequentialTestCases.ParallelTestCases(), should.BeFalse)

	sequential := newConfig(Options.AllSequential())
	assert.So(sequential.ParallelFixture(), should.BeFalse)
	assert.So(sequential.ParallelTestCases(), should.BeFalse)
}
