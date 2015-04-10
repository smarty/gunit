package generate

const rawTestFile = `//////////////////////////////////////////////////////////////////////////////
// Generated Code ////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

package {{.PackageName}}

import (
	"testing"

	"github.com/smartystreets/gunit"
)
{{range .Fixtures}}
//////////////////////////////////////////////////////////////////////////////

func Test{{.StructName}}(t *testing.T) { {{if .FixtureTeardownName}}
	defer {{.FixtureTeardownName}}()
	{{end}}{{if .FixtureSetupName}}{{.FixtureSetupName}}()

{{end}}
	fixture := gunit.NewFixture(t)
	defer fixture.Finalize()

{{range .TestCases}}
	test{{.Index}} := &{{.StructName}}{Fixture: fixture}
	test{{.Index}}.RunTestCase__(test{{.Index}}.{{.Name}}, "{{.Name | sentence}}")
{{end}}}

func (self *{{.StructName}}) RunTestCase__(test func(), description string) {
	self.T.Log(description){{if .TestTeardownName}}
	defer self.{{.TestTeardownName}}(){{end}}{{if .TestSetupName}}
	self.{{.TestSetupName}}(){{end}}
	test()
}
{{end}}{{/* range .Fixtures */}}
//////////////////////////////////////////////////////////////////////////////
`
