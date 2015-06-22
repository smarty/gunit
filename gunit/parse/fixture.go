package parse

type Fixture struct {
	Skipped    bool
	StructName string

	FixtureSetupName    string
	FixtureTeardownName string

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
