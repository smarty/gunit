package assertions

import "github.com/smarty/gunit/assert"

// so is like So, except that it only returns the string message, which is blank if the
// assertion passed. Used to facilitate testing.
func so(actual any, assertion assert.Func, expected ...any) string {
	return assertion(actual, expected...)
}
