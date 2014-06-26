package stack_test

import (
	"fmt"
	"gobook/chapter1/stacker/stack"
	"testing"
)

var myStack stack.Stack
var pStack *stack.Stack = new(stack.Stack)

func init() {
	myStack.Push("i am string")
	myStack.Push(1)
	myStack.Push(1.119)
	myStack.Push([]string{"pin", "white"})
	pStack.Push("dd")
	pStack.Push(3)
	pStack.Push(-5)
	pStack.Push(2.24)
	pStack.Push([]string{"black", "green"})
}

func TestLen(t *testing.T) {
	fmt.Printf("len(myStack) = %d\n", myStack.Len())
}

func TestLen2(t *testing.T) {
	fmt.Printf("len(pStack) = %d\n", pStack.Len())
}
