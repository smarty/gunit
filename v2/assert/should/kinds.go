package should

import "reflect"

var floatTypes = map[reflect.Kind]struct{}{
	reflect.Float32: {},
	reflect.Float64: {},
}

func isFloat(v any) bool {
	_, found := floatTypes[reflect.TypeOf(v).Kind()]
	return found
}

func asFloat(a any) float64 {
	v := reflect.ValueOf(a)
	if isSignedInteger(a) {
		return float64(v.Int())
	}
	if isUnsignedInteger(a) {
		return float64(v.Uint())
	}
	return v.Float()
}

var unsignedIntegerKinds = map[reflect.Kind]struct{}{
	reflect.Uint:    {},
	reflect.Uint8:   {},
	reflect.Uint16:  {},
	reflect.Uint32:  {},
	reflect.Uint64:  {},
	reflect.Uintptr: {},
}

func isUnsignedInteger(v any) bool {
	_, found := unsignedIntegerKinds[reflect.TypeOf(v).Kind()]
	return found
}

var signedIntegerKinds = map[reflect.Kind]struct{}{
	reflect.Int:   {},
	reflect.Int8:  {},
	reflect.Int16: {},
	reflect.Int32: {},
	reflect.Int64: {},
}

func isSignedInteger(v any) bool {
	_, found := signedIntegerKinds[reflect.TypeOf(v).Kind()]
	return found
}

func isInteger(v any) bool {
	return isSignedInteger(v) || isUnsignedInteger(v)
}

var numericKinds = map[reflect.Kind]struct{}{
	reflect.Int:     {},
	reflect.Int8:    {},
	reflect.Int16:   {},
	reflect.Int32:   {},
	reflect.Int64:   {},
	reflect.Uint:    {},
	reflect.Uint8:   {},
	reflect.Uint16:  {},
	reflect.Uint32:  {},
	reflect.Uint64:  {},
	reflect.Float32: {},
	reflect.Float64: {},
}

func isNumeric(v any) bool {
	of := reflect.TypeOf(v)
	if of == nil {
		return false
	}
	_, found := numericKinds[of.Kind()]
	return found
}

var kindsWithLength = []reflect.Kind{
	reflect.Map,
	reflect.Chan,
	reflect.Array,
	reflect.Slice,
	reflect.String,
}

var containerKinds = []reflect.Kind{
	reflect.Map,
	reflect.Array,
	reflect.Slice,
	reflect.String,
}

var orderedContainerKinds = []reflect.Kind{
	reflect.Array,
	reflect.Slice,
	reflect.String,
}

func kindSlice(kinds map[reflect.Kind]struct{}) (result []reflect.Kind) {
	for kind := range kinds {
		result = append(result, kind)
	}
	return result
}
