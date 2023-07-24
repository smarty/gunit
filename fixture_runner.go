package gunit

import (
	"os"
	"os/signal"
	"reflect"
	"testing"

	"github.com/bugVanisher/gunit/scan"
)

const FixtureParallel = "FixtureParallel"

func newFixtureRunner(
	fixture interface{},
	outerT *testing.T,
	config configuration,
	positions scan.TestCasePositions,
) *fixtureRunner {
	outerT.Parallel()
	return &fixtureRunner{
		config:          config,
		fixtureSetup:    -1,
		fixtureTeardown: -1,
		setup:           -1,
		teardown:        -1,
		outerT:          outerT,
		fixtureType:     reflect.ValueOf(fixture).Type(),
		fixture:         reflect.New(reflect.ValueOf(fixture).Type().Elem()),
		positions:       positions,
	}
}

type fixtureRunner struct {
	outerT      *testing.T
	fixtureType reflect.Type
	fixture     reflect.Value

	fixtureSetup    int
	fixtureTeardown int
	config          configuration
	setup           int
	teardown        int
	focus           []*testCase
	tests           []*testCase
	positions       scan.TestCasePositions
}

func (this *fixtureRunner) ScanFixtureForTestCases() {
	for methodIndex := 0; methodIndex < this.fixtureType.NumMethod(); methodIndex++ {
		methodName := this.fixtureType.Method(methodIndex).Name
		this.scanFixtureMethod(methodIndex, this.newFixtureMethodInfo(methodName))
	}
}

func (this *fixtureRunner) scanFixtureMethod(methodIndex int, method fixtureMethodInfo) {
	switch {
	case method.isFixTureSetup:
		this.fixtureSetup = methodIndex
	case method.isFixTureTeardown:
		this.fixtureTeardown = methodIndex
	case method.isSetup:
		this.setup = methodIndex
	case method.isTeardown:
		this.teardown = methodIndex
	case method.isFocusTest:
		this.focus = append(this.focus, this.buildTestCase(methodIndex, method))
	case method.isTest:
		this.tests = append(this.tests, this.buildTestCase(methodIndex, method))
	}
}

func (this *fixtureRunner) buildTestCase(methodIndex int, method fixtureMethodInfo) *testCase {
	return newTestCase(methodIndex, method, this.config, this.positions)
}

func (this *fixtureRunner) RunTestCases() {
	this.outerT.Helper()

	// Init Fixture for fixtureSetup and fixtureTeardown
	tmpFixture := newFixture(this.outerT, testing.Verbose(), RetrieveTestPackageName())
	this.setInnerFixture(tmpFixture)
	defer this.runFixtureTeardown()
	// Start goroutine to listen for SIGINT signal
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	go func() {
		<-sig
		// Clean up and exit
		this.runFixtureTeardown()
		this.outerT.Error("Interrupted by user")
		os.Exit(1)
	}()
	this.runFixtureSetup()

	if len(this.focus) > 0 {
		this.tests = append(this.focus, skipped(this.tests)...)
	}

	if this.config.SkippedTestCases {
		this.outerT.Skipf("Skip this Fixture:(%v)", this.fixtureType)
	} else {
		if len(this.tests) > 0 {
			this.runTestCases(this.tests)
		} else {
			this.outerT.Skipf("Fixture (%v) has no test cases.", this.fixtureType)
		}
	}
	// fix: replace inner testing
	this.setInnerFixture(tmpFixture)

}

func (this *fixtureRunner) runTestCases(cases []*testCase) {
	this.outerT.Helper()
	pkgName := RetrieveTestPackageName()
	runCases := func(t *testing.T) {
		for _, test := range cases {
			test.Prepare(this.setup, this.teardown, this.fixture, pkgName)
			test.Run(t)
		}
	}
	if this.config.ParallelTestCases() {
		this.outerT.Run(FixtureParallel, func(innerT *testing.T) {
			runCases(innerT)
		})
	} else {
		runCases(this.outerT)
	}
}

func skipped(cases []*testCase) []*testCase {
	for _, test := range cases {
		test.skipped = true
	}
	return cases
}

func (this *fixtureRunner) runFixtureSetup() {
	if this.fixtureSetup >= 0 {
		this.fixture.Method(this.fixtureSetup).Call(nil)
	}
}

func (this *fixtureRunner) runFixtureTeardown() {
	if this.fixtureTeardown >= 0 {
		this.fixture.Method(this.fixtureTeardown).Call(nil)
	}
}

func (this *fixtureRunner) setInnerFixture(innerFixture *Fixture) {
	this.fixture.Elem().FieldByName("Fixture").Set(reflect.ValueOf(innerFixture))
}
