package parse

import (
	"go/parser"
	"go/token"
)

//////////////////////////////////////////////////////////////////////////////

func ParseFixtures(code string) ([]*Fixture, error) {
	fileset := token.NewFileSet()
	file, err := parser.ParseFile(fileset, "", code, 0)
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
