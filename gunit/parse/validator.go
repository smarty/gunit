package parse

import "go/ast"

type FixtureValidator struct {
	Parent      *FixtureCollector
	FixtureName string
}

func (self *FixtureValidator) Visit(node ast.Node) ast.Visitor {
	// We start at a TypeSpec and look for an embedded pointer field: `*gunit.Fixture`.
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
	if selector.Sel.Name != "Fixture" || !isGunit || gunit.Name != "gunit" {
		return self
	}
	self.Parent.Validate(self.FixtureName)
	return nil
}
