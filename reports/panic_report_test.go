package reports

import "testing"

func TestPanicReport(t *testing.T) {
	actual := PanicReport(examplePanic, exampleStackTrace)
	if actual != expectedPanicReport {
		t.Errorf("Incorrect panic report.\nGot:\n%s\n\nWant:\n%s", actual, expectedPanicReport)
	}
}

const expectedPanicReport = `PANIC: runtime error: invalid memory address or nil pointer dereference
...
github.com/smartystreets/gunit/advanced_examples.(*Game).Roll(...)
	/Users/mike/src/github.com/smartystreets/gunit/advanced_examples/bowling_game.go:13
github.com/smartystreets/gunit/advanced_examples.(*BowlingGameScoringFixture).rollMany(0xc000091310, 0x14, 0x0)
	/Users/mike/src/github.com/smartystreets/gunit/advanced_examples/bowling_game_test.go:58 +0x38
github.com/smartystreets/gunit/advanced_examples.(*BowlingGameScoringFixture).TestAfterAllGutterBallsTheScoreShouldBeZero(0xc000091310)
	/Users/mike/src/github.com/smartystreets/gunit/advanced_examples/bowling_game_test.go:23 +0x3d`

const examplePanic = "runtime error: invalid memory address or nil pointer dereference"

/*
Consider the stack trace below. We want to filter out all lines
except those that represent actual program behavior (middle block),
not the runtime/gunit code that handles the panic (shown in the first
block) or the reflect/gunit/testing code (shown in the last block).
*/

var exampleStackTrace = []byte(`goroutine 22 [running]:
runtime/debug.Stack(0x11ec660, 0xc0000b7fa0, 0xc0000fd888)
	/usr/local/go/src/runtime/debug/stack.go:24 +0x9d
github.com/smartystreets/gunit/reports.PanicReport(0x1182b60, 0x12fe320, 0x10674a0, 0x1306400)
	/Users/mike/src/github.com/smartystreets/gunit/reports/panic_report.go:15 +0x13e
github.com/smartystreets/gunit.(*Fixture).recoverPanic(0xc0000b7f60, 0x1182b60, 0x12fe320)
	/Users/mike/src/github.com/smartystreets/gunit/fixture.go:106 +0x4f
github.com/smartystreets/gunit.(*Fixture).finalize(0xc0000b7f60)
	/Users/mike/src/github.com/smartystreets/gunit/fixture.go:97 +0x211
panic(0x1182b60, 0x12fe320)
	/usr/local/go/src/runtime/panic.go:679 +0x1b2
github.com/smartystreets/gunit/advanced_examples.(*Game).Roll(...)
	/Users/mike/src/github.com/smartystreets/gunit/advanced_examples/bowling_game.go:13
github.com/smartystreets/gunit/advanced_examples.(*BowlingGameScoringFixture).rollMany(0xc000091310, 0x14, 0x0)
	/Users/mike/src/github.com/smartystreets/gunit/advanced_examples/bowling_game_test.go:58 +0x38
github.com/smartystreets/gunit/advanced_examples.(*BowlingGameScoringFixture).TestAfterAllGutterBallsTheScoreShouldBeZero(0xc000091310)
	/Users/mike/src/github.com/smartystreets/gunit/advanced_examples/bowling_game_test.go:23 +0x3d
reflect.Value.call(0x11af5e0, 0xc000091310, 0x3e13, 0x11b8925, 0x4, 0x0, 0x0, 0x0, 0x15, 0xc00009fe40, ...)
	/usr/local/go/src/reflect/value.go:460 +0x5f6
reflect.Value.Call(0x11af5e0, 0xc000091310, 0x3e13, 0x0, 0x0, 0x0, 0x3e13, 0x0, 0x0)
	/usr/local/go/src/reflect/value.go:321 +0xb4
github.com/smartystreets/gunit.(*testCase).runTest(0xc00012c1c0)
	/Users/mike/src/github.com/smartystreets/gunit/test_case.go:86 +0x7c
github.com/smartystreets/gunit.(*testCase).runWithSetupAndTeardown(0xc00012c1c0)
	/Users/mike/src/github.com/smartystreets/gunit/test_case.go:76 +0x69
github.com/smartystreets/gunit.(*testCase).run(0xc00012c1c0, 0xc0000d4400)
	/Users/mike/src/github.com/smartystreets/gunit/test_case.go:64 +0x81
testing.tRunner(0xc0000d4400, 0xc000091140)
	/usr/local/go/src/testing/testing.go:909 +0xc9
created by testing.(*T).Run
	/usr/local/go/src/testing/testing.go:960 +0x350`)
