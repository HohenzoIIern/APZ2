package lab2

import (
	"fmt"
	"strings"
)

type StackNode struct {
	data string
	next *StackNode
}

func getStackNode(data string, top *StackNode) *StackNode {
	return &StackNode{
		data,
		top,
	}
}

type MyStack struct {
	top   *StackNode
	count int
}

func getMyStack() *MyStack {
	return &MyStack{
		nil,
		0,
	}
}

func (this MyStack) size() int {
	return this.count
}

func (this MyStack) isEmpty() bool {
	if this.size() > 0 {
		return false
	} else {
		return true
	}
}

func (this *MyStack) push(data string) {
	this.top = getStackNode(data, this.top)
	this.count++
}

func (this *MyStack) pop() string {
	var temp string = ""
	if this.isEmpty() == false {
		temp = this.top.data
		this.top = this.top.next
		this.count--
	}
	return temp
}

func isOperator(text byte) bool {
	if text == '+' || text == '-' || text == '*' ||
		text == '/' || text == '^' || text == '%' {
		return true
	}
	return false
}

func isOperands(text byte) bool {
	if (text >= '0' && text <= '9') ||
		(text >= 'a' && text <= 'z') ||
		(text >= 'A' && text <= 'Z') {
		return true
	}
	return false
}

func PrefixToInfix(prefix string) (string, error) {
	if len(prefix) <= 2 {
		return "", fmt.Errorf("incorrect expression")
	}
	var buff = strings.Split(prefix, " ")
	var size int = len(buff)
	var s *MyStack = getMyStack()
	var auxiliary string = ""
	var op1 string = ""
	var op2 string = ""
	var isValid bool = true

	for i := size - 1; i >= 0; i-- {
		if isOperator(buff[i][0]) {
			if s.size() > 1 {
				op1 = s.pop()
				op2 = s.pop()
				auxiliary = op1 + " " + buff[i] + " " + op2
				s.push(auxiliary)
			} else {
				isValid = false
			}
		} else {
			auxiliary = buff[i]
			s.push(auxiliary)
		}
	}
	if isValid == false {
		fmt.Println("Invalid Prefix : ", prefix)
	}

	return s.pop(), nil
}
