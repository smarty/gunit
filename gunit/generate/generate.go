package generate

import (
	"bytes"
	"fmt"
	"go/format"
	"text/template"

	"github.com/smartystreets/gunit/gunit/parse"
)

func TestFile(packageName string, fixtures []*parse.Fixture) string {
	buffer := bytes.NewBufferString(fmt.Sprintf(header, packageName))
	buffer.WriteString("\n///////////////////////////////////////////////////////////////////////////////\n")
	for _, fixture := range fixtures {
		buffer.WriteString(TestFunction(fixture))
		buffer.WriteString("\n\n///////////////////////////////////////////////////////////////////////////////\n\n")
	}
	buffer.WriteString(footer)
	fmt.Println(buffer.String())
	formatted, err := format.Source(buffer.Bytes())
	if err != nil {
		panic(err) // TODO: return this error
	}
	return string(formatted)
}

// TestFunction generates a test function based solely on whether the fixture has test cases that are Skipped or not.
func TestFunction(fixture *parse.Fixture) string {
	if fixture.Skipped {
		return executeTemplate(skippedFixtureTemplate, fixture)
	}
	return executeTemplate(testFixtureTemplate, fixture)
}

func executeTemplate(template *template.Template, fixture *parse.Fixture) string {
	writer := &bytes.Buffer{}
	template.Execute(writer, fixture)
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
