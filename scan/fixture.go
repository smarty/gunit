package scan

import "fmt"

type fixtureInfo struct {
	Filename   string
	StructName string
	TestCases  []*testCaseInfo
}

type testCaseInfo struct {
	CharacterPosition int
	LineNumber        int
	Name              string
}

func (this testCaseInfo) GoString() string {
	return fmt.Sprintf(
		`testCaseInfo{CharacterPosition: %d, LineNumber: %d, Name: %s}`,
		this.CharacterPosition,
		this.LineNumber,
		this.Name,
	)
}
