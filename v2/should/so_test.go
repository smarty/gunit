package should_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/smarty/gunit/v2/should"
)

func TestSoSuccess(t *testing.T) {
	fakeT := &FakeT{buffer: &bytes.Buffer{}}
	should.So(fakeT, 1, should.Equal, 1)
	if fakeT.buffer.String() != "" {
		t.Errorf("\n"+
			"expected: <blank>\n"+
			"actual:   %s", fakeT.buffer.String())
	}
}
func TestSoFailure(t *testing.T) {
	fakeT := &FakeT{buffer: &bytes.Buffer{}}
	should.So(fakeT, 1, should.Equal, 2)
	actual := strings.Join(strings.Fields(fakeT.buffer.String()), " ")
	expected := `assertion failure: Expected: (int) 2 Actual: (int) 1 ^ Stack: (filtered)`
	if !strings.HasPrefix(actual, expected) {
		t.Errorf("\n"+
			"expected: %s\n"+
			"actual:   %s", expected, actual)
	}
}

type FakeT struct {
	buffer *bytes.Buffer
}

func (this *FakeT) Helper() {}
func (this *FakeT) Error(a ...any) {
	_, _ = fmt.Fprint(this.buffer, a...)
}
