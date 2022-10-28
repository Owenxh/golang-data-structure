package stack

import (
	"fmt"

	"io.vava.datastructure/array"
)

type ArrayStack[E any] struct {
	array *array.Array[E]
}

func New[E any]() *ArrayStack[E] {
	return &ArrayStack[E]{array.New[E]()}
}

func NewWithCapacity[E any](capacity int) *ArrayStack[E] {
	return &ArrayStack[E]{array.NewWithCapacity[E](capacity)}
}

func (s *ArrayStack[E]) Push(e E) {
	s.array.AddLast(e)
}

func (s *ArrayStack[E]) Pop() E {
	return s.array.RemoveLast()
}

func (s *ArrayStack[E]) Peek() E {
	return s.array.GetLast()
}

func (s *ArrayStack[E]) Size() int {
	return s.array.Size()
}

func (s *ArrayStack[E]) IsEmpty() bool {
	return s.array.IsEmpty()
}

func (s *ArrayStack[E]) String() string {
	return fmt.Sprintf("Stack:{Bottom %v Top, size:%d}", s.array.GetAll(), s.array.Size())
}
