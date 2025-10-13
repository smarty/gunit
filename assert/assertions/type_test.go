package assertions

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (this *AssertionsFixture) TestShouldHaveSameTypeAs() {
	this.fail(so(1, ShouldHaveSameTypeAs), "This assertion requires exactly 1 comparison values (you provided 0).")
	this.fail(so(1, ShouldHaveSameTypeAs, 1, 2, 3), "This assertion requires exactly 1 comparison values (you provided 3).")

	this.fail(so(nil, ShouldHaveSameTypeAs, 0), "int|<nil>|Expected '<nil>' to be: 'int' (but was: '<nil>')!")
	this.fail(so(1, ShouldHaveSameTypeAs, "asdf"), "string|int|Expected '1' to be: 'string' (but was: 'int')!")

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

func (this *AssertionsFixture) TestShouldImplement() {
	var ioReader *io.Reader = nil
	var response http.Response = http.Response{}
	var responsePtr *http.Response = new(http.Response)
	var reader = bytes.NewBufferString("")

	this.fail(so(reader, ShouldImplement), "This assertion requires exactly 1 comparison values (you provided 0).")
	this.fail(so(reader, ShouldImplement, ioReader, ioReader), "This assertion requires exactly 1 comparison values (you provided 2).")
	this.fail(so(reader, ShouldImplement, ioReader, ioReader, ioReader), "This assertion requires exactly 1 comparison values (you provided 3).")

	this.fail(so(reader, ShouldImplement, "foo"), shouldCompareWithInterfacePointer)
	this.fail(so(reader, ShouldImplement, 1), shouldCompareWithInterfacePointer)
	this.fail(so(reader, ShouldImplement, nil), shouldCompareWithInterfacePointer)

	this.fail(so(nil, ShouldImplement, ioReader), shouldNotBeNilActual)
	this.fail(so(1, ShouldImplement, ioReader), "Expected: 'io.Reader interface support'\nActual:   '*int' does not implement the interface!")

	this.fail(so(response, ShouldImplement, ioReader), "Expected: 'io.Reader interface support'\nActual:   '*http.Response' does not implement the interface!")
	this.fail(so(responsePtr, ShouldImplement, ioReader), "Expected: 'io.Reader interface support'\nActual:   '*http.Response' does not implement the interface!")
	this.pass(so(reader, ShouldImplement, ioReader))
	this.pass(so(reader, ShouldImplement, (*io.Reader)(nil)))
}

func (this *AssertionsFixture) TestShouldNotImplement() {
	var ioReader *io.Reader = nil
	var response http.Response = http.Response{}
	var responsePtr *http.Response = new(http.Response)
	var reader io.Reader = bytes.NewBufferString("")

	this.fail(so(reader, ShouldNotImplement), "This assertion requires exactly 1 comparison values (you provided 0).")
	this.fail(so(reader, ShouldNotImplement, ioReader, ioReader), "This assertion requires exactly 1 comparison values (you provided 2).")
	this.fail(so(reader, ShouldNotImplement, ioReader, ioReader, ioReader), "This assertion requires exactly 1 comparison values (you provided 3).")

	this.fail(so(reader, ShouldNotImplement, "foo"), shouldCompareWithInterfacePointer)
	this.fail(so(reader, ShouldNotImplement, 1), shouldCompareWithInterfacePointer)
	this.fail(so(reader, ShouldNotImplement, nil), shouldCompareWithInterfacePointer)

	this.fail(so(reader, ShouldNotImplement, ioReader), "Expected         '*bytes.Buffer'\nto NOT implement   'io.Reader' (but it did)!")
	this.fail(so(nil, ShouldNotImplement, ioReader), shouldNotBeNilActual)
	this.pass(so(1, ShouldNotImplement, ioReader))
	this.pass(so(response, ShouldNotImplement, ioReader))
	this.pass(so(responsePtr, ShouldNotImplement, ioReader))
}

func (this *AssertionsFixture) TestShouldBeError() {
	this.fail(so(nil, ShouldBeError, "too", "many"), "This assertion allows 1 or fewer comparison values (you provided 2).")

	this.fail(so(1, ShouldBeError), "Expected an error value (but was 'int' instead)!")
	this.fail(so(nil, ShouldBeError), "Expected an error value (but was '<nil>' instead)!")

	error1 := errors.New("Message")

	this.fail(so(error1, ShouldBeError, 42), "The final argument to this assertion must be a string or an error value (you provided: 'int').")
	this.fail(so(error1, ShouldBeError, "Wrong error message"), `Wrong error message|Message|Expected: "Wrong error message" Actual: "Message" (Should equal)!`)

	this.pass(so(error1, ShouldBeError))
	this.pass(so(error1, ShouldBeError, error1))
	this.pass(so(error1, ShouldBeError, error1.Error()))
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
