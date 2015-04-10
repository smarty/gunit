package generate

import (
	"bytes"
	"text/template"

	"github.com/smartystreets/gunit/gunit/parse"
)

func TestFile(packageName string, parsed []*parse.Fixture) string {
	data := PackageFixtures{PackageName: packageName, Fixtures: fixtures}
	writer := &bytes.Buffer{}
	compiled.Execute(writer, data)
	return writer.String()
	// formatted, err := format.Source(writer.Bytes())
	// if err != nil {
	// 	panic(err)
	// }
	// return string(formatted)
}

type PackageFixtures struct {
	PackageName string
	Fixtures    []*parse.Fixture
}

var compiled = template.Must(
	template.
		New("testfile").Funcs(map[string]interface{}{"sentence": toSentence}).
		Parse(rawTestFile))
