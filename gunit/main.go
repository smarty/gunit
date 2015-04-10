// gunit generates testing functions by scanning for xunit-style struct-based
// fixtures that implement gunit test fixtures (see github.com/smartystreets/gunit).
package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/smartystreets/gunit/gunit/parse"
)

func init() {
	parse.DEBUG = true
}

func main() {
	// TODO: decide if we are working in the current directory or if we need to derive a directory from an import page (command line flag).
	// TODO: parse and concatenate fixtures from each *_test.go file in the target directory.
	// TODO: if there are no go files, no test files, or no fixture structs found, don't generate anything, exit code: 0
	// TODO: generate the contents of a single *_test.go file from the parsed fixtures.
	// TODO: generate checksum validation code and append it to the content generated in the previous step.
	// TODO: write the combined content to a gunit_fixtures_test.go file.

	sourceFile := "parse/example_input_test.go.txt"
	source, err := ioutil.ReadFile(sourceFile)
	fixtures, err := parse.ParseFixtures(string(source))
	fatal(err)
	fmt.Println(fixtures)
}

//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

func fatal(err error) {
	if err != nil {
		log.Fatal("STUFF:", err)
	}
}
