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
	sourceFile := "parse/example_input_test.go"
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
