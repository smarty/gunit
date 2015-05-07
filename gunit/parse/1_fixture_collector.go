package parse

import (
	"go/ast"
	"strings"
)

type FixtureCollector struct {
	candidates map[string]*Fixture
	fixtures   map[string]*Fixture
}

func NewFixtureCollector() *FixtureCollector {
	return &FixtureCollector{
		candidates: make(map[string]*Fixture),
		fixtures:   make(map[string]*Fixture),
	}
}

func (self *FixtureCollector) Collect(file *ast.File) map[string]*Fixture {
	ast.Walk(self, file) // Calls self.Visit(...) recursively which populates self.fixtures
	return self.fixtures
}

func (self *FixtureCollector) Visit(node ast.Node) ast.Visitor {
	switch t := node.(type) {
	case *ast.TypeSpec:
		name := t.Name.Name
		self.candidates[name] = &Fixture{StructName: name}
		return &FixtureValidator{Parent: self, FixtureName: name}
	default:
		return self
	}
}

func (self *FixtureCollector) Validate(fixture string) {
	self.fixtures[fixture] = self.candidates[fixture]
	delete(self.candidates, fixture)

	if strings.HasPrefix(fixture, "Skip") {
		self.fixtures[fixture].Skipped = true
	}
}
