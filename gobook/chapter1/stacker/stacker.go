package main

import (
	"fmt"
	"gobook/chapter1/stacker/stack"
)

var pStack *stack.Stack = new(stack.Stack)

func init() {
	pStack.Push("dd")
	pStack.Push(3)
	pStack.Push(-5)
	pStack.Push(2.24)
	pStack.Push([]string{"black", "green"})
}

func main() {
	for {
		item, err := pStack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}
