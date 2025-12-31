package gunit

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/smarty/gunit/v2/assert/should"
)

func TestShouldFailure(t *testing.T) {
	fakeT := &FakeT{buffer: &bytes.Buffer{}}
	fixture := &Fixture{TestingT: fakeT}
	fixture.So(1, should.Equal, 2)
	actual := strings.Join(strings.Fields(fakeT.buffer.String()), " ")
	expected := "assertion failure: Expected: (int) 2 Actual: (int) 1 ^ Stack: (filtered)"
	if !strings.HasPrefix(actual, expected) {
		t.Errorf("\n"+
			"expected: %s\n"+
			"actual:   %s", expected, actual)
	}
	if fakeT.failCount != 1 {
		t.Error("Expected 1 failure, got:", fakeT.failCount)
	}
}
func TestPrint(t *testing.T) {
	fakeT := &FakeT{buffer: &bytes.Buffer{}}
	fixture := &Fixture{TestingT: fakeT}
	fixture.Print("one")
	fixture.Printf("%s", "two")
	fixture.Println("three")
	actual := fakeT.buffer.String()
	expected := "one" + "two" + "three\n"
	if actual != expected {
		t.Errorf("\n"+
			"expected: %s\n"+
			"actual:   %s", expected, actual)
	}
}

type FakeT struct {
	failCount int
	buffer    *bytes.Buffer
}

func (this *FakeT) Cleanup(func())           { panic("not implemented") }
func (this *FakeT) Context() context.Context { panic("not implemented") }
func (this *FakeT) Errorf(string, ...any)    { panic("not implemented") }
func (this *FakeT) Fail()                    { panic("not implemented") }
func (this *FakeT) FailNow()                 { panic("not implemented") }
func (this *FakeT) Failed() bool             { panic("not implemented") }
func (this *FakeT) Fatal(a ...any)           { panic("not implemented") }
func (this *FakeT) Fatalf(string, ...any)    { panic("not implemented") }
func (this *FakeT) Logf(string, ...any)      { panic("not implemented") }
func (this *FakeT) Name() string             { panic("not implemented") }
func (this *FakeT) Output() io.Writer        { return this.buffer }
func (this *FakeT) Setenv(string, string)    { panic("not implemented") }
func (this *FakeT) Skip(...any)              { panic("not implemented") }
func (this *FakeT) SkipNow()                 { panic("not implemented") }
func (this *FakeT) Skipf(string, ...any)     { panic("not implemented") }
func (this *FakeT) Skipped() bool            { panic("not implemented") }
func (this *FakeT) TempDir() string          { panic("not implemented") }
func (this *FakeT) Helper()                  {}
func (this *FakeT) Log(a ...any) {
	_, _ = this.buffer.Write([]byte(fmt.Sprint(a...)))
}
func (this *FakeT) Error(a ...any) {
	this.failCount++
	_, _ = fmt.Fprintln(this.buffer, a...)
}
