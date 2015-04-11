package generate

import (
	"bytes"
	"fmt"
	"go/format"
	"text/template"

	"github.com/smartystreets/gunit/gunit/parse"
)

// TestFunction generates complete source code for a _test.go file from the provided fixtures.
func TestFile(packageName string, fixtures []*parse.Fixture) (string, error) {
	buffer := bytes.NewBufferString(fmt.Sprintf(header, packageName))
	buffer.WriteString("\n///////////////////////////////////////////////////////////////////////////////\n\n")
	for _, fixture := range fixtures {
		function, err := TestFunction(fixture)
		if err != nil {
			return "", err
		}
		buffer.WriteString(function)
		buffer.WriteString("\n\n///////////////////////////////////////////////////////////////////////////////\n\n")
	}
	buffer.WriteString(footer)
	formatted, err := format.Source(buffer.Bytes())
	if err != nil {
		return "", err
	}
	return string(formatted), nil
}

// TestFunction generates a test function based solely on whether the fixture has test cases that are Skipped or not.
func TestFunction(fixture *parse.Fixture) (string, error) {
	if fixture.Skipped {
		return executeTemplate(skippedFixtureTemplate, fixture)
	}
	return executeTemplate(testFixtureTemplate, fixture)
}

func executeTemplate(template *template.Template, fixture *parse.Fixture) (string, error) {
	writer := &bytes.Buffer{}
	template.Execute(writer, fixture)
	formatted, err := format.Source(writer.Bytes())
	if err != nil {
		return "", err
	}
	return string(formatted), nil
}

var skippedFixtureTemplate = template.Must(template.
	New("testFunction").Funcs(map[string]interface{}{"sentence": toSentence}).
	Parse(rawSkippedFixture))

var testFixtureTemplate = template.Must(template.
	New("testFunction").Funcs(map[string]interface{}{"sentence": toSentence}).
	Parse(rawTestFunction))
