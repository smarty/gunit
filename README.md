#### 声明：这个项目是基于 [smarty/gunit](https://github.com/smarty/gunit) 改造的，原作者已经不维护所以移到本仓库继续更新。

# gunit

`gunit`，又一个用于Go语言的测试工具。

> 不要再来了...（[GoConvey](http://goconvey.co)已经够疯狂了...但还挺酷，好吧我会关注一下...）

等等，这个工具有一些非常有趣的特性。它是由内置测试包提供的好东西、[GoConvey](http://goconvey.co)
项目中您所熟悉和喜爱的[断言](https://github.com/smarty/assertions)，[xUnit](https://en.wikipedia.org/wiki/XUnit)
测试风格（第一个真正的Go单元测试框架）的混合体，所有这些都与`go test`紧密结合在一起。

## 编写

> 啰哩啰嗦，，好吧。那么，为什么不只是使用标准的"testing"包呢？这个`gunit`有什么优势？

由"testing"包和`go test`工具建立的约定只允许在局部函数范围内：

```
func TestSomething(t *testing.T) {
	// 巴拉 巴拉 巴拉
}
```

这种有限的作用域使得共享方法和数据变得麻烦。如果试图保持测试简洁而短小，可能会变得混乱。以下展示了如何使用`gunit`
编写`*_test.go`文件：

```go

package examples

import (
	"time"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/bugVanisher/gunit"
)

func TestExampleFixture(t *testing.T) {
	gunit.Run(new(ExampleFixture), t)
}

type ExampleFixture struct {
	*gunit.Fixture // Required: Embedding this type is what makes the magic happen.

	// Declare useful state here (probably the stuff being tested, any fakes, etc...).
}

func (this *ExampleFixture) SetupStuff() {
	// This optional method will be executed before each "Test"
	// method (because it starts with "Setup").
}
func (this *ExampleFixture) TeardownStuff() {
	// This optional method will be executed after each "Test"
	// method (because it starts with "Teardown"), even if the test method panics.
}

func (this *ExampleFixture) FixtureSetupStuff() {
	// This optional method will be executed before all "Test"
	// method (because it starts with "FixtureSetup").
}
func (this *ExampleFixture) FixtureTeardownStuff() {
	// This optional method will be executed after all "Test"
	// method (because it starts with "FixtureTeardown"), even if any test method panics.
}

// This is an actual test case:
func (this *ExampleFixture) TestWithAssertions() {
	// Here's how to use the functions from the `should`
	// package at github.com/smarty/assertions/should
	// to perform assertions:
	this.So(42, should.Equal, 42)
	this.So("Hello, World!", should.ContainSubstring, "World")
}

func (this *ExampleFixture) SkipTestWithNothing() {
	// Because this method's name starts with 'Skip', it will be skipped.
}

func (this *ExampleFixture) LongTestSlowOperation() {
	// Because this method's name starts with 'Long', it will be skipped if `go test` is run with the `short` flag.
	time.Sleep(time.Hour)
	this.So(true, should.BeTrue)
}
```

-------------------------

> 所以，你只会看到一个传统的测试函数，而且只有一行代码（当然，还有一个结构体和它的方法）。这是怎么回事？

`gunit`允许用例编写者使用一个_结构体_"封装"一组相关测试用例，类似于[xUnit](https://en.wikipedia.org/wiki/XUnit)
。这使得进行前置和后置行为变得更简单，因为测试的所有状态可以声明为嵌入`gunit`包中的`Fixture`
类型的结构体字段。只需创建一个Test函数，并将fixture结构体的新实例与*
testing.T一起传递给gunit的Run函数，它将运行所有已定义的Test方法以及Setup和Teardown方法。另外，还可以使用FixtureSetup和FixtureTeardown方法在所有测试执行之前或之后定义一些操作。

## 核心特性

- xUnit风格：完整支持 Go 语言的 xUnit 风格用例编写
- 同时支持串行和并行：Fixture粒度的串行、并行控制，大大提高执行效率
- 丰富断言支持：提供多种断言方式
- Readable测试报告：以包为单位组织报告，查看和定位简洁高效
- 复杂场景：可根据实际情况继承Fixture，实现业务级的接口、UI自动化测试

## 引用

```go
import (
    "github.com/bugVanisher/gunit"
)
```

然后在你的测试用例工程的go.mod文件中添加如下语句：

```
replace github.com/bugVanisher/gunit => github.com/bugVanisher/gunit v2.0.1
```

-------------------------

### 日志打印

使用fixture logger中的Log/Debug/Info/Warn/Error方法打印日志，如下

```go
type MyFixture struct {
    *gunit.Fixture
}
// 所有Test方法执行前执行
func (g *MyFixture) FixtureSetup() {
    g.GetLogger().Info().Msg("in FixtureSetup...")
}

// 所有Test方法执行后执行
func (g *MyFixture) FixtureTeardown() {
    g.GetLogger().Info().Msg("in FixtureTearDown...")
}

// 每一个Test方法执行前执行
func (g *MyFixture) Setup() {
    g.WithLogger(c.T()).Info().Msg("in test setup...")
}

// 每一个Test方法执行后执行
func (g *MyFixture) Teardown() {
    g.GetLogger().Info().Msg("in test teardown...")
}

// 真正的测试方法A
func (g *MyFixture) TestA() {
    g.GetLogger().Description("这是TestA")
    g.GetLogger().Info().Msg("hello TestA...")
}

// 真正的测试方法B
func (g *MyFixture) TestB() {
    g.GetLogger().Description("这是TestB")
    g.GetLogger().Info().Msg("hello TestB...")
}

```

必须调用Msg或Msgf才能输出！

### 并行执行

默认情况下，所有fixture的方法都会并行运行，因为它们应该是独立的，但如果由于某种原因有需要按顺序运行fixture，可以向`Run()`
方法传入参数`gunit.Options.SequentialTestCases()`，例如在下面的例子中，这样fixture中的Test*方法将会按照ASCII码顺序执行。

```go
func TestExampleFixture(t *testing.T) {
    gunit.Run(new(ExampleFixture), t, gunit.Options.SequentialTestCases())
}
```

[Examples](https://github.com/bugVanisher/gunit/tree/master/examples)

----------------------------------------------------------------------------

对于JetBrains IDE的用户，以下是可以使用的LiveTemplate，用于生成新fixture的脚手架代码:

- Abbreviation: `fixture`
- Description: `生成 gunit Fixture 脚手架代码`
- Template Text:

```go
func Test$NAME$(t *testing.T) {
    gunit.Run(new($NAME$), t)
}

type $NAME$ struct {
    *gunit.Fixture
}

func (this *$NAME$) Setup() {
}

func (this *$NAME$) FixtureSetup() {
}

func (this *$NAME$) FixtureTeardown() {
}

func (this *$NAME$) Test$END$() {
}


```

----------------------------------------------------------------------------

## 执行&报告

#### 单独执行

`go test ./testcases/... -v`

#### 执行并生成报告

首次生成报告，先安装报告生成工具

`go get github.com/bugVanisher/gunit-test-report`

然后执行

`go test ./testcases/... -json | gunit-test-report`

将会在当前目录生成test_report.html报告。
