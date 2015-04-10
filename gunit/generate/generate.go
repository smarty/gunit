package generate

import (
	"bytes"
	"go/format"
	"text/template"

	"github.com/smartystreets/gunit/gunit/parse"
)

// TestFunction generates a test function based solely on whether the fixture has test cases that are Skipped or not.
func TestFunction(fixture *parse.Fixture) string {
	if fixture.Skipped {
		return executeTemplate(skippedFixtureTemplate, fixture)
	}
	return executeTemplate(testFixtureTemplate, fixture)
}

// FocusedTestFunction generates a test function based solely on whether the fixture has test cases that are Focused or not.
func FocusedTestFunction(fixture *parse.Fixture) string {
	panic("GOPHERS!")
	return ""
}

func executeTemplate(template *template.Template, fixture *parse.Fixture) string {
	writer := &bytes.Buffer{}
	template.Execute(writer, fixture)
	// return writer.String()
	formatted, err := format.Source(writer.Bytes())
	if err != nil {
		panic(err) // TODO: return this error.
	}
	return string(formatted)
}

var skippedFixtureTemplate = template.Must(template.
	New("testFunction").Funcs(map[string]interface{}{"sentence": toSentence}).
	Parse(rawSkippedFixture))

var testFixtureTemplate = template.Must(template.
	New("testFunction").Funcs(map[string]interface{}{"sentence": toSentence}).
	Parse(rawTestFunction))
