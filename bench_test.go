package APZ2

import (
	"fmt"
	"testing"
)

var cntRes string
var err error

func BenchmarkPrefixToInfix(b *testing.B) {
	const baseLen = 11 // number of operators in input
	inputLength := baseLen
	input := "+ 1 + 2 + 3 + 4 + 5 6"

	for i := 0; i < 20; i++ {
		input = "*" + input + input
		inputLength = inputLength*2 + 1
		b.Run(fmt.Sprintf("input length = %d", inputLength), func(b *testing.B) {
			cntRes, err = PrefixToInfix(input)
		})

	}
}
