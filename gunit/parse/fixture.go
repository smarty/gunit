package parse

type Fixture struct {
	Skipped    bool
	Focused    bool
	StructName string

	FixtureSetupName    string
	FixtureTeardownName string

	TestSetupName    string
	TestTeardownName string

	TestCaseNames        []string
	SkippedTestCaseNames []string
	FocusedTestCaseNames []string
}
