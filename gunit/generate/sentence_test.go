package generate

import "testing"

func TestIdentifierToSentence(t *testing.T) {
	for i, test := range sentenceCases {
		if test.SKIP {
			continue
		}
		actual := toSentence(test.input)
		if actual != test.expected {
			t.Errorf("Case #%d:\nInput:    [%s]\nExpected: [%s]\nActual:   [%s]", i, test.input, test.expected, actual)
		} else {
			t.Logf("✔ Case #%d OK '%s' --> '%s'", i, test.input, test.expected)
		}
	}
}

type SentenceTestCase struct {
	input    string
	expected string
	SKIP     bool
}

var sentenceCases = []SentenceTestCase{
	{
		input:    "Hello",
		expected: "Hello",
	},
	{
		input:    "HelloWorld",
		expected: "Hello world",
	},
	{
		input:    "ILikeIceCream",
		expected: "I like ice cream",
	},
	{
		input:    "HTTPIsGreat",
		expected: "HTTP is great",
	},
	{
		input:    "DoYouLikeHTTP",
		expected: "Do you like http",
	},
	{
		input:    "Hello_world",
		expected: "Hello world",
	},
	{
		input:    "WeHaveAnHTTP_API",
		expected: "We have an http api",
	},
	{
		input:    "TheQuickBrownFoxJumpsOverTheLazyDog",
		expected: "The quick brown fox jumps over the lazy dog",
	},
	{
		input:    "The_quick_brown_fox_jumps_over_the_lazy_dog",
		expected: "The quick brown fox jumps over the lazy dog",
	},
}
