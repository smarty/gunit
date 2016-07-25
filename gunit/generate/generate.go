package generate

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"text/template"

	"github.com/smartystreets/gunit/gunit/parse"
)

// TestFile generates complete source code for a _test.go file from the provided fixtures.
func TestFile(packageName string, fixtures []*parse.Fixture, checksum string) ([]byte, error) {
	buffer := bytes.NewBufferString(fmt.Sprintf(header, packageName))
	buffer.WriteString("\n///////////////////////////////////////////////////////////////////////////////\n\n")
	for _, fixture := range fixtures {
		function, err := TestCases(fixture)
		if err != nil {
			return nil, err
		}
		buffer.Write(function)
		buffer.WriteString("\n\n///////////////////////////////////////////////////////////////////////////////\n\n")
	}
	buffer.WriteString(fmt.Sprintf(footer, checksum))
	return format.Source(buffer.Bytes())
}

func TestCases(fixture *parse.Fixture) ([]byte, error) {
	writer := &bytes.Buffer{}
	err := testFixtureTemplate.Execute(writer, fixture)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return format.Source(writer.Bytes())
}

var testFixtureTemplate = template.Must(template.New("testFunction").Parse(rawTestFunction))
