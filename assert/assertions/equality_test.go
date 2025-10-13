package assertions

import (
	"time"
)

func (this *AssertionsFixture) TestShouldEqual() {
	this.fail(so(1, ShouldEqual), "This assertion requires exactly 1 comparison values (you provided 0).")
	this.fail(so(1, ShouldEqual, 1, 2), "This assertion requires exactly 1 comparison values (you provided 2).")
	this.fail(so(1, ShouldEqual, 1, 2, 3), "This assertion requires exactly 1 comparison values (you provided 3).")

	this.pass(so(1, ShouldEqual, 1))
	this.fail(so(1, ShouldEqual, 2), "Expected: 2 Actual: 1 (Should equal)!")
	this.fail(so(1, ShouldEqual, "1"), `Expected: "1" Actual: 1 (Should equal)!`)

	this.pass(so(nil, ShouldEqual, nil))

	this.pass(so(true, ShouldEqual, true))
	this.fail(so(true, ShouldEqual, false), "Expected: false Actual: true (Should equal)!")

	this.pass(so("hi", ShouldEqual, "hi"))
	this.fail(so("hi", ShouldEqual, "bye"), `Expected: "bye" Actual: "hi" (Should equal)!`)

	this.pass(so(42, ShouldEqual, uint(42)))

	this.fail(so(Thing1{"hi"}, ShouldEqual, Thing1{}), `Expected: assertions.Thing1{a:""} Actual: assertions.Thing1{a:"hi"} (Should equal)! Diff: 'assertions.Thing1{a:"hi"}'`)
	this.pass(so(Thing1{"hi"}, ShouldEqual, Thing1{"hi"}))
	this.pass(so(&Thing1{"hi"}, ShouldEqual, &Thing1{"hi"}))

	this.fail(so(Thing1{}, ShouldEqual, Thing2{}), `Expected: assertions.Thing2{a:""} Actual: assertions.Thing1{a:""} (Should equal)! Diff: 'assertions.Thing21{a:""}'`)
}
func (this *AssertionsFixture) TestShouldEqual_Issue53() {
	a := func() {}
	b := a
	this.pass(so(a, ShouldEqual, b))
	this.fail(so(a, ShouldEqual, func() {}), "[no-check]")
}
func (this *AssertionsFixture) TestShouldEqual_Issue22_GoConveyIssue437() {
	// https://github.com/smartystreets/goconvey/issues/437#issue-166724312
	type A struct {
		UID      string
		CID      string
		Metadata map[string]any
	}
	this.pass(so(
		A{UID: "joe", CID: "cid", Metadata: map[string]any{"key": "value string", "key int": 3, "key array": []any{2, "blabla"}}}, ShouldEqual,
		A{UID: "joe", CID: "cid", Metadata: map[string]any{"key": "value string", "key int": 3, "key array": []any{2, "blabla"}}}))

	// https://github.com/smartystreets/goconvey/issues/437#issuecomment-258787465
	this.pass(so(
		map[string]any{"basecost": 12000, "basemin": 60, "initialcost": 0, "initialmin": 0, "overtimecost": 0, "overtimemin": 0}, ShouldEqual,
		map[string]any{"basecost": 12000, "basemin": 60, "initialcost": 0, "initialmin": 0, "overtimecost": 0, "overtimemin": 0}))

	// https://github.com/smartystreets/goconvey/issues/437#issuecomment-263060222
	this.pass(so(
		map[string]any{"fields": []string{}, "limit": 25, "offset": 0}, ShouldEqual,
		map[string]any{"fields": []string{}, "limit": 25, "offset": 0}))
}
func (this *AssertionsFixture) TestShouldEqual_TimeValues() {
	var (
		gopherCon, _ = time.LoadLocation("America/Denver")
		elsewhere, _ = time.LoadLocation("America/New_York")

		timeNow          = time.Now().In(gopherCon)
		timeNowElsewhere = timeNow.In(elsewhere)
		timeLater        = timeNow.Add(time.Nanosecond)
	)

	this.pass(so(timeNow, ShouldEqual, timeNowElsewhere)) // Time.Equal method used to determine exact instant.
	this.pass(so(timeNow, ShouldNotEqual, timeLater))
}

