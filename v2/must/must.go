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

// NOT (a singleton) constrains all negated assertions to their own namespace.
var NOT negated

type negated struct{}

func (negated) BeChronological(actual any, expected ...any) error {
	return wrapFatal(should.NOT.BeChronological)(actual, expected...)
}
func (negated) BeEmpty(actual any, expected ...any) error {
	return wrapFatal(should.NOT.BeEmpty)(actual, expected...)
}
func (negated) BeGreaterThan(actual any, expected ...any) error {
	return wrapFatal(should.NOT.BeGreaterThan)(actual, expected...)
}
func (negated) BeGreaterThanOrEqualTo(actual any, expected ...any) error {
	return wrapFatal(should.NOT.BeGreaterThanOrEqualTo)(actual, expected...)
}
func (negated) BeIn(actual any, expected ...any) error {
	return wrapFatal(should.NOT.BeIn)(actual, expected...)
}
func (negated) BeLessThan(actual any, expected ...any) error {
	return wrapFatal(should.NOT.BeLessThan)(actual, expected...)
}
func (negated) BeLessThanOrEqualTo(actual any, expected ...any) error {
	return wrapFatal(should.NOT.BeLessThanOrEqualTo)(actual, expected...)
}
func (negated) BeNil(actual any, expected ...any) error {
	return wrapFatal(should.NOT.BeNil)(actual, expected...)
}
func (negated) Contain(actual any, expected ...any) error {
	return wrapFatal(should.NOT.Contain)(actual, expected...)
}
func (negated) Equal(actual any, expected ...any) error {
	return wrapFatal(should.NOT.Equal)(actual, expected...)
}
func (negated) HappenOn(actual any, expected ...any) error {
	return wrapFatal(should.NOT.HappenOn)(actual, expected...)
}
func (negated) Panic(actual any, expected ...any) error {
	return wrapFatal(should.NOT.Panic)(actual, expected...)
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
