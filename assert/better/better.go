// Package better contains the same listing as the should package,
// but each assertion is wrapped in behavior that decorates failure
// output with a 'fatal' prefix/indicator so that the So method can
// invoke testing.T.Fatal(...) to immediately end a test case.
package better

import (
	"github.com/smarty/gunit/assert"
	"github.com/smarty/gunit/assert/should"
)

func fatal(so assert.Func) assert.Func {
	return func(actual any, expected ...any) string {
		result := so(actual, expected...)
		if result == "" {
			return ""
		}
		return "<<<FATAL>>>\n" + result
	}
}

var (
	AlmostEqual            = fatal(should.AlmostEqual)
	BeBetween              = fatal(should.BeBetween)
	BeBetweenOrEqual       = fatal(should.BeBetweenOrEqual)
	BeBlank                = fatal(should.BeBlank)
	BeChronological        = fatal(should.BeChronological)
	BeEmpty                = fatal(should.BeEmpty)
	BeFalse                = fatal(should.BeFalse)
	BeGreaterThan          = fatal(should.BeGreaterThan)
	BeGreaterThanOrEqualTo = fatal(should.BeGreaterThanOrEqualTo)
	BeIn                   = fatal(should.BeIn)
	BeLessThan             = fatal(should.BeLessThan)
	BeLessThanOrEqualTo    = fatal(should.BeLessThanOrEqualTo)
	BeNil                  = fatal(should.BeNil)
	BeTrue                 = fatal(should.BeTrue)
	BeZeroValue            = fatal(should.BeZeroValue)
	Contain                = fatal(should.Contain)
	ContainKey             = fatal(should.ContainKey)
	ContainSubstring       = fatal(should.ContainSubstring)
	EndWith                = fatal(should.EndWith)
	Equal                  = fatal(should.Equal)
	HappenAfter            = fatal(should.HappenAfter)
	HappenBefore           = fatal(should.HappenBefore)
	HappenBetween          = fatal(should.HappenBetween)
	HappenOnOrAfter        = fatal(should.HappenOnOrAfter)
	HappenOnOrBefore       = fatal(should.HappenOnOrBefore)
	HappenOnOrBetween      = fatal(should.HappenOnOrBetween)
	HappenWithin           = fatal(should.HappenWithin)
	HaveLength             = fatal(should.HaveLength)
	HaveSameTypeAs         = fatal(should.HaveSameTypeAs)
	NotAlmostEqual         = fatal(should.NotAlmostEqual)
	NotBeBetween           = fatal(should.NotBeBetween)
	NotBeBetweenOrEqual    = fatal(should.NotBeBetweenOrEqual)
	NotBeBlank             = fatal(should.NotBeBlank)
	NotBeChronological     = fatal(should.NotBeChronological)
	NotBeEmpty             = fatal(should.NotBeEmpty)
	NotBeIn                = fatal(should.NotBeIn)
	NotBeNil               = fatal(should.NotBeNil)
	NotBeZeroValue         = fatal(should.NotBeZeroValue)
	NotContain             = fatal(should.NotContain)
	NotContainKey          = fatal(should.NotContainKey)
	NotContainSubstring    = fatal(should.NotContainSubstring)
	NotEndWith             = fatal(should.NotEndWith)
	NotEqual               = fatal(should.NotEqual)
	NotHappenOnOrBetween   = fatal(should.NotHappenOnOrBetween)
	NotHappenWithin        = fatal(should.NotHappenWithin)
	NotHaveSameTypeAs      = fatal(should.NotHaveSameTypeAs)
	NotPanic               = fatal(should.NotPanic)
	NotPanicWith           = fatal(should.NotPanicWith)
	NotStartWith           = fatal(should.NotStartWith)
	Panic                  = fatal(should.Panic)
	PanicWith              = fatal(should.PanicWith)
	StartWith              = fatal(should.StartWith)
	Wrap                   = fatal(should.Wrap)
)
