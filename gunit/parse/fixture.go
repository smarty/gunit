package parse

type Fixture struct {
	Skipped    bool
	StructName string

	FixtureSetupName    string
	FixtureTeardownName string

	TestSetupName    string
	TestTeardownName string

	TestCases []TestCase
}

type TestCase struct {
	Index      int
	Name       string
	StructName string
	Skipped    bool
}
