package parse

import "go/ast"

type FixtureValidator struct {
	Parent      *FixtureCollector
	FixtureName string
}

func (this *FixtureValidator) Visit(node ast.Node) ast.Visitor {
	// We start at a TypeSpec and look for an embedded pointer field: `*gunit.Fixture`.
	field, isField := node.(*ast.Field)
	if !isField {
		return this
	}
	pointer, isPointer := field.Type.(*ast.StarExpr)
	if !isPointer {
		return &NonPointerFixtureValidator{Parent: this.Parent, FixtureName: this.FixtureName}
	}

	selector, isSelector := pointer.X.(*ast.SelectorExpr)
	if !isSelector {
		return this
	}
	gunit, isGunit := selector.X.(*ast.Ident)
	if selector.Sel.Name != "Fixture" || !isGunit || gunit.Name != "gunit" {
		return this
	}
	this.Parent.Validate(this.FixtureName)
	return nil
}

///////////////////////////////////////////////////////////////////////////////

type NonPointerFixtureValidator struct {
	Parent      *FixtureCollector
	FixtureName string
}

func (this *NonPointerFixtureValidator) Visit(node ast.Node) ast.Visitor {
	selector, isSelector := node.(*ast.SelectorExpr)
	if !isSelector {
		return nil
	}
	gunit, isGunit := selector.X.(*ast.Ident)
	if selector.Sel.Name != "Fixture" || !isGunit || gunit.Name != "gunit" {
		return nil
	}
	this.Parent.Invalidate(this.FixtureName)
	return nil
}
