package parse

import (
	"fmt"
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

	// ast.Print(fileset, file)

	collection := NewFixtureCollector().Collect(file)
	collection = NewFixtureMethodFinder(collection).Find(file)
	collection = NewFixtureSetupTeardownFinder(collection).Find(file)

	fixtures := []*Fixture{}
	for _, fixture := range collection {
		if fixture.InvalidNonPointer {
			return nil, fmt.Errorf("The fixture struct '%s' must embed *gunit.Fixture, not gunit.Fixture.", fixture.StructName)
		}
		fixtures = append(fixtures, fixture)
	}
	return fixtures, nil
}

//////////////////////////////////////////////////////////////////////////////