func (this *AssertionsFixture) TestShouldNotEqual() {
	this.fail(so(1, ShouldNotEqual), "This assertion requires exactly 1 comparison values (you provided 0).")
	this.fail(so(1, ShouldNotEqual, 1, 2), "This assertion requires exactly 1 comparison values (you provided 2).")
	this.fail(so(1, ShouldNotEqual, 1, 2, 3), "This assertion requires exactly 1 comparison values (you provided 3).")

	this.pass(so(1, ShouldNotEqual, 2))
	this.pass(so(1, ShouldNotEqual, "1"))
	this.fail(so(1, ShouldNotEqual, 1), "Expected '1' to NOT equal '1' (but it did)!")

	this.pass(so(true, ShouldNotEqual, false))
	this.fail(so(true, ShouldNotEqual, true), "Expected 'true' to NOT equal 'true' (but it did)!")

	this.pass(so("hi", ShouldNotEqual, "bye"))
	this.fail(so("hi", ShouldNotEqual, "hi"), "Expected 'hi' to NOT equal 'hi' (but it did)!")

	this.fail(so(&Thing1{"hi"}, ShouldNotEqual, &Thing1{"hi"}), "Expected '&{hi}' to NOT equal '&{hi}' (but it did)!")
	this.fail(so(Thing1{"hi"}, ShouldNotEqual, Thing1{"hi"}), "Expected '{hi}' to NOT equal '{hi}' (but it did)!")
	this.fail(so(Thing1{}, ShouldNotEqual, Thing1{}), "Expected '{}' to NOT equal '{}' (but it did)!")
	this.pass(so(Thing1{}, ShouldNotEqual, Thing2{}))
}

func (this *AssertionsFixture) TestShouldAlmostEqual() {
	this.fail(so(1, ShouldAlmostEqual), "This assertion requires exactly one comparison value and an optional delta (you provided neither)")
	this.fail(so(1, ShouldAlmostEqual, 1, 2, 3), "This assertion requires exactly one comparison value and an optional delta (you provided more values)")
	this.fail(so(1, ShouldAlmostEqual, "1"), "The comparison value must be a numerical type, but was: string")
	this.fail(so(1, ShouldAlmostEqual, 1, "1"), "The delta value must be a numerical type, but was: string")
	this.fail(so("1", ShouldAlmostEqual, 1), "The actual value must be a numerical type, but was: string")

	// with the default delta
	this.pass(so(0.99999999999999, ShouldAlmostEqual, uint(1)))
	this.pass(so(1, ShouldAlmostEqual, 0.99999999999999))
	this.pass(so(1.3612499999999996, ShouldAlmostEqual, 1.36125))
	this.pass(so(0.7285312499999999, ShouldAlmostEqual, 0.72853125))
	this.fail(so(1, ShouldAlmostEqual, .99), "Expected '1' to almost equal '0.99' (but it didn't)!")

	// with a different delta
	this.pass(so(100.0, ShouldAlmostEqual, 110.0, 10.0))
	this.fail(so(100.0, ShouldAlmostEqual, 111.0, 10.5), "Expected '100' to almost equal '111' (but it didn't)!")

	// various ints should work
	this.pass(so(100, ShouldAlmostEqual, 100.0))
	this.pass(so(int(100), ShouldAlmostEqual, 100.0))
	this.pass(so(int8(100), ShouldAlmostEqual, 100.0))
	this.pass(so(int16(100), ShouldAlmostEqual, 100.0))
	this.pass(so(int32(100), ShouldAlmostEqual, 100.0))
	this.pass(so(int64(100), ShouldAlmostEqual, 100.0))
	this.pass(so(uint(100), ShouldAlmostEqual, 100.0))
	this.pass(so(uint8(100), ShouldAlmostEqual, 100.0))
	this.pass(so(uint16(100), ShouldAlmostEqual, 100.0))
	this.pass(so(uint32(100), ShouldAlmostEqual, 100.0))
	this.pass(so(uint64(100), ShouldAlmostEqual, 100.0))
	this.pass(so(100, ShouldAlmostEqual, 100.0))
	this.fail(so(100, ShouldAlmostEqual, 99.0), "Expected '100' to almost equal '99' (but it didn't)!")

	// floats should work
	this.pass(so(float64(100.0), ShouldAlmostEqual, float32(100.0)))
	this.fail(so(float32(100.0), ShouldAlmostEqual, 99.0, float32(0.1)), "Expected '100' to almost equal '99' (but it didn't)!")
}

