package parse

type Fixture struct {
	StructName          string
	FixtureSetupName    string
	FixtureTeardownName string

	TestSetupName    string
	TestTeardownName string

	TestCaseNames        []string
	SkippedTestCaseNames []string
	FocusedTestCaseNames []string
}
