package reports

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/smartystreets/gunit/scan"
)

func FailureReport(failure string, stack []Frame) string {
	report := newFailureReport(failure, readFile)
	report.scanStack(stack)
	return report.composeReport()
}

type failureReport struct {
	stack  []string
	files  map[string][]string
	reader fileReader

	method   string
	fixture  string
	package_ string
	failure  string
}

func newFailureReport(failure string, reader fileReader) *failureReport {
	return &failureReport{
		failure: failure,
		files:   make(map[string][]string),
		reader:  reader,
	}
}

func (this *failureReport) scanStack(stack []Frame) {
	for _, frame := range stack {
		if frame.isFromStandardLibrary() || frame.isFromGunit() {
			continue
		}
		this.parseTestName(frame.Function)
		this.loadFile(frame)
		code := this.extractLineOfCode(frame)
		filename := filepath.Base(frame.File)
		stack := fmt.Sprintf("%s // %s:%d", code, filename, frame.Line)
		this.stack = append(this.stack, strings.TrimSpace(stack))
	}
}

func (this *failureReport) loadFile(frame Frame) {
	if _, found := this.files[frame.File]; !found {
		this.files[frame.File] = strings.Split(this.reader(frame.File), "\n")
	}
}

func (this *failureReport) extractLineOfCode(frame Frame) string {
	file := this.files[frame.File]
	if len(file) < frame.Line {
		return ""
	}
	return strings.TrimSpace(file[frame.Line-1])
}

func (this *failureReport) parseTestName(name string) {
	if len(this.method) > 0 {
		return
	}
	parts := strings.Split(name, ".")
	partCount := len(parts)
	last := partCount - 1
	if partCount < 3 {
		return
	}

	if method := parts[last]; scan.IsTestCase(method) {
		this.method = method
		this.fixture = parts[last-1]
		this.package_ = strings.Join(parts[0:last-1], ".")
	}
}

func (this failureReport) composeReport() string {
	buffer := new(bytes.Buffer)
	for i, stack := range this.stack {
		fmt.Fprintf(buffer, "(%d): %s\n", len(this.stack)-i-1, stack)
	}
	fmt.Fprintf(buffer, this.failure)
	return buffer.String() + "\n\n"
}

const maxStackDepth = 32

type fileReader func(path string) string

func readFile(path string) string {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(raw)
}
