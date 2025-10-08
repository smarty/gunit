package assert_test

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/smarty/gunit/assert"
)

func TestSoSuccess(t *testing.T) {
	fakeT := &FakeT{buffer: &bytes.Buffer{}}
	assert.So(fakeT, 1, shouldEqual, 1)
	if fakeT.buffer.String() != "" {
		t.Errorf("\n"+
			"expected: <blank>\n"+
			"actual:   %s", fakeT.buffer.String(),
		)
	}
}
func TestSoFailure(t *testing.T) {
	fakeT := &FakeT{buffer: &bytes.Buffer{}}
	assert.So(fakeT, 1, shouldEqual, 2)
	actual := strings.Join(strings.Fields(fakeT.buffer.String()), " ")
	expected := `1 != 2`
	if !strings.HasPrefix(actual, expected) {
		t.Errorf("\n"+
			"expected: %s\n"+
			"actual:   %s", expected, actual)
	}
	if fakeT.failCount != 1 {
		t.Error("Expected 1 fatal failure, got:", fakeT.failCount)
	}
}

func shouldEqual(actual any, expected ...any) string {
	if reflect.DeepEqual(actual, expected[0]) {
		return ""
	}
	return fmt.Sprintf("%v != %v", actual, expected[0])
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
