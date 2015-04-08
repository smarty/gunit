package parser

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
)

//////////////////////////////////////////////////////////////////////////////

func ParseFixtures(code string) ([]*Fixture, error) {
	fileset := token.NewFileSet()
	file, err := parser.ParseFile(fileset, "", code, 0)
	Debug(fileset, file)
	if err != nil {
		return nil, err
	}

	collection := NewFixtureCollector().Collect(file)
	collection = NewFixtureMethodFinder(collection).Find(file)

	fixtures := []*Fixture{}
	for _, fixture := range collection {
		fixtures = append(fixtures, fixture)
	}
	return fixtures, nil
}

//////////////////////////////////////////////////////////////////////////////

func Debug(fileset *token.FileSet, file *ast.File) {
	if DEBUG {
		buffer := &bytes.Buffer{}
		ast.Fprint(buffer, fileset, file, nil)
		ioutil.WriteFile("ast.txt", buffer.Bytes(), 0644)
	}
}

var DEBUG = false
