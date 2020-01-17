package reports

import (
	"fmt"
	"strings"
)

func PanicReport(r interface{}, stack []byte) string {
	var builder strings.Builder
	fmt.Fprintln(&builder, "PANIC:", r)
	fmt.Fprintln(&builder, "...")

	opened, closed := false, false
	for _, line := range strings.Split(string(stack), "\n") {
		if strings.Contains(line, "/runtime/panic.go:") {
			opened = true
			continue
		}
		if !opened || closed {
			continue
		}
		if strings.Contains(line, "reflect.Value.call(0x") {
			closed = true
			continue
		}
		fmt.Fprintln(&builder, line)
	}

	return strings.TrimSpace(builder.String())

}
