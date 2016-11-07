package gunit

import (
	"fmt"
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
	ensureEmbeddedFixture(fixture)

	return &fixtureRunner{
		setup:       -1,
		teardown:    -1,
		outerT:      t,
		fixtureType: reflect.ValueOf(fixture).Type(),
	}
}

func ensureEmbeddedFixture(fixture interface{}) {
	fixtureType := reflect.TypeOf(fixture)
	_, hasEmbeddedGunitFixture := fixtureType.Elem().FieldByName("Fixture")
	if !hasEmbeddedGunitFixture {
		panic(fmt.Sprintf("Type (%v) lacks embedded *gunit.Fixture.", fixtureType))
	}
}

type fixtureRunner struct {
	outerT      *testing.T
	fixtureType reflect.Type

	setup    int
	teardown int
	tests    []*testCase
}

/**************************************************************************/

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

/**************************************************************************/

func (this *fixtureRunner) RunTestCases() {
	for _, test := range this.tests {
		test.Prepare(this.setup, this.teardown, this.fixtureType)
		test.Run(this.outerT)
	}
}

/**************************************************************************/

type testCase struct {
	methodIndex int
	description string
	skipped     bool
	long        bool

	setup            int
	teardown         int
	innerFixture     *Fixture
	outerFixtureType reflect.Type
	outerFixture     reflect.Value
}

func newTestCase(methodIndex int, method fixtureMethodInfo) *testCase {
	return &testCase{
		methodIndex: methodIndex,
		description: method.name,
		skipped:     method.isSkippedTest,
		long:        method.isLongTest,
	}
}

func (this *testCase) Prepare(setup, teardown int, outerFixtureType reflect.Type) {
	this.setup = setup
	this.teardown = teardown
	this.outerFixtureType = outerFixtureType
}

func (this *testCase) Run(t *testing.T) {
	if this.skipped {
		t.Run(this.description, skip)
	} else if this.long && testing.Short() {
		t.Run(this.description, skipLong)
	} else {
		t.Run(this.description, this.run)
	}
}

func skip(innerT *testing.T)     { innerT.Skip("Skipped test") }
func skipLong(innerT *testing.T) { innerT.Skip("Skipped long-running test") }

func (this *testCase) run(innerT *testing.T) {
	this.initializeFixture(innerT)
	defer this.innerFixture.Finalize()
	this.runWithSetupAndTeardown()
}
func (this *testCase) initializeFixture(innerT *testing.T) {
	this.innerFixture = NewFixture(innerT, testing.Verbose())
	this.outerFixture = reflect.New(this.outerFixtureType.Elem())
	this.outerFixture.Elem().FieldByName("Fixture").Set(reflect.ValueOf(this.innerFixture))
}

func (this *testCase) runWithSetupAndTeardown() {
	this.runSetup()
	defer this.runTeardown()
	this.runTest()
}

func (this *testCase) runSetup() {
	if this.setup < 0 {
		return
	}

	this.outerFixture.Method(this.setup).Call(nil)
}

func (this *testCase) runTest() {
	this.outerFixture.Method(this.methodIndex).Call(nil)
}

func (this *testCase) runTeardown() {
	if this.teardown < 0 {
		return
	}

	this.outerFixture.Method(this.teardown).Call(nil)
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

/**************************************************************************/
