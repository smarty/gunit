package generate

import (
	"bytes"
	"text/template"

	"github.com/smartystreets/gunit/gunit/parse"
)

func TestFile(packageName string, parsed []*parse.Fixture) string {
	fixtures := []*Fixture{}
	for _, p := range parsed {
		fixture := &Fixture{
			Skipped:             p.Skipped,
			Focused:             p.Focused,
			StructName:          p.StructName,
			FixtureSetupName:    p.FixtureSetupName,
			FixtureTeardownName: p.FixtureTeardownName,
			TestSetupName:       p.TestSetupName,
			TestTeardownName:    p.TestTeardownName,
		}
		// for _, focusedTestCase := range p.FocusedTestCaseNames {
		// 	fixture.FocusedTestCases = true
		// 	fixture.TestCases = append(fixture.TestCases, TestCase{
		// 		StructName: fixture.StructName,
		// 		Name:       focusedTestCase,
		// 		Focused:    true,
		// 		Index:      len(fixture.TestCases),
		// 	})
		// }
		// for _, skippedTestcase := range p.SkippedTestCaseNames {
		// 	fixture.TestCases = append(fixture.TestCases, TestCase{
		// 		StructName: fixture.StructName,
		// 		Name:       skippedTestcase,
		// 		Skipped:    true,
		// 		Index:      len(fixture.TestCases),
		// 	})
		// }
		for _, testCaseName := range p.TestCaseNames {
			fixture.TestCases = append(fixture.TestCases, TestCase{
				StructName: fixture.StructName,
				Name:       testCaseName,
				Index:      len(fixture.TestCases),
			})
		}
		fixtures = append(fixtures, fixture)
	}
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
	Fixtures    []*Fixture
}

type Fixture struct {
	Skipped    bool
	Focused    bool
	StructName string

	FixtureSetupName    string
	FixtureTeardownName string

	TestSetupName    string
	TestTeardownName string

	TestCases        []TestCase
	FocusedTestCases bool
}
type TestCase struct {
	Index      int
	StructName string
	Name       string
	Skipped    bool
	Focused    bool
}

var compiled = template.Must(
	template.
		New("testfile").Funcs(map[string]interface{}{"sentence": toSentence}).
		Parse(rawTestFile))
