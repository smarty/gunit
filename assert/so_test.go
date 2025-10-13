package assert_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/smarty/gunit/assert"
	"github.com/smarty/gunit/assert/better"
	"github.com/smarty/gunit/assert/should"
)

func TestSoSuccess(t *testing.T) {
	fakeT := &FakeT{buffer: &bytes.Buffer{}}
	assert.So(fakeT, 1, should.Equal, 1)
	if fakeT.buffer.String() != "" {
		t.Errorf("\n"+
			"expected: <blank>\n"+
			"actual:   %s", fakeT.buffer.String(),
		)
	}
}
func TestSoFailure(t *testing.T) {
	fakeT := &FakeT{buffer: &bytes.Buffer{}}
	assert.So(fakeT, 1, should.Equal, 2)
	actual := strings.Join(strings.Fields(fakeT.buffer.String()), " ")
	expected := `Expected: 2 Actual: 1 (Should equal)!`
	if !strings.HasPrefix(actual, expected) {
		t.Errorf("\n"+
			"expected: %s\n"+
			"actual:   %s", expected, actual)
	}
	if fakeT.failCount != 1 {
		t.Error("Expected 1 failure, got:", fakeT.failCount)
	}
}
func TestSoFatal(t *testing.T) {
	fakeT := &FakeT{buffer: &bytes.Buffer{}}
	assert.So(fakeT, 1, better.Equal, 2)
	actual := strings.Join(strings.Fields(fakeT.buffer.String()), " ")
	expected := "<<<FATAL>>> Expected: 2 Actual: 1 (Should equal)!"
	if !strings.HasPrefix(actual, expected) {
		t.Errorf("\n"+
			"expected: %s\n"+
			"actual:   %s", expected, actual)
	}
	if fakeT.fatalCount != 1 {
		t.Error("Expected 1 fatal failure, got:", fakeT.failCount)
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
