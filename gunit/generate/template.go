package generate

import "strings"

const GeneratedFilename = "generated_by_gunit_test.go"

//////////////////////////////////////////////////////////////////////////////

var rawSkippedFixture = strings.TrimSpace(`
func Test{{.StructName}}(t *testing.T) { 
	fixture := gunit.NewFixture(t, os.Stdout, testing.Verbose())
	defer fixture.Finalize()
	{{range .TestCases}}
	fixture.Skip("Skipping test case: '{{.Name | sentence}}'") {{end}}
}
`)

//////////////////////////////////////////////////////////////////////////////

var rawTestFunction = strings.TrimSpace(`
func Test{{.StructName}}(t *testing.T) { {{if .FixtureTeardownName}}
	defer {{.FixtureTeardownName}}()
	{{end}}{{if .FixtureSetupName}}{{.FixtureSetupName}}()

{{end}}
	fixture := gunit.NewFixture(t, os.Stdout, testing.Verbose())
	defer fixture.Finalize()

{{range .TestCases}}{{if .Skipped}}
	fixture.Skip("Skipping test case: '{{.Name | sentence}}'"){{else}}
	test{{.Index}} := &{{.StructName}}{Fixture: fixture}
	test{{.Index}}.RunTestCase__(test{{.Index}}.{{.Name}}, "{{.Name | sentence}}", {{.LongRunning}}){{end}}
{{else}}	fixture.Skip("Fixture '{{.StructName}}' has no test cases.")
{{end}}}

{{if .TestCases}}func (self *{{.StructName}}) RunTestCase__(test func(), description string, longRunning bool) {
	if longRunning && testing.Short() {
		self.Skip("Skipping long-running test case: '" + description + "'")
		return
	}
	self.Describe(description){{if .TestTeardownName}}
	defer self.{{.TestTeardownName}}(){{end}}{{if .TestSetupName}}
	self.{{.TestSetupName}}(){{end}}
	test()
}
{{end}}{{/*if .TestCases*/}}`)

//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

const header = `//////////////////////////////////////////////////////////////////////////////
// Generated Code ////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

package %s

import (
	"os"
	"testing"

	"github.com/smartystreets/gunit"
)
`

const footer = `

func init() {
	gunit.Validate("%s")
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////// Generated Code //
///////////////////////////////////////////////////////////////////////////////
`

//////////////////////////////////////////////////////////////////////////////
