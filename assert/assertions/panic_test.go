package assertions

import (
	"errors"
	"fmt"
)

func (this *AssertionsFixture) TestShouldPanic() {
	this.fail(so(func() {}, ShouldPanic, 1), "This assertion requires exactly 0 comparison values (you provided 1).")
	this.fail(so(func() {}, ShouldPanic, 1, 2, 3), "This assertion requires exactly 0 comparison values (you provided 3).")

	this.fail(so(1, ShouldPanic), shouldUseVoidNiladicFunction)
	this.fail(so(func(i int) {}, ShouldPanic), shouldUseVoidNiladicFunction)
	this.fail(so(func() int { panic("hi") }, ShouldPanic), shouldUseVoidNiladicFunction)

	this.fail(so(func() {}, ShouldPanic), shouldHavePanicked)
	this.pass(so(func() { panic("hi") }, ShouldPanic))
}

func (this *AssertionsFixture) TestShouldNotPanic() {
	this.fail(so(func() {}, ShouldNotPanic, 1), "This assertion requires exactly 0 comparison values (you provided 1).")
	this.fail(so(func() {}, ShouldNotPanic, 1, 2, 3), "This assertion requires exactly 0 comparison values (you provided 3).")

	this.fail(so(1, ShouldNotPanic), shouldUseVoidNiladicFunction)
	this.fail(so(func(i int) {}, ShouldNotPanic), shouldUseVoidNiladicFunction)

	this.fail(so(func() { panic("hi") }, ShouldNotPanic), fmt.Sprintf(shouldNotHavePanicked, "hi"))
	this.pass(so(func() {}, ShouldNotPanic))
}

func (this *AssertionsFixture) TestShouldPanicWith() {
	this.fail(so(func() {}, ShouldPanicWith), "This assertion requires exactly 1 comparison values (you provided 0).")
	this.fail(so(func() {}, ShouldPanicWith, 1, 2, 3), "This assertion requires exactly 1 comparison values (you provided 3).")

	this.fail(so(1, ShouldPanicWith, 1), shouldUseVoidNiladicFunction)
	this.fail(so(func(i int) {}, ShouldPanicWith, "hi"), shouldUseVoidNiladicFunction)
	this.fail(so(func() {}, ShouldPanicWith, "bye"), shouldHavePanicked)
	this.fail(so(func() { panic("hi") }, ShouldPanicWith, "bye"), "bye|hi|Expected func() to panic with 'bye' (but it panicked with 'hi')!")
	this.fail(so(func() { panic(errInner) }, ShouldPanicWith, errOuter), "outer inner|inner|Expected func() to panic with 'outer inner' (but it panicked with 'inner')!")

	this.pass(so(func() { panic("hi") }, ShouldPanicWith, "hi"))
	this.pass(so(func() { panic(errOuter) }, ShouldPanicWith, errInner))
}

func (this *AssertionsFixture) TestShouldNotPanicWith() {
	this.fail(so(func() {}, ShouldNotPanicWith), "This assertion requires exactly 1 comparison values (you provided 0).")
	this.fail(so(func() {}, ShouldNotPanicWith, 1, 2, 3), "This assertion requires exactly 1 comparison values (you provided 3).")

	this.fail(so(1, ShouldNotPanicWith, 1), shouldUseVoidNiladicFunction)
	this.fail(so(func(i int) {}, ShouldNotPanicWith, "hi"), shouldUseVoidNiladicFunction)
	this.fail(so(func() { panic("hi") }, ShouldNotPanicWith, "hi"), "Expected func() NOT to panic with 'hi' (but it did)!")
	this.fail(so(func() { panic(errOuter) }, ShouldNotPanicWith, errInner), "Expected func() NOT to panic with 'inner' (but it did)!")

	this.pass(so(func() {}, ShouldNotPanicWith, "bye"))
	this.pass(so(func() { panic("hi") }, ShouldNotPanicWith, "bye"))
	this.pass(so(func() { panic(errInner) }, ShouldNotPanicWith, errOuter))

}

var (
	errInner = errors.New("inner")
	errOuter = fmt.Errorf("outer %w", errInner)
)
