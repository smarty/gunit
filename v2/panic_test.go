package gunit

import (
	"strings"
	"testing"
)

func TestPanicReport(t *testing.T) {
	actualStackTrace := panicReport("boink", []byte(rawStackTrace))
	if actualStackTrace != expectedStackTrace {
		t.Error(actualStackTrace)
	}
}

var rawStackTrace = strings.TrimSpace(`
runtime/debug.Stack()
	/usr/local/go/src/runtime/debug/stack.go:24 +0x5e
github.com/smarty/gunit/v2/should.testCase.runTest.func1()
	/Users/mike/src/github.com/smarty/gunit/v2/should/run.go:139 +0xce
panic({0x10720340?, 0x10762000?})
	/usr/local/go/src/runtime/panic.go:770 +0x132
github.com/smarty/gunit/v2/examples/bowling.(*game).calculateScore(...)
	/Users/mike/src/github.com/smarty/gunit/v2/examples/bowling/bowling.go:15
github.com/smarty/gunit/v2/examples/bowling.(*GameFixture).assertScore(...)
	/Users/mike/src/github.com/smarty/gunit/v2/examples/bowling/bowling_test.go:22
github.com/smarty/gunit/v2/examples/bowling.(*GameFixture).TestAllOnes(0x0?)
	/Users/mike/src/github.com/smarty/gunit/v2/examples/bowling/bowling_test.go:40 +0x5b
reflect.Value.call({0x1075d540?, 0xc000025100?, 0xc000025100?}, {0x106bf18b, 0x4}, {0x0, 0x0, 0x10766658?})
	/usr/local/go/src/reflect/value.go:596 +0xce5
reflect.Value.Call({0x1075d540?, 0xc000025100?, 0x603?}, {0x0?, 0xc000054678?, 0x10897a60?})
	/usr/local/go/src/reflect/value.go:380 +0xb9
github.com/smarty/gunit/v2/should.testCase.runTest({0xc000108680, {0x10710a0c, 0xb}, 0xc000012330, 0x0, {0x10766658, 0x1075d540}, {0x1075d540, 0xc000024350, 0x16}}, ...)
	/Users/mike/src/github.com/smarty/gunit/v2/should/run.go:162 +0x2b1
testing.tRunner(0xc000108ea0, 0xc00002e720)
	/usr/local/go/src/testing/testing.go:1689 +0xfb
created by testing.(*T).Run in goroutine 6
	/usr/local/go/src/testing/testing.go:1742 +0x390
`)

var expectedStackTrace = strings.TrimSpace(`
PANIC: boink
...
github.com/smarty/gunit/v2/examples/bowling.(*game).calculateScore(...)
	/Users/mike/src/github.com/smarty/gunit/v2/examples/bowling/bowling.go:15
github.com/smarty/gunit/v2/examples/bowling.(*GameFixture).assertScore(...)
	/Users/mike/src/github.com/smarty/gunit/v2/examples/bowling/bowling_test.go:22
github.com/smarty/gunit/v2/examples/bowling.(*GameFixture).TestAllOnes(0x0?)
	/Users/mike/src/github.com/smarty/gunit/v2/examples/bowling/bowling_test.go:40 +0x5b
`)
