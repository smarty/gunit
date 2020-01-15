package reports

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/smartystreets/gunit/scan"
)

type failureReport struct {
	Stack []string
	Files map[string][]string

	Method  string
	Fixture string
	Package string
	Failure string
}

func FailureReport(failure string) string {
	report := &failureReport{Failure: failure, Files: make(map[string][]string)}
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
		this.LoadFile(frame)
		code := this.extractLineOfCode(frame)
		filename := filepath.Base(frame.File)
		stack := fmt.Sprintf("%s // %s:%d", code, filename, frame.Line)
		this.Stack = append(this.Stack, strings.TrimSpace(stack))
	}
}

func (this *failureReport) LoadFile(frame runtime.Frame) {
	if _, found := this.Files[frame.File]; !found {
		this.Files[frame.File] = readLines(frame.File)
	}
}
func readLines(path string) []string {
	all, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	return strings.Split(string(all), "\n")
}

func (this *failureReport) extractLineOfCode(frame runtime.Frame) string {
	file := this.Files[frame.File]
	if len(file) < frame.Line {
		return ""
	}
	return strings.TrimSpace(file[frame.Line-1])
}

func isFromGunit(frame runtime.Frame) bool {
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

func isFromStandardLibrary(frame runtime.Frame) bool {
	return strings.Contains(frame.File, "/libexec/src/") || // homebrew
		strings.Contains(frame.File, "/go/src/") // traditional
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
		fmt.Fprintf(buffer, "(%d): %s\n", len(this.Stack)-i-1, stack)
	}
	fmt.Fprintf(buffer, this.Failure)
	return buffer.String() + "\n\n"
}

const maxStackDepth = 32
