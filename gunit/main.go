// gunit generates testing functions by scanning for xunit-style struct-based
// fixtures that implement gunit test fixtures (see github.com/smartystreets/gunit).
package main

import (
	"go/build"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

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

	working, err := os.Getwd() // TODO: or a specified import path from cli
	if err != nil {
		log.Fatal(err)
	}

	importPath := strings.Replace(working, gopath+"/src/", "", 1)
	pkg, err := build.Import(importPath, "", build.AllowBinary)
	if err != nil {
		log.Fatal(err)
	}

	fixtures := []*parse.Fixture{}
	fileInfo := []os.FileInfo{}
	for _, item := range pkg.TestGoFiles {
		if item == "generated_by_gunit_test.go" {
			continue
		} else if strings.HasPrefix(item, ".") {
			continue
		} else if !strings.HasSuffix(item, "_test.go") {
			continue
		}
		path := filepath.Join(working, item)
		info, err := os.Stat(path)
		if err != nil {
			log.Fatal(err)
		}
		fileInfo = append(fileInfo, info)
		source, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}

		batch, err := parse.ParseFixtures(string(source))
		if err != nil {
			log.Fatal(err)
		}

		fixtures = append(fixtures, batch...)
	}

	files, err := ioutil.ReadDir(working)
	if err != nil {
		log.Fatal(err)
	}
	checksum := generate.Checksum(generate.SelectGoFiles(files))
	generated, err := generate.TestFile(pkg.Name, fixtures, checksum)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(filepath.Join(working, "generated_by_gunit_test.go"), generated, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
