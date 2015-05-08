// gunit generates testing functions by scanning for xunit-style struct-based
// fixtures that implement gunit test fixtures (see github.com/smartystreets/gunit).
package main

import (
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
			log.Fatal(err)
		}

		fixtures = append(fixtures, batch...)
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
	if len(code) == 0 {
		return
	}
	err := ioutil.WriteFile(filepath.Join(pkg.Dir, generate.GeneratedFilename), code, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
