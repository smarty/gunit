package parse

import (
	"go/parser"
	"go/token"
)

//////////////////////////////////////////////////////////////////////////////

func Fixtures(code string) ([]*Fixture, error) {
	file, err := parser.ParseFile(token.NewFileSet(), "", code, 0)
	if err != nil {
		return nil, err
	}

	collection := NewFixtureCollector().Collect(file)
	collection = NewFixtureMethodFinder(collection).Find(file)
	collection = NewFixtureSetupTeardownFinder(collection).Find(file)

	fixtures := []*Fixture{}
	for _, fixture := range collection {
		fixtures = append(fixtures, fixture)
	}
	return fixtures, nil
}

//////////////////////////////////////////////////////////////////////////////
