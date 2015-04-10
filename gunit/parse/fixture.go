package parse

type Fixture struct {
	Skipped    bool
	Focused    bool
	StructName string

	FixtureSetupName    string
	FixtureTeardownName string

	TestSetupName    string
	TestTeardownName string

	HasFocusedTestCases bool

	TestCases []TestCase
}

type TestCase struct {
	Index      int
	Name       string
	StructName string
	Skipped    bool
	Focused    bool
}
