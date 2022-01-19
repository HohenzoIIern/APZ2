package main

import (
	"APZ2"
	"flag"
	"fmt"
	"io"
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
			os.Stderr.WriteString("Error reading file!")
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

	handler := &APZ2.ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	if err != nil {
		fmt.Println(err.Error())
	}
}
