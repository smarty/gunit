package gunit

import (
	"reflect"
	"testing"

	"github.com/smarty/gunit/scan"
)

type testCase struct {
	methodIndex int
	description string
	skipped     bool
	long        bool
	parallel    bool

	setup           int
	teardown        int
	innerFixture    *Fixture
	outerFixture    reflect.Value
	positions       scan.TestCasePositions
	testPackageName string
}

func newTestCase(methodIndex int, method fixtureMethodInfo, config configuration, positions scan.TestCasePositions) *testCase {
	return &testCase{
		parallel:    config.ParallelTestCases(),
		methodIndex: methodIndex,
		description: method.name,
		skipped:     method.isSkippedTest || config.SkippedTestCases,
		long:        method.isLongTest || config.LongRunningTestCases,
		positions:   positions,
	}
}

func (this *testCase) Prepare(setup, teardown int, outerFixture reflect.Value, pkgName string) {
	this.setup = setup
	this.teardown = teardown
	this.outerFixture = outerFixture
	this.testPackageName = pkgName
}

func (this *testCase) Run(t *testing.T) {
	t.Helper()

	if this.skipped {
		t.Run(this.description, this.skip)
	} else if this.long && testing.Short() {
		t.Run(this.description, this.skipLong)
	} else {
		t.Run(this.description, this.run)
	}
}

func (this *testCase) skip(innerT *testing.T) {
	innerT.Skip("\n" + this.positions[innerT.Name()])
}
func (this *testCase) skipLong(innerT *testing.T) {
	innerT.Skipf("Skipped long-running test:\n" + this.positions[innerT.Name()])
}
func (this *testCase) run(innerT *testing.T) {
	innerT.Helper()
	this.initializeFixture(innerT)

	if this.parallel {
		innerT.Parallel()
	}
	defer this.innerFixture.finalize()
	this.runWithSetupAndTeardown()
	if innerT.Failed() {
		innerT.Log("Test definition:\n" + this.positions[innerT.Name()])
	}
}
func (this *testCase) initializeFixture(innerT *testing.T) {
	this.innerFixture = newFixture(innerT, testing.Verbose(), this.testPackageName)
	if this.parallel {
		// new outerFixture in every parallel test
		oldOuterFixture := this.outerFixture
		this.outerFixture = reflect.New(this.outerFixture.Type().Elem())
		this.outerFixture.Elem().Set(reflect.Indirect(oldOuterFixture))
	}
	this.outerFixture.Elem().FieldByName("Fixture").Set(reflect.ValueOf(this.innerFixture))
}

func (this *testCase) runWithSetupAndTeardown() {
	this.runSetup()
	defer this.runTeardown()
	this.runTest()
}

func (this *testCase) runSetup() {
	if this.setup >= 0 {
		this.outerFixture.Method(this.setup).Call(nil)
	}
}

func (this *testCase) runTest() {
	this.outerFixture.Method(this.methodIndex).Call(nil)
}

func (this *testCase) runTeardown() {
	if this.teardown >= 0 {
		this.outerFixture.Method(this.teardown).Call(nil)
	}
}
