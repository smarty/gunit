package gunit

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestPassingAssertion(t *testing.T) {
	// not marked as failure
	// no log output
	t.Skip("Not yet implemented")
}

func TestFailingAssertion(t *testing.T) {
	// marked as failure
	// log output not empty
	t.Skip("Not yet implemented")
}

func TestFinalize(t *testing.T) {
	// forwards the call to the inner T.
	t.Skip("Not yet implemented")
}

func TestErrorMarksFailure(t *testing.T) {
	fake := NewFakeT()
	wrapper := NewTWrapper(fake)

	wrapper.Error("hello", "world")

	if !fake.failed {
		t.Error("Calling .Error(...) should have marked the test as failed (but it didn't).")
	}
	if actual := strings.TrimSpace(wrapper.log.String()); actual != "hello world" {
		t.Errorf("Expected log: '	hello world' Actual log: '%s'", actual)
	}
}

func TestErrorfMarksFailure(t *testing.T) {
	fake := NewFakeT()
	wrapper := NewTWrapper(fake)

	wrapper.Errorf("hello %s", "world")

	if !fake.failed {
		t.Error("Calling .Errorf(...) should have marked the test as failed (but it didn't).")
	}
	if actual := strings.TrimSpace(wrapper.log.String()); actual != "hello world" {
		t.Errorf("Expected log: '	hello world' Actual log: '%s'", actual)
	}
}

func TestSkip_OnlyCallsSkipNowAtFinalize(t *testing.T) {
	fake := NewFakeT()
	wrapper := NewTWrapper(fake)

	wrapper.Skip()

	if fake.skipped {
		t.Error("SkipNow should only be called after finalize!")
		t.FailNow()
	}

	wrapper.finalize()

	if !fake.skipped {
		t.Error("SkipNow should have been called by finalize!")
	}

	if fake.failed {
		t.Error("The test was erroneously marked as failed.")
	}
}

func TestSkipf_OnlyCallsSkipNowAtFinalize(t *testing.T) {
	fake := NewFakeT()
	wrapper := NewTWrapper(fake)

	wrapper.Skipf("hi %s", "world")

	if fake.skipped {
		t.Error("SkipNow should only be called after finalize!")
		t.FailNow()
	}

	wrapper.finalize()

	if !fake.skipped {
		t.Error("SkipNow should have been called by finalize!")
	}

	if fake.failed {
		t.Error("The test was erroneously marked as failed.")
	}
}

func TestWhenLogIsEmpty_NoOutputIsGenerated(t *testing.T) {
	out = &bytes.Buffer{}
	fake := NewFakeT()
	wrapper := NewTWrapper(fake)

	wrapper.finalize()

	if log := out.(fmt.Stringer).String(); log != "" {
		t.Errorf("Expected a blank log, got: '%s'", log)
	}
}

func TestWhenLogIsFull_ButNoVerboseMode_AndNoFailure(t *testing.T) {
	out = &bytes.Buffer{}
	fake := NewFakeT()
	wrapper := NewTWrapper(fake)
	wrapper.Log("HI")

	wrapper.finalize()

	if log := out.(fmt.Stringer).String(); log != "" {
		t.Errorf("Expected a blank log, got: '%s'", log)
	}
}

func TestWhenLogIsFullAndVerboseModeIsOn_GenerateOutput(t *testing.T) {
	out = &bytes.Buffer{}
	fake := NewFakeT()
	wrapper := NewTWrapper(fake)
	verbose = true
	wrapper.Log("hi")

	wrapper.finalize()

	if log := out.(fmt.Stringer).String(); log == "" {
		t.Error("Log was empty!")
	}
}

func TestWhenLogIsFullAndTestFailed_GenerateOutput(t *testing.T) {
	out = &bytes.Buffer{}
	fake := NewFakeT()
	wrapper := NewTWrapper(fake)
	wrapper.Log("hi")
	wrapper.Fail()

	wrapper.finalize()

	if log := out.(fmt.Stringer).String(); log == "" {
		t.Error("Log was empty!")
	}
}

//////////////////////////////////////////////////////////////////////////////

type FakeT struct {
	skipped   bool
	failed    bool
	log       *bytes.Buffer
	finalized bool
}

func NewFakeT() *FakeT                      { return &FakeT{log: &bytes.Buffer{}} }
func (self *FakeT) SkipNow()                { self.skipped = true }
func (self *FakeT) Fail()                   { self.failed = true }
func (self *FakeT) Failed() bool            { return self.failed }
func (self *FakeT) Log(args ...interface{}) { self.log.WriteString(fmt.Sprintln(args...)) }
func (self *FakeT) finalize()               { self.finalized = true }
func init() {
	out = ioutil.Discard // NOTE: if you aren't seeing debug output, this is why...
}
