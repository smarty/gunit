package gunit

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/smartystreets/gunit/gunit/generate"
)

// Validate is called by generated code.
func Validate(checksum string) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		exit("Unable to resolve the test file from runtime.Caller(...).\n")
	}
	current, err := generate.Checksum(filepath.Dir(file))
	if err != nil {
		exit("Could not calculate checksum of current go files. Error: %s\n", err.Error())
	}
	if checksum != current {
		exit("The checksum provided [%s] does not match the current file listing [%s]. Please re-run the `gunit` command and try again.\n", checksum, current)
	}
}

func exit(message string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, message, args...)
	os.Exit(1)
}

////////////////////////////////////////////////////////////////////////////////

func init() {
	working, err := os.Getwd()
	if err != nil {
		exit("Could not resolve working directory. Error: %s\n", err)
	}
	_, err = os.Stat(filepath.Join(working, generate.GeneratedFilename))
	if err != nil {
		exit("Having written one or more gunit Fixtures in this package, please run `gunit` and try again.\n")
	}
}
