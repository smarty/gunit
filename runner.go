package gunit

import (
	"reflect"
	"strings"
	"testing"
)

func Run(fixture interface{}, t *testing.T) {
	runner := new(fixtureRunner)
	runner.outerT = t
	runner.outerFixture = fixture
	runner.fixtureValue = reflect.ValueOf(runner.outerFixture)
	runner.fixtureType = runner.fixtureValue.Type()

	runner.ScanFixture()
	runner.RunTestCases()
}

type fixtureRunner struct {
	outerFixture interface{}
	outerT       *testing.T
	fixtureValue reflect.Value
	fixtureType  reflect.Type

	setupYes    bool
	setup       int
	teardownYes bool
	teardown    int
	tests       []TestCaseInfo
}

func (this *fixtureRunner) ScanFixture() {
	for m := 0; m < this.fixtureType.NumMethod(); m++ {
		methodName := this.fixtureType.Method(m).Name
		if strings.HasPrefix(methodName, "Setup") {
			this.setup = m
			this.setupYes = true
		} else if strings.HasPrefix(methodName, "Teardown") {
			this.teardown = m
			this.teardownYes = true
		} else if strings.HasPrefix(methodName, "Test") {
			this.tests = append(this.tests, TestCaseInfo{
				methodIndex: m,
			})
		} else if strings.HasPrefix(methodName, "SkipTest") || strings.HasPrefix(methodName, "SkipLongTest") {
			this.tests = append(this.tests, TestCaseInfo{
				methodIndex: m,
				skipped:     true,
				long:        strings.HasPrefix(methodName, "SkipLongTest"),
			})
		} else if strings.HasPrefix(methodName, "LongTest") {
			this.tests = append(this.tests, TestCaseInfo{
				methodIndex: m,
				long:        true,
			})
		}
	}
}

func (this *fixtureRunner) description(i int) string {
	return this.fixtureType.Method(i).Name
}
func (this *fixtureRunner) RunTestCases() {
	for _, test := range this.tests {
		if test.skipped {
			this.outerT.Run(this.description(test.methodIndex), func(t *testing.T) { t.Skip() })
		}

		this.outerT.Run(this.description(test.methodIndex), func(t *testing.T) {
			fixture := reflect.New(this.fixtureType.Elem())
			inner := NewFixture(t, testing.Verbose())
			defer inner.Finalize()

			fixture.Elem().FieldByName("Fixture").Set(reflect.ValueOf(inner))

			if this.setupYes {
				fixture.Method(this.setup).Call(nil)
			}

			fixture.Method(test.methodIndex).Call(nil)

			if this.teardownYes {
				fixture.Method(this.teardown).Call(nil)
			}
		})
	}
}

type TestCaseInfo struct {
	methodIndex int
	skipped     bool
	long        bool
}
