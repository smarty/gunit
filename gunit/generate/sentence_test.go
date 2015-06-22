package generate

import "testing"

func TestIdentifierToSentence(t *testing.T) {
	for i, test := range cases {
		if test.SKIP {
			t.Log("?? Skipping case #", i)
		} else {
			if result := toSentence(test.input); result != test.expected {
				t.Errorf("\nExpected: '%s' \nActual:   '%s'", test.expected, result)
			} else {
				t.Logf("'%s' -> '%s'", test.input, result)
			}
		}
	}
}

var cases []Case = []Case{
	Case{"lowercase", "Lowercase", false},
	Case{"Class", "Class", false},
	Case{"MyClass", "My class", false},
	Case{"HTML", "HTML", false},
	Case{"PDFLoader", "PDF loader", false},
	Case{"AString", "A string", false},
	Case{"SimpleXMLParser", "Simple xml parser", false},
	Case{"GL11Version", "GL 11 version", false},
	Case{"99Bottles", "99 bottles", false},
	Case{"May5", "May 5", false},
	Case{"BFG9000", "BFG 9000", false},
	Case{"WhenSomethingHappens__ThenSomethingElseShouldHappen", "When something happens, then something else should happen", false},
	Case{"The_quick_brown_fox_jumps_over_the_lazy_dog", "The quick brown fox jumps over the lazy dog", false},
	Case{"HTTPResponseToSmartyResponseHTTP200ValidJSONBody", "HTTP response to smarty response http 200 valid json body", false},

	// These tests make sure the prefix 'Test' never begins a sentence.
	Case{"TestB", "B", false},
	Case{"TestBB", "BB", false},
	Case{"TestB1", "B 1", false},

	// These tests remove the first occurence of 'Test' in a skipped test case.
	Case{"SkipTestB", "Skip b", false},
	Case{"SkipTestBB", "Skip bb", false},

	// These tests remove the first occurence of 'LongTest' in a long-running test case.
	Case{"LongTestB", "B", false},
	Case{"LongTestBB", "BB", false},
	Case{"LongTestB1", "B 1", false},

	// These tests remove the first occurence of 'Test' in a skipped long-running test case.
	Case{"SkipLongTestB", "Skip long b", false},
	Case{"SkipLongTestBB", "Skip long bb", false},
}

type Case struct {
	input    string
	expected string
	SKIP     bool
}