func (this *AssertionsFixture) TestShouldNotAlmostEqual() {
	this.fail(so(1, ShouldNotAlmostEqual), "This assertion requires exactly one comparison value and an optional delta (you provided neither)")
	this.fail(so(1, ShouldNotAlmostEqual, 1, 2, 3), "This assertion requires exactly one comparison value and an optional delta (you provided more values)")

	// with the default delta
	this.fail(so(1, ShouldNotAlmostEqual, .99999999999999), "Expected '1' to NOT almost equal '0.99999999999999' (but it did)!")
	this.fail(so(1.3612499999999996, ShouldNotAlmostEqual, 1.36125), "Expected '1.3612499999999996' to NOT almost equal '1.36125' (but it did)!")
	this.pass(so(1, ShouldNotAlmostEqual, .99))

	// with a different delta
	this.fail(so(100.0, ShouldNotAlmostEqual, 110.0, 10.0), "Expected '100' to NOT almost equal '110' (but it did)!")
	this.pass(so(100.0, ShouldNotAlmostEqual, 111.0, 10.5))

	// ints should work
	this.fail(so(100, ShouldNotAlmostEqual, 100.0), "Expected '100' to NOT almost equal '100' (but it did)!")
	this.pass(so(100, ShouldNotAlmostEqual, 99.0))

	// float32 should work
	this.fail(so(float64(100.0), ShouldNotAlmostEqual, float32(100.0)), "Expected '100' to NOT almost equal '100' (but it did)!")
	this.pass(so(float32(100.0), ShouldNotAlmostEqual, 99.0, float32(0.1)))
}

func (this *AssertionsFixture) TestShouldBeNil() {
	this.fail(so(nil, ShouldBeNil, nil, nil, nil), "This assertion requires exactly 0 comparison values (you provided 3).")
	this.fail(so(nil, ShouldBeNil, nil), "This assertion requires exactly 0 comparison values (you provided 1).")

	this.pass(so(nil, ShouldBeNil))
	this.fail(so(1, ShouldBeNil), "Expected: nil Actual: '1'")

	var thing ThingInterface
	this.pass(so(thing, ShouldBeNil))
	thing = &ThingImplementation{}
	this.fail(so(thing, ShouldBeNil), "Expected: nil Actual: '&{}'")

	var thingOne *Thing1
	this.pass(so(thingOne, ShouldBeNil))

	var nilSlice []int = nil
	this.pass(so(nilSlice, ShouldBeNil))

	var nilMap map[string]string = nil
	this.pass(so(nilMap, ShouldBeNil))

	var nilChannel chan int = nil
	this.pass(so(nilChannel, ShouldBeNil))

	var nilFunc func() = nil
	this.pass(so(nilFunc, ShouldBeNil))

	var nilInterface any = nil
	this.pass(so(nilInterface, ShouldBeNil))
}

func (this *AssertionsFixture) TestShouldNotBeNil() {
	this.fail(so(nil, ShouldNotBeNil, nil, nil, nil), "This assertion requires exactly 0 comparison values (you provided 3).")
	this.fail(so(nil, ShouldNotBeNil, nil), "This assertion requires exactly 0 comparison values (you provided 1).")

	this.fail(so(nil, ShouldNotBeNil), "Expected '<nil>' to NOT be nil (but it was)!")
	this.pass(so(1, ShouldNotBeNil))

	var thing ThingInterface
	this.fail(so(thing, ShouldNotBeNil), "Expected '<nil>' to NOT be nil (but it was)!")
	thing = &ThingImplementation{}
	this.pass(so(thing, ShouldNotBeNil))
}

