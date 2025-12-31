package should

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strings"
	"time"

	diffmatchpatch2 "github.com/smarty/gunit/v2/assert/should/internal/go-diff/diffmatchpatch"
	"github.com/smarty/gunit/v2/assert/should/internal/go-render/render"
)

// Equal verifies that the actual value is equal to the expected value.
// It uses reflect.DeepEqual in most cases, but also compares numerics
// regardless of specific type and compares time.Time values using the
// time.Equal method.
func Equal(actual any, EXPECTED ...any) error {
	err := validateExpected(1, EXPECTED)
	if err != nil {
		return err
	}

	expected := EXPECTED[0]

	for _, spec := range equalitySpecs {
		if !spec.assertable(actual, expected) {
			continue
		}
		if spec.passes(actual, expected) {
			return nil
		}
		break
	}
	return failure(report(actual, expected))
}

// Equal negated!
func (negated) Equal(actual any, expected ...any) error {
	err := Equal(actual, expected...)
	if errors.Is(err, ErrAssertionFailure) {
		return nil
	}

	if err != nil {
		return err
	}

	return failure("\n"+
		"  expected:     %#v\n"+
		"  to not equal: %#v\n"+
		"  (but it did)",
		expected[0],
		actual,
	)
}

var equalitySpecs = []specification{
	numericEquality{},
	timeEquality{},
	deepEquality{},
}

func report(a, b any) string {
	builder := new(strings.Builder)
	builder.WriteString(simpleDiff(a, b))
	builder.WriteString(prettyDiff(a, b))
	builder.WriteString("\n")
	return builder.String()
}

func simpleDiff(a, b any) string {
	aType := fmt.Sprintf("(%v)", reflect.TypeOf(a))
	bType := fmt.Sprintf("(%v)", reflect.TypeOf(b))
	longestType := int(math.Max(float64(len(aType)), float64(len(bType))))
	aType += strings.Repeat(" ", longestType-len(aType))
	bType += strings.Repeat(" ", longestType-len(bType))
	aFormat := fmt.Sprintf(format(a), a)
	bFormat := fmt.Sprintf(format(b), b)

	builder := new(strings.Builder)
	typeDiff := diff(bType, aType)
	valueDiff := diff(bFormat, aFormat)

	_, _ = fmt.Fprintf(builder, "\n")
	_, _ = fmt.Fprintf(builder, "Expected: %s %s\n", bType, bFormat)
	_, _ = fmt.Fprintf(builder, "Actual:   %s %s\n", aType, aFormat)
	_, _ = fmt.Fprintf(builder, "          %s %s", typeDiff, valueDiff)

	if firstDiffIndex := strings.Index(valueDiff, "^"); firstDiffIndex > 40 {
		start := firstDiffIndex - 20
		_, _ = fmt.Fprintf(builder, "\nInitial discrepancy at index %d:\n", firstDiffIndex)
		_, _ = fmt.Fprintf(builder, "... %s\n", bFormat[start:])
		_, _ = fmt.Fprintf(builder, "... %s\n", aFormat[start:])
		_, _ = fmt.Fprintf(builder, "    %s", valueDiff[start:])
	}
	return builder.String()
}

func prettyDiff(actual, expected any) string {
	diff := diffmatchpatch2.New()
	diffs := diff.DiffMain(render.Render(expected), render.Render(actual), false)
	if prettyDiffIsLikelyToBeHelpful(diffs) {
		return fmt.Sprintf("\nDiff: '%s'", diff.DiffPrettyText(diffs))
	}
	return ""
}

// prettyDiffIsLikelyToBeHelpful returns true if the diff listing contains
// more 'equal' segments than 'deleted'/'inserted' segments.
func prettyDiffIsLikelyToBeHelpful(diffs []diffmatchpatch2.Diff) bool {
	equal, deleted, inserted := measureDiffTypeLengths(diffs)
	return equal > deleted && equal > inserted
}

func measureDiffTypeLengths(diffs []diffmatchpatch2.Diff) (equal, deleted, inserted int) {
	for _, segment := range diffs {
		switch segment.Type {
		case diffmatchpatch2.DiffEqual:
			equal += len(segment.Text)
		case diffmatchpatch2.DiffDelete:
			deleted += len(segment.Text)
		case diffmatchpatch2.DiffInsert:
			inserted += len(segment.Text)
		}
	}
	return equal, deleted, inserted
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
