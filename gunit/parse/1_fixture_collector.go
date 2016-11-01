package parse

import "go/ast"

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

func (this *FixtureCollector) Collect(file *ast.File) map[string]*Fixture {
	ast.Walk(this, file) // Calls this.Visit(...) recursively which populates this.fixtures
	return this.fixtures
}

func (this *FixtureCollector) Visit(node ast.Node) ast.Visitor {
	switch t := node.(type) {
	case *ast.TypeSpec:
		name := t.Name.Name
		this.candidates[name] = &Fixture{StructName: name}
		return &FixtureValidator{Parent: this, FixtureName: name}
	default:
		return this
	}
}

func (this *FixtureCollector) Validate(fixture string) {
	this.fixtures[fixture] = this.candidates[fixture]
	delete(this.candidates, fixture)
}

func (this *FixtureCollector) Invalidate(fixture string) {
	this.candidates[fixture].InvalidNonPointer = true
	this.Validate(fixture)
}
