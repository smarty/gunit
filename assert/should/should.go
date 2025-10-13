// Package should is simply a rewording of the assertion
// functions in the assertions package.
package should

import "github.com/smarty/gunit/assert/assertions"

var (
	AlmostEqual            = assertions.ShouldAlmostEqual
	BeBetween              = assertions.ShouldBeBetween
	BeBetweenOrEqual       = assertions.ShouldBeBetweenOrEqual
	BeBlank                = assertions.ShouldBeBlank
	BeChronological        = assertions.ShouldBeChronological
	BeEmpty                = assertions.ShouldBeEmpty
	BeFalse                = assertions.ShouldBeFalse
	BeGreaterThan          = assertions.ShouldBeGreaterThan
	BeGreaterThanOrEqualTo = assertions.ShouldBeGreaterThanOrEqualTo
	BeIn                   = assertions.ShouldBeIn
	BeLessThan             = assertions.ShouldBeLessThan
	BeLessThanOrEqualTo    = assertions.ShouldBeLessThanOrEqualTo
	BeNil                  = assertions.ShouldBeNil
	BeTrue                 = assertions.ShouldBeTrue
	BeZeroValue            = assertions.ShouldBeZeroValue
	Contain                = assertions.ShouldContain
	ContainKey             = assertions.ShouldContainKey
	ContainSubstring       = assertions.ShouldContainSubstring
	EndWith                = assertions.ShouldEndWith
	Equal                  = assertions.ShouldEqual
	HappenAfter            = assertions.ShouldHappenAfter
	HappenBefore           = assertions.ShouldHappenBefore
	HappenBetween          = assertions.ShouldHappenBetween
	HappenOnOrAfter        = assertions.ShouldHappenOnOrAfter
	HappenOnOrBefore       = assertions.ShouldHappenOnOrBefore
	HappenOnOrBetween      = assertions.ShouldHappenOnOrBetween
	HappenWithin           = assertions.ShouldHappenWithin
	HaveLength             = assertions.ShouldHaveLength
	HaveSameTypeAs         = assertions.ShouldHaveSameTypeAs
	NotAlmostEqual         = assertions.ShouldNotAlmostEqual
	NotBeBetween           = assertions.ShouldNotBeBetween
	NotBeBetweenOrEqual    = assertions.ShouldNotBeBetweenOrEqual
	NotBeBlank             = assertions.ShouldNotBeBlank
	NotBeChronological     = assertions.ShouldNotBeChronological
	NotBeEmpty             = assertions.ShouldNotBeEmpty
	NotBeIn                = assertions.ShouldNotBeIn
	NotBeNil               = assertions.ShouldNotBeNil
	NotBeZeroValue         = assertions.ShouldNotBeZeroValue
	NotContain             = assertions.ShouldNotContain
	NotContainKey          = assertions.ShouldNotContainKey
	NotContainSubstring    = assertions.ShouldNotContainSubstring
	NotEndWith             = assertions.ShouldNotEndWith
	NotEqual               = assertions.ShouldNotEqual
	NotHappenOnOrBetween   = assertions.ShouldNotHappenOnOrBetween
	NotHappenWithin        = assertions.ShouldNotHappenWithin
	NotHaveSameTypeAs      = assertions.ShouldNotHaveSameTypeAs
	NotPanic               = assertions.ShouldNotPanic
	NotPanicWith           = assertions.ShouldNotPanicWith
	NotStartWith           = assertions.ShouldNotStartWith
	Panic                  = assertions.ShouldPanic
	PanicWith              = assertions.ShouldPanicWith
	StartWith              = assertions.ShouldStartWith
	Wrap                   = assertions.ShouldWrap
)
