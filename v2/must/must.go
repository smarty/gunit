// Package must provides the same assertions as the should package, but the
// error returned in failure conditions results in a call to *testing.T.Fatal(),
// halting the currently running test.
package must

import (
	"fmt"

	"github.com/smarty/gunit/v2/should"
)

var (
	BeChronological        = wrapFatal(should.BeChronological)
	BeEmpty                = wrapFatal(should.BeEmpty)
	BeFalse                = wrapFatal(should.BeFalse)
	BeGreaterThan          = wrapFatal(should.BeGreaterThan)
	BeGreaterThanOrEqualTo = wrapFatal(should.BeGreaterThanOrEqualTo)
	BeIn                   = wrapFatal(should.BeIn)
	BeLessThan             = wrapFatal(should.BeLessThan)
	BeLessThanOrEqualTo    = wrapFatal(should.BeLessThanOrEqualTo)
	BeNil                  = wrapFatal(should.BeNil)
	BeTrue                 = wrapFatal(should.BeTrue)
	Contain                = wrapFatal(should.Contain)
	EndWith                = wrapFatal(should.EndWith)
	Equal                  = wrapFatal(should.Equal)
	HappenAfter            = wrapFatal(should.HappenAfter)
	HappenBefore           = wrapFatal(should.HappenBefore)
	HappenOn               = wrapFatal(should.HappenOn)
	HappenWithin           = wrapFatal(should.HappenWithin)
	HaveLength             = wrapFatal(should.HaveLength)
	Panic                  = wrapFatal(should.Panic)
	StartWith              = wrapFatal(should.StartWith)
	WrapError              = wrapFatal(should.WrapError)
)

// NOT constrains all negated assertions to their own 'namespace'.
var NOT = struct {
	BeChronological        should.Assertion
	BeEmpty                should.Assertion
	BeGreaterThan          should.Assertion
	BeGreaterThanOrEqualTo should.Assertion
	BeIn                   should.Assertion
	BeLessThan             should.Assertion
	BeLessThanOrEqualTo    should.Assertion
	BeNil                  should.Assertion
	Contain                should.Assertion
	Equal                  should.Assertion
	HappenOn               should.Assertion
	Panic                  should.Assertion
}{
	BeChronological:        wrapFatal(should.NOT.BeChronological),
	BeEmpty:                wrapFatal(should.NOT.BeEmpty),
	BeGreaterThan:          wrapFatal(should.NOT.BeGreaterThan),
	BeGreaterThanOrEqualTo: wrapFatal(should.NOT.BeGreaterThanOrEqualTo),
	BeIn:                   wrapFatal(should.NOT.BeIn),
	BeLessThan:             wrapFatal(should.NOT.BeLessThan),
	BeLessThanOrEqualTo:    wrapFatal(should.NOT.BeLessThanOrEqualTo),
	BeNil:                  wrapFatal(should.NOT.BeNil),
	Contain:                wrapFatal(should.NOT.Contain),
	Equal:                  wrapFatal(should.NOT.Equal),
	HappenOn:               wrapFatal(should.NOT.HappenOn),
	Panic:                  wrapFatal(should.NOT.Panic),
}

func wrapFatal(original should.Assertion) should.Assertion {
	return func(actual any, expected ...any) error {
		err := original(actual, expected...)
		if err != nil {
			return fmt.Errorf("%w %w", should.ErrFatalAssertionFailure, err)
		}
		return nil
	}
}
