package stack

import (
	"errors"
)

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack Stack) Cap() int {
	return cap(stack)
}

func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}

func (stack *Stack) Push(x interface{}) {
	*stack = append(*stack, x)
}

func (stack Stack) Top() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("Can't Top from an empty stack!")
	}
	return stack[len(stack)-1], nil
}

func (stack *Stack) Pop() (interface{}, error) {
	theFinalStack := *stack
	if len(theFinalStack) == 0 {
		return nil, errors.New("Can't Pop from an empty stack!")
	}
	x := theFinalStack[len(theFinalStack)-1]
	*stack = theFinalStack[:len(theFinalStack)-1]
	return x, nil
}
