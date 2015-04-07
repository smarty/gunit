// gunit generates testing functions by scanning for xunit-style struct-based
// fixtures that implement gunit test fixtures (see github.com/smartystreets/gunit).
package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	sourceFile := "example_input_test.go"
	source, err := ioutil.ReadFile(sourceFile)
	fixtures, err := ParseFixtures(string(source))
	fatal(err)
	fmt.Println(fixtures)
}

//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func debug(fileset *token.FileSet, file *ast.File) {
	ast.Print(fileset, file)
}

func (self Fixture) String() string {
	buffer := bytes.NewBufferString("Fixture Name: " + self.StructName + "\n")
	return strings.TrimSpace(buffer.String())
}
