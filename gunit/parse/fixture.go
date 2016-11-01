package parse

type Fixture struct {
	Filename   string
	StructName string

	TestSetupName    string
	TestTeardownName string

	TestCases        []TestCase
	InvalidTestCases []string

	InvalidNonPointer bool
}

type TestCase struct {
	Index       int
	Name        string
	StructName  string
	Skipped     bool
	LongRunning bool

	InvalidNonPointer bool
}
