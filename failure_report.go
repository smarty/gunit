package gunit

import (
	"bytes"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/smartystreets/gunit/scan"
)

type failureReport struct {
	Stack   []string
	Method  string
	Fixture string
	Package string
	Failure string
}

func newFailureReport(failure string) string {
	report := &failureReport{Failure: failure}
	report.ScanStack()
	return report.String()
}

func (this *failureReport) ScanStack() {
	stack := make([]uintptr, maxStackDepth)
	runtime.Callers(0, stack)
	frames := runtime.CallersFrames(stack)
	for {
		frame, more := frames.Next()
		if !more {
			break
		}
		if isFromStandardLibrary(frame) || isFromGunit(frame) {
			continue
		}
		this.ParseTestName(frame.Function)
		this.Stack = append(this.Stack, fmt.Sprintf("%s:%d", frame.File, frame.Line))
	}
}

func isFromGunit(frame runtime.Frame) bool {
	const gunitFolder = "github.com/smartystreets/gunit"
	const goModuleVersionSeparator = "@" // Go module path w/ '@' separator example:
	// /Users/mike/go/pkg/mod/github.com/smartystreets/gunit@v1.0.1-0.20190705210239-badfae8b004a/failure_report.go:23

	dir := filepath.Dir(frame.File)
	parts := strings.Split(dir, goModuleVersionSeparator)
	if len(parts) > 1 {
		dir = parts[0]
	}
	return strings.HasSuffix(dir, gunitFolder)
}

func isFromStandardLibrary(frame runtime.Frame) bool {
	return strings.Contains(frame.File, "libexec/src/")
}

func (this *failureReport) ParseTestName(name string) {
	if len(this.Method) > 0 {
		return
	}
	parts := strings.Split(name, ".")
	partCount := len(parts)
	last := partCount - 1
	if partCount < 3 {
		return
	}

	if method := parts[last]; scan.IsTestCase(method) {
		this.Method = method
		this.Fixture = parts[last-1]
		this.Package = strings.Join(parts[0:last-1], ".")
	}
}

func (this failureReport) String() string {
	buffer := new(bytes.Buffer)
	for i, stack := range this.Stack {
		fmt.Fprintf(buffer, "(%d):      %s\n", len(this.Stack)-i-1, stack)
	}
	fmt.Fprintf(buffer, this.Failure)
	return buffer.String() + "\n\n"
}

const maxStackDepth = 32
