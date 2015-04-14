// gunit generates testing functions by scanning for xunit-style struct-based
// fixtures that implement gunit test fixtures (see github.com/smartystreets/gunit).
package main

import (
	"crypto/md5"
	"encoding/hex"
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
	flag.StringVar(&importPath, "package", "", "The import path for the package to run `gunit` on.")
	flag.Parse()
}

func main() {
	if importPath == "" {
		gopath := os.Getenv("GOPATH")
		if gopath == "" {
			panic("must have gopath!")
		}

		working, err := os.Getwd() // TODO: or a specified import path from cli
		if err != nil {
			log.Fatal(err)
		}

		importPath = strings.Replace(working, gopath+"/src/", "", 1)
	}

	pkg, err := build.Import(importPath, "", build.AllowBinary)
	if err != nil {
		log.Fatal(err)
	}

	fixtures := []*parse.Fixture{}
	fileInfo := []os.FileInfo{}
	for _, item := range pkg.TestGoFiles {
		if item == generate.GeneratedFilename {
			continue
		} else if strings.HasPrefix(item, ".") {
			continue
		} else if !strings.HasSuffix(item, "_test.go") {
			continue
		}
		path := filepath.Join(pkg.Dir, item)
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

	if len(fixtures) == 0 {
		return
	}

	files, err := ioutil.ReadDir(pkg.Dir)
	if err != nil {
		log.Fatal(err)
	}
	contents, err := generate.ReadFiles(pkg.Dir, generate.SelectGoFiles(files))
	if err != nil {
		log.Fatal(err)
	}
	hash := md5.Sum(contents)
	buffer := make([]byte, len(hash))
	copy(buffer, hash[:])
	checksum := hex.EncodeToString(buffer)
	generated, err := generate.TestFile(pkg.Name, fixtures, checksum)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(filepath.Join(pkg.Dir, generate.GeneratedFilename), generated, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
