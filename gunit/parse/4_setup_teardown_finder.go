package parse

import (
	"go/ast"
	"strings"
)

//////////////////////////////////////////////////////////////////////////////

type FixtureSetupTeardownFinder struct {
	fixtures map[string]*Fixture
}

func NewFixtureSetupTeardownFinder(fixtures map[string]*Fixture) *FixtureSetupTeardownFinder {
	return &FixtureSetupTeardownFinder{fixtures: fixtures}
}

func (self *FixtureSetupTeardownFinder) Find(file *ast.File) map[string]*Fixture {
	ast.Walk(self, file) // Calls self.Visit(...) recursively.
	return self.fixtures
}

func (self *FixtureSetupTeardownFinder) Visit(node ast.Node) ast.Visitor {
	self.associateFixtureSetup(NewFixtureSetupTeardownParser(node, "Setup").Parse())
	self.associateFixtureTeardown(NewFixtureSetupTeardownParser(node, "Teardown").Parse())
	return self
}
func (self *FixtureSetupTeardownFinder) associateFixtureSetup(fixtureName, setupName string) {
	if fixture, found := self.fixtures[fixtureName]; found {
		fixture.FixtureSetupName = setupName
	}
}
func (self *FixtureSetupTeardownFinder) associateFixtureTeardown(fixtureName, teardownName string) {
	if fixture, found := self.fixtures[fixtureName]; found {
		fixture.FixtureTeardownName = teardownName
	}
}

//////////////////////////////////////////////////////////////////////////////

type FixtureSetupTeardownParser struct {
	node   ast.Node
	prefix string
}

func NewFixtureSetupTeardownParser(node ast.Node, prefix string) *FixtureSetupTeardownParser {
	return &FixtureSetupTeardownParser{
		node:   node,
		prefix: prefix,
	}
}

func (self *FixtureSetupTeardownParser) Parse() (fixtureName, setupName string) {
	function, isFunction := self.node.(*ast.FuncDecl)
	if !isFunction {
		return "", ""
	}

	if !isStandalone(function) {
		return "", ""
	}

	if !isExportedAndVoidAndNiladic(function) {
		return "", ""
	}

	setupName = function.Name.Name
	if !self.functionNameIsValid(setupName) {
		return "", ""
	}

	return self.fixtureWithSetup(setupName)
}

func (self *FixtureSetupTeardownParser) fixtureWithSetup(setup string) (string, string) {
	fixture := setup[len(self.prefix):]
	return fixture, setup
}

func (self *FixtureSetupTeardownParser) functionNameIsValid(name string) bool {
	return strings.HasPrefix(name, self.prefix) && len(name) > len(self.prefix)
}

//////////////////////////////////////////////////////////////////////////////

func isStandalone(function *ast.FuncDecl) bool {
	return function.Recv == nil
}

//////////////////////////////////////////////////////////////////////////////
