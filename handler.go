package APZ2

import (
	"io"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (c *ComputeHandler) Compute() error {
	expr := ""
	partial := make([]byte, 8)
	for {
		count, err := (c.Input).Read(partial)
		if err != nil {
			break
		}
		expr += string(partial[:count])
	}
	res, err := PrefixToInfix(expr)
	if err != nil {
		return err
	}
	c.Output.Write([]byte(res))
	return nil
}
