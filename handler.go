package APZ2

import (
	"bufio"
	"io"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (c *ComputeHandler) Compute() error {
	fileScanner := bufio.NewScanner(c.Input)

	for fileScanner.Scan() {
		input := fileScanner.Text()
		res, err := PrefixToInfix(input)
		if err != nil {
			return err
		}
		c.Output.Write([]byte(res))
	}
	return nil
}
