// These tests demonstrate the very ugliness this project aims to eliminate.
// Lots of duplicate setup logic, etc...
package gunit

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/smartystreets/assertions/should"
)

func TestFinalizeAfterNoActions(t *testing.T) {
	defer reset()
	log, fake, fixture := setup(false)

	fixture.Finalize()

	if fake.failed {
		t.Error("Fake should not have been marked as failed.")
	}
	if fake.skipped {
		t.Error("Fake should not have been marked as skipped.")
	}
	if log.Len() > 0 {
		t.Errorf("Output was not blank: '%s'", log.String())
	}
}

func TestFinalizeAfterPass_NotVerbose(t *testing.T) {
	defer reset()
	log, _, fixture := setup(false)

	fixture.Describe("Hello")
	fixture.Finalize()

	if output := strings.TrimSpace(log.String()); output != "" {
		t.Errorf("Unexpected output: '%s'", log.String())
	}
}

func TestFinalizeAfterPass_Verbose(t *testing.T) {
	defer reset()
	log, _, fixture := setup(true)

	fixture.Describe("Hello")
	fixture.Finalize()

	if output := strings.TrimSpace(log.String()); output != "- Hello" {
		t.Errorf("Unexpected output: '%s'", log.String())
	}
}

func TestFinalizeAfterFailure(t *testing.T) {
	defer reset()
	log, fake, fixture := setup(false)

	fixture.Describe("Hello")
	fake.Fail()

	fixture.Finalize()

	if output := strings.TrimSpace(log.String()); output != "- Hello" {
		t.Errorf("Unexpected output: '%s'", output)
	}
}

func TestFinalizeAfterSkip_NotVerbose(t *testing.T) {
	defer reset()
	log, fake, fixture := setup(false)

	fixture.Skip("Hello")

	fixture.Finalize()

	if !fake.skipped {
		t.Error("SkipNow() was not called.")
	}
	if log.Len() > 0 {
		t.Errorf("Unexpected output: '%s'", log.String())
	}
}

func TestSoPasses(t *testing.T) {
	defer reset()
	log, fake, fixture := setup(false)

	fixture.So(true, should.BeTrue)
	fixture.Finalize()

	if log.Len() > 0 {
		t.Errorf("Unexpected ouput: '%s'", log.String())
	}
	if fake.failed {
		t.Error("Test was erroneously marked as failed.")
	}
}

func TestSoFailsAndLogs(t *testing.T) {
	defer reset()
	log, fake, fixture := setup(false)

	fixture.So(true, should.BeFalse)
	fixture.Finalize()

	if output := log.String(); !strings.Contains(output, "Expected:") {
		t.Errorf("Unexpected ouput: '%s'", log.String())
	}
	if !fake.failed {
		t.Error("Test should have been marked as failed.")
	}
}

func TestOkPasses(t *testing.T) {
	defer reset()
	log, fake, fixture := setup(false)

	fixture.Ok(true)
	fixture.Finalize()

	if log.Len() > 0 {
		t.Errorf("Unexpected ouput: '%s'", log.String())
	}
	if fake.failed {
		t.Error("Test was erroneously marked as failed.")
	}
}

func TestOkFailsAndLogs(t *testing.T) {
	defer reset()
	log, fake, fixture := setup(false)

	fixture.Ok(false)
	fixture.Finalize()

	if output := log.String(); !strings.Contains(output, "Expected condition to be true, was false instead.") {
		t.Errorf("Unexpected ouput: '%s'", log.String())
	}
	if !fake.failed {
		t.Error("Test should have been marked as failed.")
	}
}

func TestOkWithCustomMessageFailsAndLogs(t *testing.T) {
	defer reset()
	log, fake, fixture := setup(false)

	fixture.Ok(false, "gophers!")
	fixture.Finalize()

	if output := log.String(); !strings.Contains(output, "gophers!") {
		t.Errorf("Unexpected ouput: '%s'", log.String())
	}
	if !fake.failed {
		t.Error("Test should have been marked as failed.")
	}
}

func TestErrorFailsAndLogs(t *testing.T) {
	defer reset()
	log, fake, fixture := setup(false)

	fixture.Error("1", "2", "3")
	fixture.Finalize()

	if !fake.failed {
		t.Error("Test should have been marked as failed.")
	}
	if output := log.String(); !strings.Contains(output, "123") {
		t.Errorf("Expected string containing: '123' Got: '%s'", output)
	}
}

func TestErrorfFailsAndLogs(t *testing.T) {
	defer reset()
	log, fake, fixture := setup(false)

	fixture.Errorf("%s%s%s", "1", "2", "3")
	fixture.Finalize()

	if !fake.failed {
		t.Error("Test should have been marked as failed.")
	}
	if output := log.String(); !strings.Contains(output, "123") {
		t.Errorf("Expected string containing: '123' Got: '%s'", output)
	}
}

//////////////////////////////////////////////////////////////////////////////

func setup(verbosity bool) (log *bytes.Buffer, fake *FakeTT, fixture *Fixture) {
	patchVerbosity(verbosity)
	log = patchOutput()
	fake = &FakeTT{}
	fixture = NewFixture(fake)
	return log, fake, fixture
}
func patchVerbosity(verbosity bool) {
	verbose = func() bool { return verbosity }
}
func patchOutput() *bytes.Buffer {
	output := &bytes.Buffer{}
	out = output
	return output
}
func reset() {
	out = os.Stdout
	verbose = testing.Verbose
}

//////////////////////////////////////////////////////////////////////////////

type FakeTT struct {
	failed  bool
	skipped bool
}

func (self *FakeTT) Fail()        { self.failed = true }
func (self *FakeTT) Failed() bool { return self.failed }
func (self *FakeTT) SkipNow()     { self.skipped = true }

//////////////////////////////////////////////////////////////////////////////
