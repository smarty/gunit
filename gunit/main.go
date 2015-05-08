// gunit generates testing functions by scanning for xunit-style struct-based
// fixtures that implement gunit test fixtures (see github.com/smartystreets/gunit).
package main

import (
	"bytes"
	"flag"
	"go/build"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/smartystreets/gunit/gunit/generate"
	"github.com/smartystreets/gunit/gunit/parse"
)

var importPath string

func init() {
	log.SetFlags(log.Lshortfile)
	flag.StringVar(&importPath, "package", "", "The import path of the package for which a gunit test file will be generated.")
	flag.Parse()
}

func main() {
	pkg := resolvePackage()
	fixtures := parseFixtures(pkg)
	code := generateTestFileContents(pkg, fixtures)
	writeTestFile(pkg, code)
}

func resolvePackage() *build.Package {
	importPath := resolveImportPath()

	pkg, err := build.Import(importPath, "", build.AllowBinary)
	if err != nil {
		log.Fatal(err)
	}
	return pkg
}
func resolveImportPath() string {
	if importPath != "" {
		return importPath
	}

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		log.Fatal("$GOPATH environment variable required.")
	}

	working, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return strings.Replace(working, gopath+"/src/", "", 1)
}

func parseFixtures(pkg *build.Package) []*parse.Fixture {
	fixtures := []*parse.Fixture{}
	badFixtures := new(bytes.Buffer)

	for _, item := range pkg.TestGoFiles {
		if item == generate.GeneratedFilename {
			continue
		}

		source, err := ioutil.ReadFile(filepath.Join(pkg.Dir, item))
		if err != nil {
			log.Fatal(err)
		}

		batch, err := parse.Fixtures(string(source))
		if err != nil {
			badFixtures.WriteString(err.Error())
		}

		fixtures = append(fixtures, batch...)
	}

	if badFixtures.Len() > 0 {
		log.Fatal("The following incorrectly defined fixtures and/or test methods were found:" + badFixtures.String())
	}

	return fixtures
}

func generateTestFileContents(pkg *build.Package, fixtures []*parse.Fixture) []byte {
	if len(fixtures) == 0 {
		return nil
	}

	checksum, err := generate.Checksum(pkg.Dir)
	if err != nil {
		log.Fatal(err)
	}

	generated, err := generate.TestFile(pkg.Name, fixtures, checksum)
	if err != nil {
		log.Fatal(err)
	}

	return generated
}

func writeTestFile(pkg *build.Package, code []byte) {
	filename := filepath.Join(pkg.Dir, generate.GeneratedFilename)

	if len(code) == 0 {
		removeExistingGeneratedFile(filename)
		return
	}

	err := ioutil.WriteFile(filename, code, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func removeExistingGeneratedFile(filename string) {
	if _, err := os.Stat(filename); err == nil {
		err := os.Remove(filename)
		if err != nil {
			log.Fatal(err)
		}
	}
}
