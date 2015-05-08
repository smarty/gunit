package parse

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

//////////////////////////////////////////////////////////////////////////////

func Fixtures(code string) ([]*Fixture, error) {
	fileset := token.NewFileSet()
	file, err := parser.ParseFile(fileset, "", code, 0)
	if err != nil {
		return nil, err
	}
	// ast.Print(fileset, file) // helps with debugging...
	return findAndListFixtures(file)
}

func findAndListFixtures(file *ast.File) ([]*Fixture, error) {
	collection := NewFixtureCollector().Collect(file)
	collection = NewFixtureMethodFinder(collection).Find(file)
	collection = NewFixtureSetupTeardownFinder(collection).Find(file)

	return listFixtures(collection)
}

func listFixtures(collection map[string]*Fixture) ([]*Fixture, error) {
	fixtures := []*Fixture{}
	errorMessage := new(bytes.Buffer)

	for _, fixture := range collection {
		accountForErrors(errorMessage, fixture)
		fixtures = append(fixtures, fixture)
	}
	if errorMessage.Len() > 0 {
		return nil, errors.New(errorMessage.String())
	}
	return fixtures, nil
}

func accountForErrors(errorMessage *bytes.Buffer, fixture *Fixture) {
	if fixture.InvalidNonPointer {
		errorMessage.WriteString(fmt.Sprintf("\n- The fixture struct '%s' must embed `*gunit.Fixture` (pointer), not `gunit.Fixture`.", fixture.StructName))
	}
	for _, methodName := range fixture.InvalidTestCases {
		errorMessage.WriteString(fmt.Sprintf("\n- %s.%s must be declared with a pointer receiver.", fixture.StructName, methodName))
	}
}
