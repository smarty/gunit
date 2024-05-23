package gunit

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestSuccess(t *testing.T) {
	fakeT := &FakeT{buffer: &bytes.Buffer{}}
	fixture := &Fixture{TestingT: fakeT}
	fixture.So(1, shouldEqual, 1)
	actual := fakeT.buffer.String()
	expected := ""
	if actual != expected {
		t.Errorf("\n"+
			"expected: %s\n"+
			"actual:   %s", expected, actual)
	}
	if fakeT.failCount != 0 {
		t.Error("Expected 0 failures, got:", fakeT.failCount)
	}
}
func TestFailure(t *testing.T) {
	fakeT := &FakeT{buffer: &bytes.Buffer{}}
	fixture := &Fixture{TestingT: fakeT}
	fixture.So(1, shouldEqual, 2)
	actual := fakeT.buffer.String()
	expected := "shouldEqual failed: 1 vs 2\n"
	if actual != expected {
		t.Errorf("\n"+
			"expected: %s\n"+
			"actual:   %s", expected, actual)
	}
	if fakeT.failCount != 1 {
		t.Error("Expected 1 failure, got:", fakeT.failCount)
	}
}
func TestWrite(t *testing.T) {
	fakeT := &FakeT{buffer: &bytes.Buffer{}}
	fixture := &Fixture{TestingT: fakeT}
	fixture.Log("Hello, world!")
	actual := fakeT.buffer.String()
	expected := "Hello, world!"
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

func (this *FakeT) Cleanup(func())        { panic("not implemented") }
func (this *FakeT) Errorf(string, ...any) { panic("not implemented") }
func (this *FakeT) Fail()                 { panic("not implemented") }
func (this *FakeT) FailNow()              { panic("not implemented") }
func (this *FakeT) Failed() bool          { panic("not implemented") }
func (this *FakeT) Fatal(...any)          { panic("not implemented") }
func (this *FakeT) Fatalf(string, ...any) { panic("not implemented") }
func (this *FakeT) Logf(string, ...any)   { panic("not implemented") }
func (this *FakeT) Name() string          { panic("not implemented") }
func (this *FakeT) Setenv(string, string) { panic("not implemented") }
func (this *FakeT) Skip(...any)           { panic("not implemented") }
func (this *FakeT) SkipNow()              { panic("not implemented") }
func (this *FakeT) Skipf(string, ...any)  { panic("not implemented") }
func (this *FakeT) Skipped() bool         { panic("not implemented") }
func (this *FakeT) TempDir() string       { panic("not implemented") }
func (this *FakeT) Helper()               {}
func (this *FakeT) Log(a ...any) {
	_, _ = this.buffer.Write([]byte(fmt.Sprint(a...)))
}
func (this *FakeT) Error(a ...any) {
	this.failCount++
	_, _ = fmt.Fprintln(this.buffer, a...)
}

// TODO: replace with should.Equal when defined
func shouldEqual(actual any, expected ...any) error {
	if reflect.DeepEqual(actual, expected[0]) {
		return nil
	}
	return fmt.Errorf("shouldEqual failed: %v vs %v", actual, expected[0])
}
