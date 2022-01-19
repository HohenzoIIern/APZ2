package lab2

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type OutputSim struct {
	called bool
}

func (o *OutputSim) Write(p []byte) (n int, err error) {
	o.called = true
	return 0, nil
}

func TestComputerHandleCorrect(t *testing.T) {
	testExpression := "+ * 22 b ^ c d"
	input := strings.NewReader(testExpression)
	output := OutputSim{}
	handler := ComputeHandler{
		Input:  input,
		Output: &output,
	}
	err := handler.Compute()
	assert.Nil(t, err)
	assert.Equal(t, true, output.called)
}

func TestComputerHandleIncorrect(t *testing.T) {
	test := ""
	input := strings.NewReader(test)
	output := OutputSim{}
	handler := ComputeHandler{
		Input:  input,
		Output: &output,
	}
	err := handler.Compute()
	assert.NotNil(t, err)
}
