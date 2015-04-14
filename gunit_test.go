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
	patchVerbosity(false)
	log := patchOutput()
	fake := &FakeTT{}
	fixture := NewFixture(fake)

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
	patchVerbosity(false)
	log := patchOutput()
	fake := &FakeTT{}
	fixture := NewFixture(fake)
	fixture.Describe("Hello")
	fixture.Finalize()

	if output := strings.TrimSpace(log.String()); output != "" {
		t.Errorf("Unexpected output: '%s'", log.String())
	}
}

func TestFinalizeAfterPass_Verbose(t *testing.T) {
	defer reset()
	patchVerbosity(true)
	log := patchOutput()
	fake := &FakeTT{}
	fixture := NewFixture(fake)
	fixture.Describe("Hello")
	fixture.Finalize()

	if output := strings.TrimSpace(log.String()); output != "- Hello" {
		t.Errorf("Unexpected output: '%s'", log.String())
	}
}

func TestFinalizeAfterFailure(t *testing.T) {
	defer reset()
	patchVerbosity(false)
	log := patchOutput()
	fake := &FakeTT{}
	fixture := NewFixture(fake)

	fixture.Describe("Hello")
	fake.Fail()

	fixture.Finalize()

	if output := strings.TrimSpace(log.String()); output != "- Hello" {
		t.Errorf("Unexpected output: '%s'", output)
	}
}

func TestFinalizeAfterSkip_NotVerbose(t *testing.T) {
	defer reset()
	patchVerbosity(false)
	log := patchOutput()
	fake := &FakeTT{}
	fixture := NewFixture(fake)

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
	patchVerbosity(false)
	log := patchOutput()
	fake := &FakeTT{}
	fixture := NewFixture(fake)
	fixture.So(true, should.BeTrue)
	fixture.Finalize()

	if log.Len() > 0 {
		t.Errorf("Unexpected ouput: '%s'", log.String())
	}
	if fake.failed {
		t.Error("Test was erroneously marked as failed.")
	}
}

func TestSoFails(t *testing.T) {
	defer reset()
	patchVerbosity(false)
	log := patchOutput()
	fake := &FakeTT{}
	fixture := NewFixture(fake)
	fixture.So(true, should.BeFalse)
	fixture.Finalize()

	if output := log.String(); !strings.Contains(output, "Expected:") {
		t.Errorf("Unexpected ouput: '%s'", log.String())
	}
	if !fake.failed {
		t.Error("Test should have been marked as failed.")
	}
}

//////////////////////////////////////////////////////////////////////////////

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
