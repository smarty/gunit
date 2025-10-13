package assertions

import "fmt"

func (this *AssertionsFixture) TestShouldHaveSameTypeAs() {
	this.fail(so(1, ShouldHaveSameTypeAs), "This assertion requires exactly 1 comparison values (you provided 0).")
	this.fail(so(1, ShouldHaveSameTypeAs, 1, 2, 3), "This assertion requires exactly 1 comparison values (you provided 3).")

	this.fail(so(nil, ShouldHaveSameTypeAs, 0), "Expected '<nil>' to be: 'int' (but was: '<nil>')!")
	this.fail(so(1, ShouldHaveSameTypeAs, "asdf"), "Expected '1' to be: 'string' (but was: 'int')!")

	this.pass(so(1, ShouldHaveSameTypeAs, 0))
	this.pass(so(nil, ShouldHaveSameTypeAs, nil))
}

func (this *AssertionsFixture) TestShouldNotHaveSameTypeAs() {
	this.fail(so(1, ShouldNotHaveSameTypeAs), "This assertion requires exactly 1 comparison values (you provided 0).")
	this.fail(so(1, ShouldNotHaveSameTypeAs, 1, 2, 3), "This assertion requires exactly 1 comparison values (you provided 3).")

	this.fail(so(1, ShouldNotHaveSameTypeAs, 0), "Expected '1' to NOT be: 'int' (but it was)!")
	this.fail(so(nil, ShouldNotHaveSameTypeAs, nil), "Expected '<nil>' to NOT be: '<nil>' (but it was)!")

	this.pass(so(nil, ShouldNotHaveSameTypeAs, 0))
	this.pass(so(1, ShouldNotHaveSameTypeAs, "asdf"))
}

func (this *AssertionsFixture) TestShouldWrapError() {
	inner := fmt.Errorf("inner")
	middle := fmt.Errorf("middle(%w)", inner)
	outer := fmt.Errorf("outer(%w)", middle)

	this.fail(so(outer, ShouldWrap, "too", "many"), "This assertion requires exactly 1 comparison values (you provided 2).")
	this.fail(so(outer, ShouldWrap), "This assertion requires exactly 1 comparison values (you provided 0).")

	this.fail(so(42, ShouldWrap, 42), "The first and last arguments to this assertion must both be error values (you provided: 'int' and 'int').")
	this.fail(so(inner, ShouldWrap, 42), "The first and last arguments to this assertion must both be error values (you provided: '*errors.errorString' and 'int').")
	this.fail(so(42, ShouldWrap, inner), "The first and last arguments to this assertion must both be error values (you provided: 'int' and '*errors.errorString').")

	this.fail(so(inner, ShouldWrap, outer), `Expected error("inner") to wrap error("outer(middle(inner))") but it didn't.`)
	this.pass(so(middle, ShouldWrap, inner))
	this.pass(so(outer, ShouldWrap, middle))
	this.pass(so(outer, ShouldWrap, inner))
}
