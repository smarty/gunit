package reports

import (
	"path/filepath"
	"runtime"
	"strings"
)

func StackTrace() (frames []Frame) {
	stack := make([]uintptr, maxStackDepth)
	runtime.Callers(0, stack)
	scanner := runtime.CallersFrames(stack)
	for {
		frame, more := scanner.Next()
		if !more {
			break
		}
		frames = append(frames, Frame{
			File:     frame.File,
			Line:     frame.Line,
			Function: frame.Function,
		})
	}
	return frames
}

type Frame struct {
	Function string
	File     string
	Line     int
}

func (frame Frame) isFromStandardLibrary() bool {
	return strings.Contains(frame.File, "/libexec/src/") || // homebrew
		strings.Contains(frame.File, "/go/src/") // traditional
}

func (frame Frame) isFromGunit() bool {
	const gunitBasicExamples = "github.com/smartystreets/gunit/basic_examples"
	const gunitAdvancedExamples = "github.com/smartystreets/gunit/advanced_examples"
	const gunitFolder = "github.com/smartystreets/gunit"
	const goModuleVersionSeparator = "@" // Go module path w/ '@' separator example:
	// /Users/mike/go/pkg/mod/github.com/smartystreets/gunit@v1.0.1-0.20190705210239-badfae8b004a/reports/failure_report.go:23

	dir := filepath.Dir(frame.File)
	parts := strings.Split(dir, goModuleVersionSeparator)
	if len(parts) > 1 {
		dir = parts[0]
	}
	if strings.Contains(dir, gunitBasicExamples) || strings.Contains(dir, gunitAdvancedExamples) {
		return false
	}
	return strings.Contains(dir, gunitFolder)
}
