package generate

import (
	"regexp"
	"strings"
)

// TODO: implement apostrophe insertion for known contractions (http://www.enchantedlearning.com/grammar/contractions/list.shtml)

var (
	// These regexes allow us to insert a space...
	first  = regexp.MustCompile("([A-Z][a-z]+)") // ...at conventional word boundaries (HelloWorld -> Hello World),
	second = regexp.MustCompile("([A-Z]+)")      // ...between word and uppercase acronym (TestHTTP -> Test HTTP),
	third  = regexp.MustCompile("([^A-Za-z ]+)") // ...and between uppercase acronyms and numerics (HTTP200 -> HTTP 200).
	// Reference: http://stackoverflow.com/a/8837360/605022
)

// toSentence turns identifiers (pascal-cased or underscored) into sentences.
// It replaces underscores with spaces (ie. 'Super_awesome' -> 'super awesome').
// It inserts spaces at casing boundaries (id. 'SuperAwesome' -> 'Super awesome').
// It counts UPPERCASE acronyms as words (ie. 'ILikeHTTP' -> 'I like http').
func toSentence(input string) string {
	input = removeTestAndLongPrefix(input)
	input = breakWordsApart(input)
	input = removeMultipleSpaces(input)
	input = strings.TrimSpace(input)
	return titleCaseSentence(input)
}
func removeTestAndLongPrefix(input string) string {
	if strings.HasPrefix(input, "Test") {
		input = input[len("Test"):]
	} else if strings.HasPrefix(input, "SkipTest") {
		input = input[:len("Test")] + input[len("SkipTest"):]
	} else if strings.HasPrefix(input, "LongTest") {
		input = input[len("LongTest"):]
	} else if strings.HasPrefix(input, "SkipLongTest") {
		input = input[:len("SkipLong")] + input[len("SkipLongTest"):]
	}
	return input
}
func breakWordsApart(input string) string {
	input = strings.Replace(input, "__", ", ", -1)
	input = strings.Replace(input, "_", " ", -1)
	input = first.ReplaceAllString(input, " $1")
	input = second.ReplaceAllString(input, " $1")
	input = third.ReplaceAllString(input, " $1")
	input = strings.Replace(input, " , ", ", ", -1)
	return input
}
func removeMultipleSpaces(input string) string {
	for strings.Contains(input, "  ") {
		input = strings.Replace(input, "  ", " ", -1)
	}
	return input
}
func titleCaseSentence(input string) string {
	words := strings.Split(input, " ")
	first := []string{strings.Title(words[0])}
	for _, word := range words[1:] {
		first = append(first, strings.ToLower(word))
	}
	return strings.Join(first, " ")
}
