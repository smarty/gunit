package assertions

import (
	"reflect"
	"testing"
)

type Assertion struct{ t *testing.T }

func New(t *testing.T) *Assertion {
	return &Assertion{t: t}
}

func (this *Assertion) AssertEqual(expected, actual interface{}) bool {
	this.t.Helper()

	equal := expected == actual
	if !equal {
		this.t.Errorf("\nExpected: [%v]\nActual:   [%v]", expected, actual)
	}
	return equal
}

func (this *Assertion) AssertNil(value interface{}) bool {
	this.t.Helper()

	isNil := value == nil || reflect.ValueOf(value).IsNil()
	if !isNil {
		this.t.Errorf("Expected [%v] to be nil, but it wasn't.", value)
	}
	return isNil
}

func (this *Assertion) AssertTrue(value bool) bool {
	this.t.Helper()

	if !value {
		this.t.Error("Expected true, got false instead")
	}
	return value
}

func (this *Assertion) AssertFalse(value bool) bool {
	this.t.Helper()

	if value {
		this.t.Error("Expected false, got true instead")
	}
	return !value
}

func (this *Assertion) AssertDeepEqual(expected, actual interface{}) bool {
	this.t.Helper()

	equal := reflect.DeepEqual(expected, actual)
	if !equal {
		this.t.Errorf("\nExpected: [%v]\nActual:   [%v]", expected, actual)
	}
	return equal
}
