package gunit_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/smarty/gunit/v2"
)

func TestSuiteWithoutEmbeddedFixture(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Error("Expected panic didn't occur.")
		}
	}()
	gunit.Run(&Suite00{})
}

type Suite00 struct{}

///////////////////////////

func TestSuiteWithSetupsAndTeardowns(t *testing.T) {
	fixture := &Suite01{T: gunit.New(t)}

	gunit.Run(fixture, gunit.Options.IntegrationTests())

	fixture.AssertEqual(fixture.events, []string{
		"SetupSuite",
		"Setup",
		"Test",
		"Teardown",
		"TeardownSuite",
	})
}

type Suite01 struct {
	*gunit.T
	events []string
}

func (this *Suite01) SetupSuite()         { this.record("SetupSuite") }
func (this *Suite01) TeardownSuite()      { this.record("TeardownSuite") }
func (this *Suite01) Setup()              { this.record("Setup") }
func (this *Suite01) Teardown()           { this.record("Teardown") }
func (this *Suite01) Test()               { this.record("Test") }
func (this *Suite01) record(event string) { this.events = append(this.events, event) }

///////////////////////////

func TestFreshFixture(t *testing.T) {
	fixture := &Suite02{T: gunit.New(t)}
	gunit.Run(fixture, gunit.Options.UnitTests())
	fixture.AssertEqual(fixture.counter, 0)
}

type Suite02 struct {
	*gunit.T
	counter int
}

func (this *Suite02) TestSomething() {
	_, _ = this.Write([]byte("*** this should appear in the test log!"))
	this.counter++
}

///////////////////////////

func TestSkip(t *testing.T) {
	fixture := &Suite03{T: gunit.New(t)}
	gunit.Run(fixture)
	fixture.AssertFalse(t.Failed())
}

type Suite03 struct{ *gunit.T }

func (this *Suite03) SkipTestThatFails() {
	this.AssertEqual(1, 2)
}

///////////////////////////

func TestFocus(t *testing.T) {
	fixture := &Suite04{
		T:      gunit.New(t),
		events: make(map[string]struct{}),
	}

	gunit.Run(fixture, gunit.Options.SharedFixture())

	fixture.AssertFalse(t.Failed())
	fixture.AssertEqual(fixture.events, map[string]struct{}{"1": {}})
}

type Suite04 struct {
	*gunit.T
	events map[string]struct{}
}

func (this *Suite04) FocusTest1() {
	this.events["1"] = struct{}{}
}
func (this *Suite04) TestThatFails() {
	this.AssertEqual(1, 2)
}

///////////////////////////

func TestSuiteWithSetupsAndTeardownsSkippedEntirelyIfAllTestsSkipped(t *testing.T) {
	fixture := &Suite05{T: gunit.New(t)}

	gunit.Run(fixture, gunit.Options.SharedFixture())

	fixture.AssertNil(fixture.events)
}

type Suite05 struct {
	*gunit.T
	events []string
}

func (this *Suite05) SetupSuite()         { this.record("SetupSuite") }
func (this *Suite05) TeardownSuite()      { this.record("TeardownSuite") }
func (this *Suite05) Setup()              { this.record("Setup") }
func (this *Suite05) Teardown()           { this.record("Teardown") }
func (this *Suite05) SkipTest()           { this.record("SkipTest") }
func (this *Suite05) record(event string) { this.events = append(this.events, event) }

///////////////////////////

func TestSuiteWithSkippedTests(t *testing.T) {
	fixture := &Suite06{T: gunit.New(t)}

	gunit.Run(fixture, gunit.Options.SharedFixture())

	fixture.AssertEqual(fixture.events, []string{
		"SetupSuite",
		"Setup",
		"Test1",
		"Teardown",
		"TeardownSuite",
	})
}

type Suite06 struct {
	*gunit.T
	events []string
}

func (this *Suite06) SetupSuite()         { this.record("SetupSuite") }
func (this *Suite06) TeardownSuite()      { this.record("TeardownSuite") }
func (this *Suite06) Setup()              { this.record("Setup") }
func (this *Suite06) Teardown()           { this.record("Teardown") }
func (this *Suite06) Test1()              { this.record("Test1") }
func (this *Suite06) SkipTest2()          { this.record("SkipTest2") }
func (this *Suite06) record(event string) { this.events = append(this.events, event) }

///////////////////////////

func shouldEqual(actual any, expected ...any) string {
	if actual == expected[0] {
		return ""
	}
	return fmt.Sprintf("shouldEqual failed: %v vs %v", actual, expected)
}
func TestFixture_PassingAssertions(t *testing.T) {
	fixture := gunit.New(t)
	_, _ = io.WriteString(fixture, "Hello, world!")
	fixture.AssertNil(nil)
	fixture.AssertNotNil(42)
	fixture.AssertTrue(true)
	fixture.AssertFalse(false)
	fixture.AssertEqual(1, 1)
	fixture.AssertNotEqual(1, 2)
	fixture.AssertError(nil, nil)
	fixture.So(1, shouldEqual, 1)
}

type FakeT struct {
	failCount int
	failures  *bytes.Buffer
}

func (this *FakeT) Helper()      {}
func (this *FakeT) Log(a ...any) {}
func (this *FakeT) Error(a ...any) {
	this.failCount++
	_, _ = fmt.Fprintln(this.failures, a...)
}

func TestFixture_FailingAssertions(t *testing.T) {
	fakeT := &FakeT{failures: bytes.NewBuffer(nil)}
	fixture := gunit.New(fakeT)
	fixture.AssertNil(42)
	fixture.AssertNotNil(nil)
	fixture.AssertTrue(false)
	fixture.AssertFalse(true)
	fixture.AssertEqual(1, 2)
	fixture.AssertNotEqual(1, 1)
	fixture.AssertError(errors.New("boink"), errors.New("yoink"))
	fixture.So(1, shouldEqual, 2)
	t.Log(fakeT.failures.String())
	if fakeT.failCount != 8 {
		t.Error("Expected 8 failures, got:", fakeT.failCount)
	}
}
