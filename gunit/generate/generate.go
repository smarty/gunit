package generate

import (
	"bytes"
	"fmt"
	"go/format"
	"text/template"

	"github.com/smartystreets/gunit/gunit/parse"
)

// TODO: need to return an error as well (if formatting fails, the source code won't compile and we shouldn't write the contents to a *_test.go file...).
func TestFile(packageName string, parsed []*parse.Fixture) string {
	data := PackageFixtures{PackageName: packageName, Fixtures: parsed}
	writer := &bytes.Buffer{}
	compiled.Execute(writer, data)
	// return writer.String()
	formatted, err := format.Source(writer.Bytes())
	if err != nil {
		panic(err)
	}
	fmt.Println(string(formatted))
	return string(formatted)
}

type PackageFixtures struct {
	PackageName string
	Fixtures    []*parse.Fixture
}

var compiled = template.Must(template.
	New("testfile").Funcs(map[string]interface{}{"sentence": toSentence}).
	Parse(rawTestFile))
