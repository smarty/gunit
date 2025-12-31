package should

type specification interface {
	assertable(a, b any) bool
	passes(a, b any) bool
}
