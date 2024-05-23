package must

import (
	"errors"
	"fmt"

	"github.com/smarty/gunit/v2/should"
)

var (
	BeChronological = wrapFatal(should.BeChronological)
	// TODO: all the assertions
)

func wrapFatal(original should.Assertion) should.Assertion {
	return func(actual any, expected ...any) error {
		err := original(actual, expected...)
		if err != nil {
			return fmt.Errorf("%w%w", ErrFatalAssertionFailure, err)
		}
		return nil
	}
}

var ErrFatalAssertionFailure = errors.New("")
