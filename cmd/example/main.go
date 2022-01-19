package main

import (
	"flag"
	"fmt"
	"io"
	"lab2"
	"os"
	"strings"
)

var (
	inputExpression = flag.String("e", "", "Expression")
	inputFile       = flag.String("f", "", "Input file")
	outputFile      = flag.String("o", "", "Output file")
)

func main() {
	flag.Parse()
	var input io.Reader = nil
	var output = os.Stdout

	if *inputExpression != "" {
		input = strings.NewReader(*inputExpression)
	}

	if *inputFile != "" {
		f, err := os.Open(*inputFile)
		if f != nil {
			f, err = os.Open("input.txt")
		}
		if err != nil {
			fmt.Print("Error\n")
		}
		input = f
	}

	if *outputFile != "" {
		f, err := os.Create(*outputFile)
		if err != nil {
			fmt.Print("Error\n")
		}
		output = f
	}

	if input == nil {
		fmt.Print("no input\n")
		return
	}
	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}
	err := handler.Compute()
	if err != nil {
		fmt.Println(err.Error())
	}
}
