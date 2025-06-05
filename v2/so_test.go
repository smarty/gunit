package gunit_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/smarty/gunit/v2"
	"github.com/smarty/gunit/v2/must"
	"github.com/smarty/gunit/v2/should"
)

func TestSoSuccess(t *testing.T) {
	fakeT := &FakeT{buffer: &bytes.Buffer{}}
	gunit.So(fakeT, 1, should.Equal, 1)
	if fakeT.buffer.String() != "" {
		t.Errorf("\n"+
			"expected: <blank>\n"+
			"actual:   %s", fakeT.buffer.String())
	}
}
func TestSoFailure(t *testing.T) {
	fakeT := &FakeT{buffer: &bytes.Buffer{}}
	gunit.So(fakeT, 1, should.Equal, 2)
	actual := strings.Join(strings.Fields(fakeT.buffer.String()), " ")
	expected := `assertion failure: Expected: (int) 2 Actual: (int) 1 ^ Stack: (filtered)`
	if !strings.HasPrefix(actual, expected) {
		t.Errorf("\n"+
			"expected: %s\n"+
			"actual:   %s", expected, actual)
	}
	if fakeT.failCount != 1 {
		t.Error("Expected 1 fatal failure, got:", fakeT.failCount)
	}
}
func TestSoFatalSuccess(t *testing.T) {
	fakeT := &FakeT{buffer: &bytes.Buffer{}}
	gunit.So(fakeT, 1, must.Equal, 1)
	if fakeT.buffer.String() != "" {
		t.Errorf("\n"+
			"expected: <blank>\n"+
			"actual:   %s", fakeT.buffer.String())
	}
}
func TestSoFatalFailure(t *testing.T) {
	fakeT := &FakeT{buffer: &bytes.Buffer{}}
	gunit.So(fakeT, 1, must.Equal, 2)
	actual := strings.Join(strings.Fields(fakeT.buffer.String()), " ")
	expected := "fatal assertion failure: Expected: (int) 2 Actual: (int) 1 ^ Stack: (filtered)"
	if !strings.HasPrefix(actual, expected) {
		t.Errorf("\n"+
			"expected: %s\n"+
			"actual:   %s", expected, actual)
	}
	if fakeT.fatalCount != 1 {
		t.Error("Expected 1 fatal failure, got:", fakeT.fatalCount)
	}
}

type FakeT struct {
	buffer     *bytes.Buffer
	failCount  int
	fatalCount int
}

func (this *FakeT) Helper() {}
func (this *FakeT) Error(a ...any) {
	this.failCount++
	_, _ = fmt.Fprint(this.buffer, a...)
}
func (this *FakeT) Fatal(a ...any) {
	this.fatalCount++
	_, _ = fmt.Fprint(this.buffer, a...)
}
