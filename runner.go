package gunit

import (
	"reflect"
	"strings"
	"testing"
)

func Run(fixture interface{}, t *testing.T) {
	runner := newFixtureRunner(fixture, t)
	runner.ScanFixtureForTestCases()
	runner.RunTestCases()
}

func newFixtureRunner(fixture interface{}, t *testing.T) *fixtureRunner {
	return &fixtureRunner{
		setup:       -1,
		teardown:    -1,
		outerT:      t,
		fixtureType: reflect.ValueOf(fixture).Type(),
	}
}

type fixtureRunner struct {
	outerT        *testing.T
	fixtureType   reflect.Type

	setup         int
	teardown      int
	tests         []testCaseInfo

	activeTest    testCaseInfo
	activeFixture reflect.Value
}

func (this *fixtureRunner) ScanFixtureForTestCases() {
	for methodIndex := 0; methodIndex < this.fixtureType.NumMethod(); methodIndex++ {
		methodName := this.fixtureType.Method(methodIndex).Name
		this.scanFixtureMethod(methodIndex, this.newFixtureMethodInfo(methodName))
	}
}

func (this *fixtureRunner) scanFixtureMethod(methodIndex int, method fixtureMethodInfo) {
	switch {
	case method.isSetup:
		this.setup = methodIndex
	case method.isTeardown:
		this.teardown = methodIndex
	case method.isTest:
		this.tests = append(this.tests, newTestCase(methodIndex, method))
	}
}

func (this *fixtureRunner) RunTestCases() {
	for _, test := range this.tests {
		this.activeTest = test
		this.runTestCase()
	}
}
func (this *fixtureRunner) runTestCase() {
	if this.activeTest.skipped {
		this.outerT.Run(this.activeTest.description, this.skip)
	} else if this.activeTest.long && testing.Short() {
		this.outerT.Run(this.activeTest.description, this.skipLong)
	} else {
		this.outerT.Run(this.activeTest.description, this.run)
	}
}
func (this *fixtureRunner) skip(t *testing.T)     { t.Skip("Skipped test") }
func (this *fixtureRunner) skipLong(t *testing.T) { t.Skip("Skipped long-running test") }
func (this *fixtureRunner) run(t *testing.T) {
	inner := this.initializeFixture(t)
	defer inner.Finalize()
	this.runTestCaseWithSetupAndTeardown()
}
func (this *fixtureRunner) initializeFixture(t *testing.T) *Fixture {
	inner := NewFixture(t, testing.Verbose())
	this.activeFixture = reflect.New(this.fixtureType.Elem())
	this.activeFixture.Elem().FieldByName("Fixture").Set(reflect.ValueOf(inner))
	return inner
}
func (this *fixtureRunner) runTestCaseWithSetupAndTeardown() {
	this.runSetup()
	this.runTest()
	this.runTeardown()
}
func (this *fixtureRunner) runSetup() {
	if this.setup < 0 {
		return
	}
	this.activeFixture.Method(this.setup).Call(nil)
}
func (this *fixtureRunner) runTest() {
	this.activeFixture.Method(this.activeTest.methodIndex).Call(nil)
}
func (this *fixtureRunner) runTeardown() {
	if this.teardown < 0 {
		return
	}
	this.activeFixture.Method(this.teardown).Call(nil)
}

/**************************************************************************/

type testCaseInfo struct {
	methodIndex int
	description string
	skipped     bool
	long        bool
}

func newTestCase(methodIndex int, method fixtureMethodInfo) testCaseInfo {
	return testCaseInfo{
		methodIndex: methodIndex,
		description: method.name,
		skipped:     method.isSkippedTest,
		long:        method.isLongTest,
	}
}

/**************************************************************************/

type fixtureMethodInfo struct {
	name          string
	isSetup       bool
	isTeardown    bool
	isTest        bool
	isLongTest    bool
	isSkippedTest bool
}

func (this *fixtureRunner) newFixtureMethodInfo(name string) fixtureMethodInfo {
	isTest := strings.HasPrefix(name, "Test")
	isLongTest := strings.HasPrefix(name, "LongTest")
	isSkippedTest := strings.HasPrefix(name, "SkipTest")
	isSkippedLongTest := strings.HasPrefix(name, "SkipLongTest")

	return fixtureMethodInfo{
		name:          name,
		isSetup:       strings.HasPrefix(name, "Setup"),
		isTeardown:    strings.HasPrefix(name, "Teardown"),
		isLongTest:    isLongTest,
		isSkippedTest: isSkippedTest || isSkippedLongTest,
		isTest:        isTest || isLongTest || isSkippedTest || isSkippedLongTest,
	}
}
