package gunit_test

import (
	"testing"

	"github.com/smarty/gunit/v2"
	"github.com/smarty/gunit/v2/should"
)

func TestSuiteWithoutEmbeddedFixture(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Expected panic didn't occur.")
		}
	}()
	gunit.Run(&Suite00{}, t)
}

type Suite00 struct{}

func (this *Suite00) Test() {}

///////////////////////////

func TestSuiteWithSetupsAndTeardowns(t *testing.T) {
	fixture := &Suite01{}

	gunit.Run(fixture, t, gunit.Options.IntegrationTests())

	fixture.So(fixture.events, should.Equal, []string{
		"SetupSuite",
		"Setup",
		"Test",
		"Teardown",
		"TeardownSuite",
	})
}

type Suite01 struct {
	*gunit.Fixture
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
	fixture := &Suite02{}
	gunit.Run(fixture, t, gunit.Options.UnitTests())
	fixture.So(fixture.counter, should.Equal, 0)
}

type Suite02 struct {
	*gunit.Fixture
	counter int
}

func (this *Suite02) TestSomething() {
	_, _ = this.Write([]byte("*** this should appear in the test log!"))
	this.counter++
}

///////////////////////////

func TestSkip(t *testing.T) {
	fixture := &Suite03{}
	gunit.Run(fixture, t)
	fixture.So(t.Failed(), should.Equal, false)
}

type Suite03 struct{ *gunit.Fixture }

func (this *Suite03) SkipTestThatFails() {
	this.So(1, should.Equal, 2)
}

///////////////////////////

func TestFocus(t *testing.T) {
	fixture := &Suite04{events: make(map[string]struct{})}

	gunit.Run(fixture, t, gunit.Options.SharedFixture())

	fixture.So(t.Failed(), should.Equal, false)
	fixture.So(fixture.events, should.Equal, map[string]struct{}{"1": {}})
}

type Suite04 struct {
	*gunit.Fixture
	events map[string]struct{}
}

func (this *Suite04) FocusTest1() {
	this.events["1"] = struct{}{}
}
func (this *Suite04) TestThatFails() {
	this.So(1, should.Equal, 2)
}

///////////////////////////

func TestSuiteWithSetupsAndTeardownsSkippedEntirelyIfAllTestsSkipped(t *testing.T) {
	fixture := &Suite05{}

	gunit.Run(fixture, t, gunit.Options.SharedFixture())

	fixture.So(fixture.events, should.Equal, nil)
}

type Suite05 struct {
	*gunit.Fixture
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
	fixture := &Suite06{}

	gunit.Run(fixture, t, gunit.Options.SharedFixture())

	fixture.So(fixture.events, should.Equal, []string{
		"SetupSuite",
		"Setup",
		"Test1",
		"Teardown",
		"TeardownSuite",
	})
}

type Suite06 struct {
	*gunit.Fixture
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

func TestSuiteThatPanics(t *testing.T) {
	t.Skip("Unskip to see this test fail with a nicely filtered and formatted stack trace.")

	fixture := &Suite08{}

	gunit.Run(fixture, t, gunit.Options.SharedFixture())
}

type Suite08 struct {
	*gunit.Fixture
}

func (this *Suite08) Test1() {
	panic("boink")
}