func (this *AssertionsFixture) TestShouldBeTrue() {
	this.fail(so(true, ShouldBeTrue, 1, 2, 3), "This assertion requires exactly 0 comparison values (you provided 3).")
	this.fail(so(true, ShouldBeTrue, 1), "This assertion requires exactly 0 comparison values (you provided 1).")

	this.fail(so(false, ShouldBeTrue), "Expected: true Actual: false")
	this.fail(so(1, ShouldBeTrue), "Expected: true Actual: 1")
	this.pass(so(true, ShouldBeTrue))
}

func (this *AssertionsFixture) TestShouldBeFalse() {
	this.fail(so(false, ShouldBeFalse, 1, 2, 3), "This assertion requires exactly 0 comparison values (you provided 3).")
	this.fail(so(false, ShouldBeFalse, 1), "This assertion requires exactly 0 comparison values (you provided 1).")

	this.fail(so(true, ShouldBeFalse), "Expected: false Actual: true")
	this.fail(so(1, ShouldBeFalse), "Expected: false Actual: 1")
	this.pass(so(false, ShouldBeFalse))
}

func (this *AssertionsFixture) TestShouldBeZeroValue() {
	this.fail(so(0, ShouldBeZeroValue, 1, 2, 3), "This assertion requires exactly 0 comparison values (you provided 3).")
	this.fail(so(false, ShouldBeZeroValue, true), "This assertion requires exactly 0 comparison values (you provided 1).")

	this.fail(so(1, ShouldBeZeroValue), "'1' should have been the zero value")                             // "Expected: (zero value) Actual: 1")
	this.fail(so(true, ShouldBeZeroValue), "'true' should have been the zero value")                       // "Expected: (zero value) Actual: true")
	this.fail(so("123", ShouldBeZeroValue), "'123' should have been the zero value")                       // "Expected: (zero value) Actual: 123")
	this.fail(so(" ", ShouldBeZeroValue), "' ' should have been the zero value")                           // "Expected: (zero value) Actual:  ")
	this.fail(so([]string{"Nonempty"}, ShouldBeZeroValue), "'[Nonempty]' should have been the zero value") // "Expected: (zero value) Actual: [Nonempty]")
	this.fail(so(struct{ a string }{a: "asdf"}, ShouldBeZeroValue), "'{a:asdf}' should have been the zero value")
	this.pass(so(0, ShouldBeZeroValue))
	this.pass(so(false, ShouldBeZeroValue))
	this.pass(so("", ShouldBeZeroValue))
	this.pass(so(struct{}{}, ShouldBeZeroValue))
}

func (this *AssertionsFixture) TestShouldNotBeZeroValue() {
	this.fail(so(0, ShouldNotBeZeroValue, 1, 2, 3), "This assertion requires exactly 0 comparison values (you provided 3).")
	this.fail(so(false, ShouldNotBeZeroValue, true), "This assertion requires exactly 0 comparison values (you provided 1).")

	this.fail(so(0, ShouldNotBeZeroValue), "'0' should NOT have been the zero value")
	this.fail(so(false, ShouldNotBeZeroValue), "'false' should NOT have been the zero value")
	this.fail(so("", ShouldNotBeZeroValue), "'' should NOT have been the zero value")
	this.fail(so(struct{}{}, ShouldNotBeZeroValue), "'{}' should NOT have been the zero value")

	this.pass(so(1, ShouldNotBeZeroValue))
	this.pass(so(true, ShouldNotBeZeroValue))
	this.pass(so("123", ShouldNotBeZeroValue))
	this.pass(so(" ", ShouldNotBeZeroValue))
	this.pass(so([]string{"Nonempty"}, ShouldNotBeZeroValue))
	this.pass(so(struct{ a string }{a: "asdf"}, ShouldNotBeZeroValue))
}
