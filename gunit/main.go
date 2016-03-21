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
	"path"
	"path/filepath"
	"strings"

	"github.com/smartystreets/gunit/gunit/generate"
	"github.com/smartystreets/gunit/gunit/parse"
)

var importPath string

func init() {
	logger = log.New(os.Stderr, "", log.Lshortfile)
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
	fatal(err)
	return pkg
}
func resolveImportPath() string {
	if importPath != "" {
		return importPath
	}

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		logger.Fatal("$GOPATH environment variable required.")
	}

	working, err := os.Getwd()
	fatal(err)

	dirs := strings.Split(gopath, ":")
	for _, dir := range dirs {
		srcDir := path.Join(dir, "src") + "/"
		packageName := strings.Replace(working, srcDir, "", 1)
		if packageName != working {
			return packageName
		}
	}

	logger.Fatal("Cannot determine package name from current directory; must run gunit from within a package")
	panic("Not reachable")
}

func parseFixtures(pkg *build.Package) []*parse.Fixture {
	fixtures := []*parse.Fixture{}
	badFixtures := new(bytes.Buffer)

	for _, filename := range pkg.TestGoFiles {
		if filename == generate.GeneratedFilename {
			continue
		}

		source, err := ioutil.ReadFile(filepath.Join(pkg.Dir, filename))
		fatal(err)

		batch, err := parse.Fixtures(string(source))
		if err != nil {
			badFixtures.WriteString(err.Error())
		}

		for _, fixture := range batch {
			fixture.Filename = filename
			fixtures = append(fixtures, fixture)
		}
	}

	if badFixtures.Len() > 0 {
		logger.Fatal("The following incorrectly defined fixtures and/or test methods were found:" + badFixtures.String())
	}

	return fixtures
}

func generateTestFileContents(pkg *build.Package, fixtures []*parse.Fixture) []byte {
	if len(fixtures) == 0 {
		return nil
	}

	checksum, err := generate.Checksum(pkg.Dir)
	fatal(err)

	generated, err := generate.TestFile(pkg.Name, fixtures, checksum)
	fatal(err)

	return generated
}

func writeTestFile(pkg *build.Package, code []byte) {
	filename := filepath.Join(pkg.Dir, generate.GeneratedFilename)

	if len(code) == 0 {
		removeExistingGeneratedFile(filename)
		return
	}

	err := ioutil.WriteFile(filename, code, 0644)
	fatal(err)
}

func removeExistingGeneratedFile(filename string) {
	if _, err := os.Stat(filename); err == nil {
		err := os.Remove(filename)
		fatal(err)
	}
}

func fatal(err error) {
	if err != nil {
		logger.Output(2, err.Error())
		os.Exit(1)
	}
}

var logger *log.Logger
