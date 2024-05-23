package should_test

import (
	"testing"

	"github.com/smarty/gunit/v2/should"
)

func TestShouldHaveLength(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.HaveLength, " EXPECTED", "EXTRA")
	assert.KindMismatch(true, should.HaveLength, 0)

	assert.KindMismatch(42, should.HaveLength, 0)
	assert.KindMismatch("", should.HaveLength, "")

	assert.Pass([]string(nil), should.HaveLength, 0)
	assert.Pass([]string{}, should.HaveLength, 0)
	assert.Pass([]string{""}, should.HaveLength, 1)
	assert.Fail([]string{""}, should.HaveLength, 2)

	assert.Pass([0]string{}, should.HaveLength, 0) // The only possible empty array!
	assert.Fail([1]string{}, should.HaveLength, 2)

	assert.Pass(chan string(nil), should.HaveLength, 0)
	assert.Fail(nonEmptyChannel(), should.HaveLength, 2)

	assert.Pass(map[string]string{"": ""}, should.HaveLength, 1)
	assert.Fail(map[string]string{"": ""}, should.HaveLength, 2)

	assert.Pass("", should.HaveLength, 0)
	assert.Pass("123", should.HaveLength, 3)
	assert.Fail("123", should.HaveLength, 4)
}
