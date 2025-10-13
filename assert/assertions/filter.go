package assertions

import "fmt"

const (
	success                = ""
	needExactValues        = "This assertion requires exactly %d comparison values (you provided %d)."
	needNonEmptyCollection = "This assertion requires at least 1 comparison value (you provided 0)."
)

func need(needed int, expected []any) string {
	if len(expected) != needed {
		return fmt.Sprintf(needExactValues, needed, len(expected))
	}
	return success
}

func atLeast(minimum int, expected []any) string {
	if len(expected) < minimum {
		return needNonEmptyCollection
	}
	return success
}
