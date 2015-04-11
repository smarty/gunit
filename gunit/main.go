// gunit generates testing functions by scanning for xunit-style struct-based
// fixtures that implement gunit test fixtures (see github.com/smartystreets/gunit).
package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"go/build"

	"github.com/smartystreets/gunit/gunit/generate"
	"github.com/smartystreets/gunit/gunit/parse"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		panic("must have gopath!")
	}

	working, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	importPath := strings.Replace(working, gopath+"/src/", "", 1)
	pkg, err := build.Import(importPath, "", build.AllowBinary)
	if err != nil {
		log.Fatal(err)
	}

	fixtures := []*parse.Fixture{}
	for _, item := range pkg.TestGoFiles {
		if item == "generated_by_gunit_test.go" {
			continue
		}
		if strings.HasPrefix(item, ".") {
			continue
		} else if !strings.HasSuffix(item, "_test.go") {
			continue
		}
		source, err := ioutil.ReadFile(filepath.Join(working, item))
		if err != nil {
			log.Fatal(err)
		}

		batch, err := parse.ParseFixtures(string(source))
		if err != nil {
			log.Fatal(err)
		}

		fixtures = append(fixtures, batch...)
	}

	generated, err := generate.TestFile(pkg.Name, fixtures)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(filepath.Join(working, "generated_by_gunit_test.go"), generated, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: decide if we are working in the current directory (later: or if we need to derive a directory from an import path (command line flag).)
	// TODO: parse and concatenate fixtures from each *_test.go file in the target directory.
	// TODO: if there are no go files, no test files, or no fixture structs found, don't generate anything, exit code: 0
	// TODO: generate the contents of a single *_test.go file from the parsed fixtures.
	// TODO (later): generate checksum validation code and append it to the content generated in the previous step.
	// TODO: write the combined content to a gunit_fixtures_test.go file.
}

//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////
