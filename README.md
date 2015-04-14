# gunit

### TODO:

- [X] checksums
- [X] main logic
- [X] unit test the gunit package
- [X] extend scantest to run gunit
- [ ] should the gunit command be at the root instead?\
- [ ] remove 'Test' from output?
- [ ] apostrophes in contractions
- [ ] focus? (technically the -run flag achieves this)
- [ ] better README
- [ ] extend goconvey to run gunit


-------------------------

Here's the basic idea of what the test author would implement:

```

func SetupMyFixture()    {}
func TeardownMyFixture() {}

type MyFixture struct {
	*gunit.TestCase
}

func (this *MyFixture) Setup...()    {}
func (this *MyFixture) Teardown...() {}

func (this *MyFixture) Test1...() {}
func (this *MyFixture) Test2...() {}
func (this *MyFixture) Test3...() {}

```

-------------------------

Here's what the generated code could look like (sans logging):

```

func TestMyFixture(t *testing.T) {
	defer TeardownMyFixture()
	SetupMyFixture()

	test0 := &MyFixture{TestCase: gunit.TestCase(t)}
	test0.RunTestCase(test0.Test1)

	test1 := &MyFixture{TestCase: gunit.TestCase(t)}
	test1.RunTestCase(test1.Test2)

	test2 := &MyFixture{TestCase: gunit.TestCase(t)}
	test2.RunTestCase(test2.Test3)
}

func (self *MyFixture) RunTestCase(test func()) {
	defer self.Teardown...()
	self.Setup...()
	test()
}

```