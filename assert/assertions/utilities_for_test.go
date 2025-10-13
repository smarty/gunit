package assertions

import (
	"strings"
	"testing"

	"github.com/smarty/gunit/assert/assertions/internal/unit"
)

/**************************************************************************/

func TestAssertionsFixture(t *testing.T) {
	unit.Run(new(AssertionsFixture), t)
}

type AssertionsFixture struct {
	*unit.Fixture
}

func (this *AssertionsFixture) pass(result string) {
	this.Assert(result == success, result)
}

func (this *AssertionsFixture) fail(actual string, expected string) {
	actual = format(actual)
	expected = format(expected)

	if expected == "[no-check]" && actual == "" {
		this.Errorf("Expected fail, but assertion passed.")
	} else if expected == "[no-check]" {
		return
	}
	if actual != expected {
		if actual == "" {
			actual = "(empty)"
		}
		this.Errorf("Expected: %s\nActual:   %s\n", expected, actual)
	}
}
func format(message string) string {
	message = strings.ReplaceAll(message, "\n", " ")
	for strings.Contains(message, "  ") {
		message = strings.ReplaceAll(message, "  ", " ")
	}
	message = strings.ReplaceAll(message, "\x1b[32m", "")
	message = strings.ReplaceAll(message, "\x1b[31m", "")
	message = strings.ReplaceAll(message, "\x1b[0m", "")
	return message
}

/**************************************************************************/

type Thing1 struct {
	a string
}
type Thing2 struct {
	a string
}

type ThingInterface interface {
	Hi()
}

type ThingImplementation struct{}

func (self *ThingImplementation) Hi() {}
