package gunit

import (
	"context"
	"io"
)

// TestingT represents the functional subset from *testing.T needed by Fixture.
type TestingT interface {
	Helper()
	Name() string
	Log(args ...any)
	Fail()
	Failed() bool
	Fatalf(format string, args ...any)
	Errorf(format string, args ...any)
	Context() context.Context
	Output() io.Writer
}
