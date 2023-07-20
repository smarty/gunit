package gunit

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"path"
	"runtime"
	"strings"
)

const (
	Test       = "FixtureTestName"
	Package    = "FixturePackageName"
	Title      = "FixtureTestTitle"
	RequestApi = "FixtureRequestApi"
)

type Logger struct {
	t               TestingT
	testPackageName string
}

func (f *Logger) Log() *zerolog.Event {
	return log.Log().Str(Test, f.t.Name()).Str(Package, f.testPackageName)
}

func (f *Logger) Debug() *zerolog.Event {
	return log.Debug().Str(Test, f.t.Name()).Str(Package, f.testPackageName)
}

func (f *Logger) Info() *zerolog.Event {
	return log.Info().Str(Test, f.t.Name()).Str(Package, f.testPackageName)
}
func (f *Logger) Warn() *zerolog.Event {
	return log.Warn().Str(Test, f.t.Name()).Str(Package, f.testPackageName)
}

func (f *Logger) ErrorLog() *zerolog.Event {
	return log.Error().Str(Test, f.t.Name()).Str(Package, f.testPackageName)
}

func (f *Logger) Description(desc string) {
	f.Log().Str(Title, desc).Msg("")
}

func RetrieveTestPackageName() string {
	start := 0
	pc, file, _, _ := runtime.Caller(start)
	_, fileName := path.Split(file)

	if !strings.HasSuffix(fileName, "_test.go") {
		// find caller from _test.go
		for i := start + 1; i < 30; i++ {
			pc, file, _, _ = runtime.Caller(i)
			_, fileName = path.Split(file)
			if strings.HasSuffix(fileName, "_test.go") {
				break
			}
		}
	}
	parts := strings.Split(runtime.FuncForPC(pc).Name(), "/")

	last := strings.Split(parts[len(parts)-1], ".")

	pkgPath := strings.Join(parts[:len(parts)-1], "/") + "/" + last[0]
	return pkgPath
}
