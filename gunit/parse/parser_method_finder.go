package parse

import (
	"go/ast"
	"strings"
)

type FixtureMethodFinder struct {
	fixtures map[string]*Fixture
}

func NewFixtureMethodFinder(fixtures map[string]*Fixture) *FixtureMethodFinder {
	return &FixtureMethodFinder{fixtures: fixtures}
}

func (self *FixtureMethodFinder) Find(file *ast.File) map[string]*Fixture {
	ast.Walk(self, file) // Calls self.Visit(...) recursively.
	return self.fixtures
}

func (self *FixtureMethodFinder) Visit(node ast.Node) ast.Visitor {
	function, isFunction := node.(*ast.FuncDecl)
	if !isFunction {
		return self
	}

	fixture := self.resolveFixture(function)
	if fixture == nil {
		return self
	}

	if !isExportedAndVoidAndNiladic(function) {
		return self
	}

	attach(function, fixture)
	return nil
}

func (self *FixtureMethodFinder) resolveFixture(function *ast.FuncDecl) *Fixture {
	if function.Recv.NumFields() == 0 {
		return nil
	}
	receiver, isPointer := function.Recv.List[0].Type.(*ast.StarExpr)
	if !isPointer {
		return nil
	}
	fixtureName := receiver.X.(*ast.Ident).Name
	fixture, functionMatchesFixture := self.fixtures[fixtureName]
	if !functionMatchesFixture {
		return nil
	}
	return fixture
}

func isExportedAndVoidAndNiladic(function *ast.FuncDecl) bool {
	if isExported := function.Name.IsExported(); !isExported {
		return false
	}
	if isNiladic := function.Type.Params.NumFields() == 0; !isNiladic {
		return false
	}
	isVoid := function.Type.Results.NumFields() == 0
	return isVoid
}

func attach(function *ast.FuncDecl, fixture *Fixture) {
	name := function.Name.Name

	if strings.HasPrefix(name, "Test") {
		fixture.TestCases = append(fixture.TestCases, TestCase{
			Index:      len(fixture.TestCases),
			Name:       name,
			StructName: fixture.StructName,
		})

	} else if strings.HasPrefix(name, "Setup") {
		fixture.TestSetupName = name

	} else if strings.HasPrefix(name, "Teardown") {
		fixture.TestTeardownName = name

	} else if strings.HasPrefix(name, "SkipTest") {
		fixture.TestCases = append(fixture.TestCases, TestCase{
			Index:      len(fixture.TestCases),
			Name:       name,
			StructName: fixture.StructName,
			Skipped:    true,
		})
	}
}
