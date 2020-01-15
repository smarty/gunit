package reports

import (
	"fmt"
	"runtime/debug"
	"strings"
)

func PanicReport(r interface{}) string {
	var builder strings.Builder
	fmt.Fprintln(&builder, "PANIC:", r)
	fmt.Fprintln(&builder, "...")
	for _, line := range strings.Split(string(debug.Stack()), "\n") {
		if panicLineIsFromGoRuntime(line) || panicLineIsFromGunit(line) {
			continue
		}
		fmt.Fprintln(&builder, line)
	}
	return strings.TrimSpace(builder.String())
}

func panicLineIsFromGunit(line string) bool {
	if strings.Contains(line, "github.com/smartystreets/gunit/fixture.go") {
		return true
	}
	if strings.Contains(line, "github.com/smartystreets/gunit.(*") {
		return true
	}
	if strings.Contains(line, "github.com/smartystreets/gunit/test_case.go") {
		return true
	}
	if strings.Contains(line, "github.com/smartystreets/gunit/reports/panic_report.go") {
		return true
	}
	if strings.Contains(line, "github.com/smartystreets/gunit/reports.PanicReport(") {
		return true
	}
	if strings.Contains(line, "github.com/smartystreets/gunit.") {
		return true
	}
	return false
}
func panicLineIsFromGoRuntime(line string) bool {
	if strings.Contains(line, "go/src") {
		return true
	}
	if strings.Contains(line, "/libexec/src/") {
		return true
	}
	if strings.Contains(line, "reflect.Value") {
		return true
	}
	if strings.Contains(line, "testing.tRunner") {
		return true
	}
	if strings.Contains(line, "testing.(*T).Run") {
		return true
	}
	if strings.Contains(line, "goroutine") && strings.Contains(line, "[running]:") {
		return true
	}
	if strings.Contains(line, "runtime/debug.Stack(") {
		return true
	}
	if strings.Contains(line, "panic(0x") {
		return true
	}
	return false
}
