---
name: migrate-gunit-v2
description: Migrate Go test code from gunit v1 (github.com/smarty/gunit) to gunit v2 (github.com/smarty/gunit/v2). Use when a caller asks to upgrade, migrate, or convert tests, fixtures, or assertions from gunit v1 (or the legacy smarty/assertions module) to the v2 module.
---

# Migrating from gunit v1 to gunit v2

This skill guides the migration of Go test code from `github.com/smarty/gunit`
(v1) to `github.com/smarty/gunit/v2`. The two modules are similar in spirit;
most migrations are mechanical. Work one package at a time and run the tests
after each package.

## Step 1: Module wiring

```
go get github.com/smarty/gunit/v2
```

Update imports in every `_test.go` file:

| v1 import                                    | v2 import                                  |
|----------------------------------------------|--------------------------------------------|
| `github.com/smarty/gunit`                    | `github.com/smarty/gunit/v2`               |
| `github.com/smarty/gunit/assert`             | `github.com/smarty/gunit/v2/assert`        |
| `github.com/smarty/gunit/assert/should`      | `github.com/smarty/gunit/v2/assert/should` |
| `github.com/smarty/gunit/assert/better`      | `github.com/smarty/gunit/v2/assert/better` |
| `github.com/smarty/assertions/should`        | `github.com/smarty/gunit/v2/assert/should` |
| `github.com/smartystreets/gunit`             | `github.com/smarty/gunit/v2`               |
| `github.com/smartystreets/assertions/should` | `github.com/smarty/gunit/v2/assert/should` |

The import identifier stays `gunit` (the v2 package is named `gunit`), so
`gunit.Run`, `gunit.Fixture`, and `should.*` references keep compiling once
imports are updated. When the whole module is migrated, run `go mod tidy` to
drop the v1 requirement.

## Step 2: Runner and options

The basic pattern is unchanged:

```go
func TestMyFixture(t *testing.T) {
	gunit.Run(new(MyFixture), t)
}

type MyFixture struct {
	*gunit.Fixture
}
```

Defaults match: v2's default (`Options.UnitTests()`) gives parallel execution
with a fresh fixture instance per test case, which is what v1 did by default.

| v1                                    | v2                                                       |
|---------------------------------------|----------------------------------------------------------|
| `gunit.Run(f, t)`                     | `gunit.Run(f, t)` (unchanged)                            |
| `gunit.RunSequential(f, t)`           | `gunit.Run(f, t, gunit.Options.IntegrationTests())`      |
| `gunit.Options.SkipAll()`             | `gunit.Options.SkipAll()` (unchanged)                    |
| `gunit.Options.AllSequential()`       | `gunit.Options.IntegrationTests()`                       |
| `gunit.Options.SequentialFixture()`   | `gunit.Options.IntegrationTests()` (closest; see caveat) |
| `gunit.Options.SequentialTestCases()` | `gunit.Options.IntegrationTests()` (closest; see caveat) |
| `gunit.Options.LongRunning()`         | `gunit.Options.LongRunning()` (unchanged)                |

Caveats:

- **Shared vs fresh fixtures**: v1 always instantiated a fresh fixture per
  test case, even in sequential modes. v2's `IntegrationTests()` (and
  `SharedFixture()`) reuses ONE fixture instance across all test cases. If
  the fixture carries per-test state, reset it in `Setup()` — which still
  runs before each test — or flag the behavioral difference to the caller.
- **LongRunning is included in IntegrationTests**: v2's `IntegrationTests()`
  composite applies `LongRunning()`, so any suite mapped to it (including
  migrated `RunSequential`/`AllSequential` suites) will be skipped entirely
  under `go test -short`. That is usually desirable for integration suites;
  if the caller wants shared fixtures WITHOUT the `-short` skip, use
  `gunit.Options.SharedFixture()` instead.

