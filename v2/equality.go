package gunit

import (
	"errors"
	"fmt"
	"math"
	"os"
	"reflect"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

var errAssertionFailure = errors.New("assertion failure")

// equal verifies that the actual value is equal to the expected value.
// It uses reflect.DeepEqual in most cases, but also compares numerics
// regardless of specific type and compares time.Time values using the
// time.Equal method.
func equal(a, b any) error {
	for _, spec := range equalitySpecs {
		if !spec.assertable(a, b) {
			continue
		}
		if spec.passes(a, b) {
			return nil
		}
		break
	}
	return failure(report(a, b))
}

type specification interface {
	assertable(a, b any) bool
	passes(a, b any) bool
}

var equalitySpecs = []specification{
	numericEquality{},
	timeEquality{},
	deepEquality{},
}

func report(a, b any) string {
	aType := fmt.Sprintf("(%v)", reflect.TypeOf(a))
	bType := fmt.Sprintf("(%v)", reflect.TypeOf(b))
	longestType := int(math.Max(float64(len(aType)), float64(len(bType))))
	aType += strings.Repeat(" ", longestType-len(aType))
	bType += strings.Repeat(" ", longestType-len(bType))
	aFormat := fmt.Sprintf(format(a), a)
	bFormat := fmt.Sprintf(format(b), b)
	typeDiff := diff(aType, bType)
	valueDiff := diff(aFormat, bFormat)

	builder := new(strings.Builder)
	_, _ = fmt.Fprintf(builder, "\n")
	_, _ = fmt.Fprintf(builder, "a. %s %s\n", aType, aFormat)
	_, _ = fmt.Fprintf(builder, "b. %s %s\n", bType, bFormat)
	_, _ = fmt.Fprintf(builder, "          %s %s", typeDiff, valueDiff)

	if firstDiffIndex := strings.Index(valueDiff, "^"); firstDiffIndex > 40 {
		start := firstDiffIndex - 20
		_, _ = fmt.Fprintf(builder, "\nInitial discrepancy at index %d:\n", firstDiffIndex)
		_, _ = fmt.Fprintf(builder, "... %s\n", aFormat[start:])
		_, _ = fmt.Fprintf(builder, "... %s\n", bFormat[start:])
		_, _ = fmt.Fprintf(builder, "    %s", valueDiff[start:])
	}

	return builder.String()
}
func format(v any) string {
	if isNumeric(v) || isTime(v) {
		return "%v"
	} else {
		return "%#v"
	}
}
func diff(a, b string) string {
	result := new(strings.Builder)
	for x := 0; x < len(a) && x < len(b); x++ {
		if x >= len(a) || x >= len(b) || a[x] != b[x] {
			result.WriteString("^")
		} else {
			result.WriteString(" ")
		}
	}
	return result.String()
}

// deepEquality compares any two values using reflect.DeepEqual.
// https://golang.org/pkg/reflect/#DeepEqual
type deepEquality struct{}

func (deepEquality) assertable(a, b any) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}
func (deepEquality) passes(a, b any) bool {
	return reflect.DeepEqual(a, b)
}

// numericEquality compares numeric values using the built-in equality
// operator (`==`). Values of differing numeric reflect.Kind are each
// converted to the type of the other and are compared with `==` in both
// directions, with one exception: two mixed integers (one signed and one
// unsigned) are always unequal in the case that the unsigned value is
// greater than math.MaxInt64. https://golang.org/pkg/reflect/#Kind
type numericEquality struct{}

func (numericEquality) assertable(a, b any) bool {
	return isNumeric(a) && isNumeric(b)
}
func (numericEquality) passes(a, b any) bool {
	aValue := reflect.ValueOf(a)
	bValue := reflect.ValueOf(b)
	if isUnsignedInteger(a) && isSignedInteger(b) && aValue.Uint() >= math.MaxInt64 {
		return false
	}
	if isSignedInteger(a) && isUnsignedInteger(b) && bValue.Uint() >= math.MaxInt64 {
		return false
	}
	aAsB := aValue.Convert(bValue.Type()).Interface()
	bAsA := bValue.Convert(aValue.Type()).Interface()
	return a == bAsA && b == aAsB
}

// timeEquality compares values both of type time.Time using their Equal method.
// https://golang.org/pkg/time/#Time.Equal
type timeEquality struct{}

func (timeEquality) assertable(a, b any) bool {
	return isTime(a) && isTime(b)
}
func (timeEquality) passes(a, b any) bool {
	return a.(time.Time).Equal(b.(time.Time))
}
func isTime(v any) bool {
	_, ok := v.(time.Time)
	return ok
}

var numericKinds = map[reflect.Kind]struct{}{
	reflect.Int:     {},
	reflect.Int8:    {},
	reflect.Int16:   {},
	reflect.Int32:   {},
	reflect.Int64:   {},
	reflect.Uint:    {},
	reflect.Uint8:   {},
	reflect.Uint16:  {},
	reflect.Uint32:  {},
	reflect.Uint64:  {},
	reflect.Float32: {},
	reflect.Float64: {},
}

func isNumeric(v any) bool {
	of := reflect.TypeOf(v)
	if of == nil {
		return false
	}
	_, found := numericKinds[of.Kind()]
	return found
}

var unsignedIntegerKinds = map[reflect.Kind]struct{}{
	reflect.Uint:    {},
	reflect.Uint8:   {},
	reflect.Uint16:  {},
	reflect.Uint32:  {},
	reflect.Uint64:  {},
	reflect.Uintptr: {},
}

func isUnsignedInteger(v any) bool {
	_, found := unsignedIntegerKinds[reflect.TypeOf(v).Kind()]
	return found
}

var signedIntegerKinds = map[reflect.Kind]struct{}{
	reflect.Int:   {},
	reflect.Int8:  {},
	reflect.Int16: {},
	reflect.Int32: {},
	reflect.Int64: {},
}

func isSignedInteger(v any) bool {
	_, found := signedIntegerKinds[reflect.TypeOf(v).Kind()]
	return found
}

func failure(format string, args ...any) error {
	trace := stack()
	if len(trace) > 0 {
		format += "\nStack: (filtered)\n%s"
		args = append(args, trace)
	}
	return wrap(errAssertionFailure, format, args...)
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
	return "> " + strings.Join(filtered, "\n> ")
}
func readSourceCodeLine(fileLineRaw string) (string, bool) {
	fileLineJoined := strings.Fields(strings.TrimSpace(fileLineRaw))[0]
	fileLine := strings.Split(fileLineJoined, ":")
	sourceCode, _ := os.ReadFile(fileLine[0])
	sourceCodeLines := strings.Split(string(sourceCode), "\n")
	lineNumber, _ := strconv.Atoi(fileLine[1])
	lineNumber--
	lineNumber = min(len(sourceCodeLines)-1, lineNumber)
	return sourceCodeLines[lineNumber], true
}
func wrap(inner error, format string, args ...any) error {
	return fmt.Errorf("%w: "+fmt.Sprintf(format, args...), inner)
}
