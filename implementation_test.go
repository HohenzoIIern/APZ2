package APZ2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrefixToInfix(t *testing.T) {
	res, err := PrefixToInfix("+ * 22 b ^ c d")
	if assert.Nil(t, err) {
		assert.Equal(t, "22 * b + c ^ d", res)
	}
}

func TestPrefixToInfixExample(t *testing.T) {
	res, err := PrefixToInfix("+ - - - 11 a 3 4 * 12 a")
	assert.Equal(t, nil, err)
	assert.Equal(t, "11 - a - 3 - 4 + 12 * a", res)
}

func TestPostfixToPrefixAllOperators(t *testing.T) {
	res, err := PrefixToInfix("+ - 7 11 / 54 * 6 3")
	assert.Nil(t, err)
	assert.Equal(t, "7 - 11 + 54 / 6 * 3", res)
}

func TestPostfixToPrefixEmptyInput(t *testing.T) {
	res, err := PrefixToInfix("")
	assert.Equal(t, "", res)
	assert.NotNil(t, err)
}
