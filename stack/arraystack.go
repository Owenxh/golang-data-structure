package stack

import (
	"fmt"

	"io.vava.datastructure/array"
)

type ArrayStack struct {
	array *array.Array
}

func New() *ArrayStack {
	return &ArrayStack{array.New()}
}

func NewWithCapacity(capacity int) *ArrayStack {
	return &ArrayStack{array.NewWithCapacity(capacity)}
}

func (s *ArrayStack) Push(e int) {
	s.array.AddLast(e)
}

func (s *ArrayStack) Pop() int {
	return s.array.RemoveLast()
}

func (s *ArrayStack) Peek() int {
	return s.array.GetLast()
}

func (s *ArrayStack) Size() int {
	return s.array.Size()
}

func (s *ArrayStack) IsEmpty() bool {
	return s.array.IsEmpty()
}

func (s *ArrayStack) String() string {
	return fmt.Sprintf("Stack:{Bottom %v Top, size:%d}", s.array.GetAll(), s.array.Size())
}
