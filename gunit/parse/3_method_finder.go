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

	if function.Recv.NumFields() == 0 {
		return nil
	}

	receiver, isPointer := function.Recv.List[0].Type.(*ast.StarExpr)
	if !isPointer {
		return &FixtureMethodInvalidator{function: function.Name.Name, fixtures: self.fixtures}
	}

	fixtureName := receiver.X.(*ast.Ident).Name
	fixture, functionMatchesFixture := self.fixtures[fixtureName]
	if !functionMatchesFixture {
		return nil
	}

	if !isExportedAndVoidAndNiladic(function) {
		return self
	}

	attach(function, fixture)
	return nil
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

//////////////////////////////////////////////////////////////////////////////

type FixtureMethodInvalidator struct {
	function string
	fixtures map[string]*Fixture
}

func (self *FixtureMethodInvalidator) Visit(node ast.Node) ast.Visitor {
	receiverList, isReceiverList := node.(*ast.FieldList)
	if !isReceiverList {
		return nil
	}

	if receiverList.NumFields() != 1 {
		return nil
	}

	receiver := receiverList.List[0]

	fixtureName := receiver.Type.(*ast.Ident).Name
	fixture, functionMatchesFixture := self.fixtures[fixtureName]
	if !functionMatchesFixture {
		return nil
	}

	fixture.InvalidTestCases = append(fixture.InvalidTestCases, self.function)
	return nil
}
