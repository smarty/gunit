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
	ErrExpectedCountInvalid = errors.New("expected count invalid")
	ErrTypeMismatch         = errors.New("type mismatch")
	ErrKindMismatch         = errors.New("kind mismatch")
	ErrAssertionFailure     = errors.New("assertion failure")
)

func failure(format string, args ...any) error {
	trace := stack()
	if len(trace) > 0 {
		format += "\nStack: (filtered)\n%s"
		args = append(args, trace)
	}
	return wrap(ErrAssertionFailure, format, args...)
}
func stack() string {
	lines := strings.Split(string(debug.Stack()), "\n")
	var filtered []string
	for x := 1; x < len(lines)-1; x += 2 {
		fileLineRaw := lines[x+1]
		if strings.Contains(fileLineRaw, "_test.go:") {
			filtered = append(filtered, lines[x], fileLineRaw)
			line, ok := readSourceCodeLine(fileLineRaw)
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
func readSourceCodeLine(fileLineRaw string) (string, bool) {
	fileLineJoined := strings.Fields(strings.TrimSpace(fileLineRaw))[0]
	fileLine := strings.Split(fileLineJoined, ":")
	sourceCode, _ := os.ReadFile(fileLine[0])
	sourceCodeLines := strings.Split(string(sourceCode), "\n")
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
