package main

import (
	"fmt"
	"strconv"
)

type Stack struct {
	elements []int
	size     int
}

func New() *Stack {
	return &Stack{make([]int, 10, 10), 0}
}

func (s *Stack) resize(capacity int) {
	arr := make([]int, capacity, capacity)
	for i := 0; i < s.size; i++ {
		arr[i] = s.elements[i]
	}

	s.elements = arr
}

func (s *Stack) Push(e int) {
	s.add(s.size, e)
}

func (s *Stack) add(index int, e int) {
	if index < 0 || index > s.size {
		panic("Add index mustn't less than 0 or large than Array's size")
	}

	if s.size >= len(s.elements) {
		s.resize(len(s.elements) * 2)
	}

	for i := s.size - 1; i >= index; i-- {
		s.elements[i+1] = s.elements[i]
	}

	s.elements[index] = e
	s.size++
}

func (s *Stack) Pop() int {
	return s.remove(s.size - 1)
}

func (s *Stack) remove(index int) int {
	if index < 0 || index >= s.size {
		panic("Remove index must >= 0 and < Array's size")
	}

	res := s.elements[index]

	for i := index + 1; i < s.size; i++ {
		s.elements[i-1] = s.elements[i]
	}
	s.elements[s.size-1] = 0
	s.size--

	if s.size > 0 && s.size <= len(s.elements)/4 {
		s.resize(len(s.elements) / 2)
	}

	return res
}

func (s *Stack) Peek() int {
	return s.Get(s.size - 1)
}

func (s *Stack) Get(index int) int {
	if index < 0 || index >= s.size {
		panic("Find index must >= 0 and < Array's size")
	}

	return s.elements[index]
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func calPoints(ops []string) int {
	stack := New()
	for _, v := range ops {
		if v == "C" {
			stack.Pop()
		} else if v == "D" {
			s1 := stack.Peek()
			stack.Push(s1 * 2)
		} else if v == "+" {
			s1 := stack.Pop()
			s2 := stack.Peek()
			stack.Push(s1)
			stack.Push(s1 + s2)
		} else {
			s, _ := strconv.Atoi(v)
			stack.Push(s)
		}
	}
	total := 0
	for i := 0; i < stack.size; i++ {
		total += stack.Get(i)
	}
	return total
}

func main() {
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))
	fmt.Println(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}))
}
