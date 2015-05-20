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
	Case{"TestB", "Test b", false},
	Case{"TestBB", "Test bb", false},
	Case{"TestB1", "Test b 1", false},
	Case{"TestHTTPResponseToSmartyResponseHTTP200ValidJSONBody", "Test http response to smarty response http 200 valid json body", false},
	Case{"The_quick_brown_fox_jumps_over_the_lazy_dog", "The quick brown fox jumps over the lazy dog", false},
}

type Case struct {
	input    string
	expected string
	SKIP     bool
}
