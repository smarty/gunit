package generate

import (
	"strings"
	"unicode"

	"bytes"
)

// toSentence turns identifiers (pascal-cased or underscored) into sentences.
// It replaces underscores with spaces (ie. 'Super_awesome' -> 'super awesome').
// It inserts spaces at casing boundaries (id. 'SuperAwesome' -> 'Super awesome').
// It counts UPPERCASE acronyms as words (ie. 'ILikeHTTP' -> 'I like http').
func toSentence(identifier string) string {
	var (
		sentence = &bytes.Buffer{}
		current  = NewWordBuffer()
	)

	for _, c := range identifier {
		if !current.Accept(c) {
			sentence.WriteString(current.String())
			current.Next()
		}
	}
	sentence.WriteString(current.String())
	return strings.TrimSpace(sentence.String())
}

//////////////////////////////////////////////////////////////////////////////

type word struct {
	firstWord bool
	current   []rune
	next      []rune
}

func NewWordBuffer() *word {
	return &word{firstWord: true}
}

// Accept decides what to do with the incoming (next) letter. It returns false if
// the word is complete and a new word should be started, by calling Next().
func (self *word) Accept(in rune) bool {
	if NewFirstCharacterSpecification(self.current).IsSatisfiedBy(in) {
		self.current = append(self.current, in)

	} else if NewSecondCharacterSpecification(self.current).IsSatisfiedBy(in) {
		self.current = append(self.current, in)

	} else if isUnderscoreWordBoundary(in) {
		return false

	} else if NewNextCharacterInWordSpecification(self.current).IsSatisfiedBy(in) {
		self.current = append(self.current, in)

	} else if NewWordBoundarySpecification(self.current).IsSatisfiedBy(in) {
		self.next = append(self.next, in)
		return false

	} else if NewAcronymEndedSpecification(self.current).IsSatisfiedBy(in) {
		self.chopAcronym(in)
		return false
	}

	return true
}

func (self *word) chopAcronym(in rune) {
	last := len(self.current) - 1
	self.next = append(self.next, self.current[last], in)
	self.current = self.current[:last]
}

// Spits out the currently assembed word.
func (self *word) String() string {
	if self.firstWord {
		return strings.Title(string(self.current))
	}
	return " " + strings.ToLower(string(self.current))
}

// Replaces the current word with any characters gathered for the next word.
func (self *word) Next() {
	self.firstWord = false
	self.current, self.next = self.next, nil
}

//////////////////////////////////////////////////////////////////////////////

func upper(r rune) bool   { return unicode.IsUpper(r) }
func lower(r rune) bool   { return unicode.IsLower(r) }
func toLower(r rune) rune { return unicode.ToLower(r) }

//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

type FirstCharacterSpecification struct{ word []rune }

func NewFirstCharacterSpecification(word []rune) *FirstCharacterSpecification {
	return &FirstCharacterSpecification{word: word}
}
func (self *FirstCharacterSpecification) IsSatisfiedBy(character rune) bool {
	return len(self.word) == 0
}

//////////////////////////////////////////////////////////////////////////////

type SecondCharacterSpecification struct{ word []rune }

func NewSecondCharacterSpecification(word []rune) *SecondCharacterSpecification {
	return &SecondCharacterSpecification{word: word}
}

func (self *SecondCharacterSpecification) IsSatisfiedBy(character rune) bool {
	if len(self.word) != 1 {
		return false
	}
	if upper(character) {
		return false
	}
	return true
}

//////////////////////////////////////////////////////////////////////////////

type NextCharacterInWordSpecification struct {
	word              []rune
	previousCharacter rune
}

func NewNextCharacterInWordSpecification(word []rune) *NextCharacterInWordSpecification {
	last := len(word) - 1
	return &NextCharacterInWordSpecification{
		word:              word,
		previousCharacter: word[last],
	}
}
func (self *NextCharacterInWordSpecification) IsSatisfiedBy(character rune) bool {
	return upper(character) == upper(self.previousCharacter)
}

//////////////////////////////////////////////////////////////////////////////

func isUnderscoreWordBoundary(c rune) bool {
	return c == '_'
}

//////////////////////////////////////////////////////////////////////////////

type WordBoundarySpecification struct{ word []rune }

func NewWordBoundarySpecification(word []rune) *WordBoundarySpecification {
	return &WordBoundarySpecification{word: word}
}

func (self *WordBoundarySpecification) IsSatisfiedBy(character rune) bool {
	last := len(self.word) - 1
	return upper(character) && lower(self.word[last])
}

//////////////////////////////////////////////////////////////////////////////

type AcronymEndedSpecification struct{ word []rune }

func NewAcronymEndedSpecification(word []rune) *AcronymEndedSpecification {
	return &AcronymEndedSpecification{word: word}
}

func (self *AcronymEndedSpecification) IsSatisfiedBy(character rune) bool {
	last := len(self.word) - 1
	if !upper(self.word[last]) {
		return false
	}
	penultimate := len(self.word) - 2
	if !upper(self.word[penultimate]) {
		return false
	}
	if upper(character) {
		return false
	}
	return true
}

//////////////////////////////////////////////////////////////////////////////
