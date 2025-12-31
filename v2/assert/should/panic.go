package should

import "errors"

// Panic invokes the func() provided as actual and recovers from any
// panic. It returns an error if actual() does not result in a panic.
func Panic(actual any, expected ...any) (err error) {
	err = NOT.Panic(actual, expected...)
	if errors.Is(err, ErrAssertionFailure) {
		return nil
	}

	if err != nil {
		return err
	}

	return failure("provided func did not panic as expected")
}

// Panic (negated!) expects the func() provided as actual to run without panicking.
func (negated) Panic(actual any, expected ...any) (err error) {
	err = validateExpected(0, expected)
	if err != nil {
		return err
	}

	err = validateType(actual, func() {})
	if err != nil {
		return err
	}

	panicked := true
	defer func() {
		r := recover()
		if panicked {
			err = failure(""+
				"provided func should not have"+
				"panicked but it did with: %s", r,
			)
		}
	}()

	actual.(func())()
	panicked = false
	return nil
}
