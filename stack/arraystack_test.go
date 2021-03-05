package stack_test

import (
	"fmt"
	"testing"

	"io.vava.datastructure/stack"
)

func TestArrayStack_Push(t *testing.T) {
	stack := stack.NewWithCapacity(8)

	for i := 1; i <= 10; i++ {
		stack.Push(i)
		fmt.Println(stack)
	}

	fmt.Printf("Top element = %v\n", stack.Peek())
	fmt.Printf("Is empty = %v\n", stack.IsEmpty())

	for i := 1; i <= 10; i++ {
		stack.Pop()
		fmt.Println(stack)
	}

	fmt.Printf("Is empty = %v\n", stack.IsEmpty())
}
