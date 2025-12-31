package should

import (
	"runtime"
	"runtime/debug"
	"strings"
	"testing"
)

func TestStack(t *testing.T) {
	actual := stack(readLines, string(debug.Stack())) // Look for me in expected
	_, thisFile, _, _ := runtime.Caller(0)
	if !strings.Contains(actual, thisFile) {
		t.Error("Missing this file in stack trace.")
	}
	if !strings.Contains(actual, "Look for me in expected") {
		t.Error("Missing expected string 'Look for me in expected'")
	}
}
func TestStack_FailureToReadFileInStackTrace(t *testing.T) {
	actual := stack(fakeReadLines, string(debug.Stack())) // Don't look for me in expected
	_, thisFile, _, _ := runtime.Caller(0)
	if !strings.Contains(actual, thisFile) {
		t.Error("Missing this file in stack trace.")
	}
	if strings.Contains(actual, "Don't look for me in expected") {
		t.Error("Hmm, the string wasn't supposed to be there...")
	}
}
func TestStack_NoFilteredLines(t *testing.T) {
	actual := stack(fakeReadLines, "nothing to see here")
	if actual != "" {
		t.Error("Hmm, with no stack trace, no output should have been returned...")
	}
}
func fakeReadLines(string) []string {
	return nil
}