New in v2 (offer these where useful, don't force them):

- `SetupSuite()` / `TeardownSuite()` methods run once per fixture, before and
  after all test cases — useful with `IntegrationTests()` for expensive
  resources.
- Options `FreshFixture()`, `SharedFixture()`, `ParallelFixture()`,
  `ParallelTests()`, `UnitTests()` allow finer control.

Unchanged conventions: `Setup()`/`Teardown()` per-test hooks, `Test*` method
discovery, `SkipTest*` and `FocusTest*` prefixes, and
`this.Run(name, func(fixture *gunit.Fixture))` for table-driven subtests.

**Dropped prefixes**: v1 also recognized `LongTest*`, `SkipLongTest*`, and
`FocusLongTest*` method prefixes (skipped under `go test -short`). v2 does
not — a `LongTest*` method would be silently ignored by the v2 runner. Rename
to `Test*` (or `SkipTest*`/`FocusTest*`). When the whole fixture is
long-running, pass `gunit.Options.LongRunning()` (or `IntegrationTests()`,
which includes it) to `gunit.Run`; when only some methods are, guard those
methods individually with `if testing.Short() { this.Skip(...) }`.

## Step 3: Fixture method changes

| v1 usage                              | v2 replacement                                      |
|----------------------------------------|------------------------------------------------------|
| `this.T()`                             | Fixture embeds `TestingT`: call `this.Log(...)`, `this.Name()`, `this.Cleanup(...)`, etc. directly, or use `this.TestingT` |
| `ok := this.So(...)` (returns bool)    | `this.So(...)` returns nothing; if the bool guarded an early return or short-circuit, use the `better` package instead (see below) |
| `this.Assert(cond, msg)`               | `this.So(cond, should.BeTrue)` or `if !cond { this.Error(msg) }` |
| `this.AssertEqual(expected, actual)`   | `this.So(actual, should.Equal, expected)` — NOTE: argument order flips |
| `this.AssertDeepEqual(expected, actual)` | `this.So(actual, should.Equal, expected)`          |
| `this.AssertSprintEqual(e, a)`         | `this.So(fmt.Sprint(a), should.Equal, fmt.Sprint(e))` |
| `this.AssertSprintfEqual(e, a, f)`     | `this.So(fmt.Sprintf(f, a), should.Equal, fmt.Sprintf(f, e))` |
| `this.Write(p)` (io.Writer)            | `this.Output().Write(p)`; for `log.SetOutput(fixture)` use `log.SetOutput(this.Output())` |
| `this.Print/Printf/Println(...)`       | Unchanged                                            |
| `this.Error/Errorf(...)`               | Unchanged (now provided via embedded `TestingT`)     |
| `this.Failed()`, `this.Name()`         | Unchanged                                            |

### Early returns guarded by `So`'s bool result

In v1, `So` returned `false` on failure, and callers often used that to bail
out of a test before dereferencing something unsafe:

```go
// v1
if !this.So(err, should.BeNil) {
	return
}
this.So(result.Value, should.Equal, 42)
```

In v2, replace the guard with the `github.com/smarty/gunit/v2/assert/better`
package. It provides the same assertions as `should` (including a
`better.NOT` namespace), but a failure is fatal — `Fixture.So` calls
`t.Fatal`, halting the test immediately, so no early return is needed:

```go
// v2
this.So(err, better.BeNil)
this.So(result.Value, should.Equal, 42)
```

Use `better` only where the remainder of the test cannot proceed meaningfully
after the failure; keep `should` for ordinary assertions.

## Step 4: Assertion mapping

v2 replaces `Not*` assertion names with a `should.NOT` namespace and drops or
merges several rarely-used assertions. Full mapping:

| v1 (`should.`)          | v2 (`should.`)                                     |
|--------------------------|-----------------------------------------------------|
| `AlmostEqual`            | Not implemented — see "Missing assertions"          |
| `BeBetween`              | Not implemented — compose `BeGreaterThan` + `BeLessThan` |
| `BeBetweenOrEqual`       | Not implemented — compose `*OrEqualTo` variants     |
| `BeBlank`                | `BeEmpty` (or `Equal, ""`)                          |
| `BeChronological`        | `BeChronological`                                   |
| `BeEmpty`                | `BeEmpty`                                           |
| `BeFalse`                | `BeFalse`                                           |
| `BeGreaterThan`          | `BeGreaterThan`                                     |
| `BeGreaterThanOrEqualTo` | `BeGreaterThanOrEqualTo`                            |
| `BeIn`                   | `BeIn`                                              |
| `BeLessThan`             | `BeLessThan`                                        |
| `BeLessThanOrEqualTo`    | `BeLessThanOrEqualTo`                               |
| `BeNil`                  | `BeNil`                                             |
| `BeTrue`                 | `BeTrue`                                            |
| `BeZeroValue`            | `BeZeroValue`                                       |
| `Contain`                | `Contain`                                           |
| `ContainKey`             | `Contain` (v2 `Contain` handles map keys)           |
| `ContainSubstring`       | `Contain` (v2 `Contain` handles substrings/runes)   |
| `EndWith`                | `EndWith`                                           |
| `Equal`                  | `Equal`                                             |
| `HappenAfter`            | `HappenAfter`                                       |
| `HappenBefore`           | `HappenBefore`                                      |
| `HappenBetween`          | Not implemented — compose `HappenAfter` + `HappenBefore` |
| `HappenOnOrAfter`        | Not implemented — compose `HappenOn` / `HappenAfter` |
| `HappenOnOrBefore`       | Not implemented — compose `HappenOn` / `HappenBefore` |
| `HappenOnOrBetween`      | Not implemented — compose from the above            |
| `HappenWithin`           | `HappenWithin`                                      |
| `HaveLength`             | `HaveLength`                                        |
| `HaveSameTypeAs`         | Not implemented                                     |
| `NotAlmostEqual`         | Not implemented                                     |
| `NotBeBetween`           | Not implemented                                     |
| `NotBeBetweenOrEqual`    | Not implemented                                     |
| `NotBeBlank`             | `NOT.BeEmpty`                                       |
| `NotBeChronological`     | `NOT.BeChronological`                               |
| `NotBeEmpty`             | `NOT.BeEmpty`                                       |
| `NotBeIn`                | `NOT.BeIn`                                          |
| `NotBeNil`               | `NOT.BeNil`                                         |
| `NotBeZeroValue`         | `NOT.BeZeroValue`                                   |
| `NotContain`             | `NOT.Contain`                                       |
| `NotContainKey`          | `NOT.Contain`                                       |
| `NotContainSubstring`    | `NOT.Contain`                                       |
| `NotEndWith`             | Not implemented                                     |
| `NotEqual`               | `NOT.Equal`                                         |
| `NotHappenOnOrBetween`   | Not implemented                                     |
| `NotHappenWithin`        | Not implemented                                     |
| `NotHaveSameTypeAs`      | Not implemented                                     |
| `NotPanic`               | `NOT.Panic`                                         |
| `NotPanicWith`           | Not implemented                                     |
| `NotStartWith`           | Not implemented                                     |
| `Panic`                  | `Panic`                                             |
| `PanicWith`              | Not implemented                                     |
| `StartWith`              | `StartWith`                                         |
| `Wrap`                   | `WrapError`                                         |

Also available in v2 but new (no v1 counterpart): `should.NOT.BeGreaterThan`,
`should.NOT.BeGreaterThanOrEqualTo`, `should.NOT.BeLessThan`,
`should.NOT.BeLessThanOrEqualTo`, `should.HappenOn`, `should.NOT.HappenOn`.

Example rewrite:

```go
// v1
this.So(m, should.ContainKey, "id")
this.So(s, should.NotContainSubstring, "secret")
this.So(err, should.Wrap, io.EOF)

// v2
this.So(m, should.Contain, "id")
this.So(s, should.NOT.Contain, "secret")
this.So(err, should.WrapError, io.EOF)
```

### Missing assertions

When a v1 assertion has no v2 equivalent, prefer (in order):

1. Compose from existing v2 assertions when semantics allow (e.g. two `So`
   calls for `BeBetween`).
2. Write a small custom assertion (see next section).
3. Fall back to plain checks: `if cond { this.Errorf(...) }`.

Flag every such rewrite to the caller — do not silently weaken an assertion.

## Step 5: Custom assertions

The assertion contract changed from string-based to error-based:

```go
// v1: assert.Func — return "" on success, message on failure
func BeShiny(actual any, expected ...any) string {
	if isShiny(actual) {
		return ""
	}
	return fmt.Sprintf("wanted shiny, got: %v", actual)
}

// v2: assert.Assertion — return nil on success, error on failure
func BeShiny(actual any, expected ...any) error {
	if isShiny(actual) {
		return nil
	}
	return fmt.Errorf("%w: wanted shiny, got: %v", should.ErrAssertionFailure, actual)
}
```

A v1 result prefixed with `"<<<FATAL>>>\n"` halted the test; in v2, wrap
`should.ErrFatalAssertionFailure` instead — `assert.So` calls `t.Fatal` when
`errors.Is(err, should.ErrFatalAssertionFailure)`.

## Step 6: Verify

Run the migrated tests (prefer `make test` if the project has it, otherwise
`go test ./...`). All tests must pass before moving to the next package.
Watch for:

- Tests that relied on `So`'s boolean return to short-circuit — these must
  use `better.*` assertions (fatal) so the test still halts before unsafe
  code.
- Argument-order bugs from `AssertEqual(expected, actual)` →
  `So(actual, should.Equal, expected)`.
- State leakage in fixtures migrated to `IntegrationTests()`/`SharedFixture()`.
- Behavioral drift where a "Not implemented" assertion was rewritten.
