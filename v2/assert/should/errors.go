package should

import (
	"errors"
	"fmt"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
)

var (
	ErrExpectedCountInvalid  = errors.New("expected count invalid")
	ErrTypeMismatch          = errors.New("type mismatch")
	ErrKindMismatch          = errors.New("kind mismatch")
	ErrAssertionFailure      = errors.New("assertion failure")
	ErrFatalAssertionFailure = errors.New("fatal")
)

func failure(format string, args ...any) error {
	trace := stack(readLines, string(debug.Stack()))
	if len(trace) > 0 {
		format += "\nStack: (filtered)\n%s"
		args = append(args, trace)
	}
	return wrap(ErrAssertionFailure, format, args...)
}
func stack(readLines func(string) []string, stackTrace string) string {
	lines := strings.Split(stackTrace, "\n")
	var filtered []string
	for x := 1; x < len(lines)-1; x += 2 {
		fileLineRaw := lines[x+1]
		if strings.Contains(fileLineRaw, "_test.go:") {
			filtered = append(filtered, lines[x], fileLineRaw)
			line, ok := readSourceCodeLine(readLines, fileLineRaw)
			if ok {
				filtered = append(filtered, "  "+line)
			}

		}
	}
	if len(filtered) == 0 {
		return ""
	}
	return "> " + strings.Join(filtered, "\n> ")
}
func readLines(path string) []string {
	content, _ := os.ReadFile(path)
	return strings.Split(string(content), "\n")
}
func readSourceCodeLine(readLines func(string) []string, fileLineRaw string) (string, bool) {
	fileLineJoined := strings.Fields(strings.TrimSpace(fileLineRaw))[0]
	fileLine := strings.Split(fileLineJoined, ":")
	sourceCodeLines := readLines(fileLine[0])
	lineNumber, _ := strconv.Atoi(fileLine[1])
	lineNumber--
	if len(sourceCodeLines) <= lineNumber {
		return "", false
	}
	return sourceCodeLines[lineNumber], true
}
func wrap(inner error, format string, args ...any) error {
	return fmt.Errorf("%w: "+fmt.Sprintf(format, args...), inner)
}
