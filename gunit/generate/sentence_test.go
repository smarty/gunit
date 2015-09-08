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
	Case{"lowercase", "lowercase", false},
	Case{"Class", "class", false},
	Case{"MyClass", "my_class", false},
	Case{"HTML", "html", false},
	Case{"PDFLoader", "pdf_loader", false},
	Case{"AString", "a_string", false},
	Case{"SimpleXMLParser", "simple_xml_parser", false},
	Case{"GL11Version", "gl_11_version", false},
	Case{"99Bottles", "99_bottles", false},
	Case{"May5", "may_5", false},
	Case{"BFG9000", "bfg_9000", false},
	Case{"The_quick_brown_fox_jumps_over_the_lazy_dog", "the_quick_brown_fox_jumps_over_the_lazy_dog", false},
	Case{"HTTPResponseToSmartyResponseHTTP200ValidJSONBody", "http_response_to_smarty_response_http_200_valid_json_body", false},

	// These tests make sure the prefix 'Test' never begins a sentence.
	Case{"TestB", "b", false},
	Case{"TestBB", "bb", false},
	Case{"TestB1", "b_1", false},

	// These tests remove the first occurence of 'Test' in a skipped test case.
	Case{"SkipTestB", "skip_b", false},
	Case{"SkipTestBB", "skip_bb", false},

	// These tests remove the first occurence of 'LongTest' in a long-running test case.
	Case{"LongTestB", "b", false},
	Case{"LongTestBB", "bb", false},
	Case{"LongTestB1", "b_1", false},

	// These tests remove the first occurence of 'Test' in a skipped long-running test case.
	Case{"SkipLongTestB", "skip_long_b", false},
	Case{"SkipLongTestBB", "skip_long_bb", false},
}

type Case struct {
	input    string
	expected string
	SKIP     bool
}
