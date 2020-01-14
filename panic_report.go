package gunit

import (
	"runtime/debug"
	"strings"
)

func panicReport() string {
	stack := strings.Split(string(debug.Stack()), "\n")
	var filtered []string
	filtered = append(filtered, "...")
	for _, line := range stack {
		if panicLineIsFromGoRuntime(line) || panicLineIsFromGunit(line) {
			continue
		}
		filtered = append(filtered, line)
	}
	return strings.TrimSpace(strings.Join(filtered, "\n"))
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
	if strings.Contains(line, "github.com/smartystreets/gunit/panic_report.go") {
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
