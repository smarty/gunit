package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type Fixture struct {
	StructName       string
	TestSetupName    string
	TestTeardownName string
	TestCaseNames    []string
}

//////////////////////////////////////////////////////////////////////////////

func ParseFixtures(code string) ([]Fixture, error) {
	fileset := token.NewFileSet()
	file, err := parser.ParseFile(fileset, "", code, 0)
	if err != nil {
		return nil, err
	}

	return NewFixtureCollector().Collect(file), nil
}

type FixtureCollector struct {
	candidates map[string]Fixture
	fixtures   map[string]Fixture
}

func NewFixtureCollector() *FixtureCollector {
	return &FixtureCollector{
		candidates: make(map[string]Fixture),
		fixtures:   make(map[string]Fixture),
	}
}

func (self *FixtureCollector) Collect(file *ast.File) []Fixture {
	ast.Walk(self, file) // Calls self.Visit(...) recursively which populates self.fixtures
	fixtures := []Fixture{}
	for _, fixture := range self.fixtures {
		fixtures = append(fixtures, fixture)
	}
	return fixtures
}

func (self *FixtureCollector) Visit(node ast.Node) ast.Visitor {
	if s, ok := node.(*ast.TypeSpec); ok {
		name := s.Name.Name
		self.candidates[name] = Fixture{StructName: name}
		return &FixtureValidator{Parent: self, FixtureName: name}
	} else {
		return self
	}
}

func (self *FixtureCollector) Validate(fixture string) {
	self.fixtures[fixture] = self.candidates[fixture]
	delete(self.candidates, fixture)
}
func (self *FixtureCollector) RegisterSetup(fixture, setup string) {
	// TODO
}
func (self *FixtureCollector) RegisterTeardown(fixture, teardown string) {
	// TODO
}
func (self *FixtureCollector) RegisterTestFunction(fixture, function string) {
	// TODO
}

//////////////////////////////////////////////////////////////////////////////

type FixtureValidator struct {
	Parent      *FixtureCollector
	FixtureName string
}

func (self *FixtureValidator) Visit(node ast.Node) ast.Visitor {
	// We start at a TypeSpec and look for an embedded pointer field: `*gunit.TestCase`.
	field, isField := node.(*ast.Field)
	if !isField {
		return self
	}
	pointer, isPointer := field.Type.(*ast.StarExpr)
	if !isPointer {
		return self
	}
	selector, isSelector := pointer.X.(*ast.SelectorExpr)
	if !isSelector {
		return self
	}
	gunit, isGunit := selector.X.(*ast.Ident)
	if selector.Sel.Name != "TestCase" || !isGunit || gunit.Name != "gunit" {
		return self
	}
	self.Parent.Validate(self.FixtureName)
	return nil
}
