# Introduction

If you're brand new to gunit please see the [v1 module introduction](../README_v1_intro.md).

The v2 release is very similar, in spirit, to v1. Here are some interesting bits of context regarding this release:

1. As of v1.6.0 gunit has provided its own assertions package. Having assertions in a separate module is a reflection of a time when [GoConvey](https://github.com/smartystreets/goconvey) and gunit were both in use at Smarty. We've fully moved away from GoConvey for some time now so we decided to inline the assertions.
2. While I had fully written off a v2 release for some time, our recent adoption of AI tools has made certain migrations and other chores much more palatable. Upgrading to v2 has turned out to be pretty trivial.
3. The test runner code of v2 is far simpler to understand and maintain, and the functional options are more intuitive (you basically just specify whether you are unit testing or integration testing).
4. v2 has better support for integration testing scenarios as it allows for an instance of a fixture to be shared between test cases (with the `gunit.Options.IntegrationTests()` option), and optionally defined SetupSuite() and TeardownSuite() methods that are invoked once for the fixture (before and after, respectively).
5. The assertions in the v2 module have been fully rewritten. The structure of their implementation is much easier to read and maintain, and they don't depend on the vendored [oglematchers](https://github.com/smarty/gunit/tree/master/assert/assertions/internal/oglematchers) package. I hope to remove other vendored packages in the future.
6. Some assertions were combined, for example, `should.Contain` and `should.ContainSubstring` are now just `should.Contain`. There are a few other examples of that.
7. Some assertions that we just never used were not implemented in v2.
8. Overall, I'm just much happier with the v2 code.

The v1 module remains in this repository, separate from, but alongside, this new v2 module. This was very intentional and should facilitate ongoing maintenance of both versions.

