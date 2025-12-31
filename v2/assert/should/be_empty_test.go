package should_test

import (
	"testing"

	"github.com/smarty/gunit/v2/assert/should"
)

func TestShouldBeEmpty(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.BeEmpty, "EXTRA")

	assert.KindMismatch(42, should.BeEmpty)

	assert.Pass([]string(nil), should.BeEmpty)
	assert.Pass(make([]string, 0, 0), should.BeEmpty)
	assert.Pass(make([]string, 0, 1), should.BeEmpty)
	assert.Fail([]string{""}, should.BeEmpty)

	assert.Pass([0]string{}, should.BeEmpty) // The only possible empty array!
	assert.Fail([1]string{}, should.BeEmpty)

	assert.Pass(chan string(nil), should.BeEmpty)
	assert.Pass(make(chan string), should.BeEmpty)
	assert.Pass(make(chan string, 1), should.BeEmpty)
	assert.Fail(nonEmptyChannel(), should.BeEmpty)

	assert.Pass(map[string]string(nil), should.BeEmpty)
	assert.Pass(make(map[string]string), should.BeEmpty)
	assert.Pass(make(map[string]string, 1), should.BeEmpty)
	assert.Fail(map[string]string{"": ""}, should.BeEmpty)

	assert.Pass("", should.BeEmpty)
	assert.Pass(*new(string), should.BeEmpty)
	assert.Fail(" ", should.BeEmpty)
}

func TestShouldNotBeEmpty(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.NOT.BeEmpty, "EXTRA")
	assert.KindMismatch(42, should.NOT.BeEmpty)

	assert.Fail([]string(nil), should.NOT.BeEmpty)
	assert.Fail(make([]string, 0, 0), should.NOT.BeEmpty)
	assert.Fail(make([]string, 0, 1), should.NOT.BeEmpty)
	assert.Pass([]string{""}, should.NOT.BeEmpty)

	assert.Fail([0]string{}, should.NOT.BeEmpty)
	assert.Pass([1]string{}, should.NOT.BeEmpty)

	assert.Fail(chan string(nil), should.NOT.BeEmpty)
	assert.Fail(make(chan string), should.NOT.BeEmpty)
	assert.Fail(make(chan string, 1), should.NOT.BeEmpty)
	assert.Pass(nonEmptyChannel(), should.NOT.BeEmpty)

	assert.Fail(map[string]string(nil), should.NOT.BeEmpty)
	assert.Fail(make(map[string]string), should.NOT.BeEmpty)
	assert.Fail(make(map[string]string, 1), should.NOT.BeEmpty)
	assert.Pass(map[string]string{"": ""}, should.NOT.BeEmpty)

	assert.Fail("", should.NOT.BeEmpty)
	assert.Fail(*new(string), should.NOT.BeEmpty)
	assert.Pass(" ", should.NOT.BeEmpty)
}

func nonEmptyChannel() chan string {
	c := make(chan string, 1)
	c <- ""
	return c
}
