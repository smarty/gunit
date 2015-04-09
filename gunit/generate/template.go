package generate

const rawTestFile = `//////////////////////////////////////////////////////////////////////////////
// Generated Code ////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

package {{.PackageName}}

import (
	"testing"

	"github.com/smartystreets/gunit"
)

//////////////////////////////////////////////////////////////////////////////

{{range .Fixtures}}func Test{{.StructName}}(t *testing.T) {
	fixture := gunit.NewTestCase(t)
	defer fixture.Finalize()

	{{range .TestCases}}test{{.Index}} := &{{.StructName}}{TestCase: fixture}
	test{{.Index}}.RunTestCase(test{{.Index}}.{{.Name}}, "{{.Name | sentence}}"){{end}}
}

func (self *{{.StructName}}) RunTestCase(test func(), description string) {
	self.T.Log(description)
	test()
}

//////////////////////////////////////////////////////////////////////////////
{{end}}{{/* range .Fixtures */}}`

const blahblah = `//////////////////////////////////////////////////////////////////////////////
// Generated Code ////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

package {{.PackageName}}

import (
	"testing"

	"github.com/smartystreets/gunit"
)

//////////////////////////////////////////////////////////////////////////////

{{range .Fixtures}}

func Test{{.StructName}}(t *testing.T) {
	{{if .Skipped}}
	t.Skip("Fixture marked as skipped:", {{.StructName}})
	{{else}}{{/* TODO: long-running */}}{{/* TODO: no test cases */}}

	fixture := gunit.NewTestCase(t)
	defer fixture.Finalize()

	{{if .FixtureTeardownName}}defer {{.FixtureTeardownName}}(){{end}}
	{{if .FixtureSetupName}}{{.FixtureSetupName}}(){{end}}

	{{range $index, $TestCaseName := .TestCaseNames}}
	test{{$index}} := &{ {{.StructName}} }{TestCase: fixture}
	test{{$index}}.RunTestCase(
		test{{$index}}.{{$TestCaseName}},
		"{{$TestCaseName | sentence}}")
	{{end}}{{/* range .TestCaseNames */}}

	{{end}}{{/* if .FixtureSkipped */}}
}

func (self *{{.StructName}}) RunTestCase(test func(), description string) {
	self.T.Log(description)
	{{if .TestCaseTeardownName}}defer self.{{.TestCaseTeardownName}}(){{end}}
	{{if .TestCaseSetupName}}self.{{.TestCaseSetupName}}(){{end}}
	test()
}

//////////////////////////////////////////////////////////////////////////////
{{end}}{{/* range .Fixtures */}}
`
