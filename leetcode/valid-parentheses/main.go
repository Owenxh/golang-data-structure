// https://leetcode-cn.com/problems/valid-parentheses/

package main

import (
	"fmt"
)

type Stack struct {
	elements []string
	size     int
}

func New() *Stack {
	return &Stack{make([]string, 10, 10), 0}
}

func (s *Stack) resize(capacity int) {
	arr := make([]string, capacity, capacity)
	for i := 0; i < s.size; i++ {
		arr[i] = s.elements[i]
	}

	s.elements = arr
}

func (s *Stack) Push(e string) {
	s.add(s.size, e)
}

func (s *Stack) add(index int, e string) {
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

func (s *Stack) Pop() string {
	return s.remove(s.size - 1)
}

func (s *Stack) remove(index int) string {
	if index < 0 || index >= s.size {
		panic("Remove index must >= 0 and < Array's size")
	}

	res := s.elements[index]

	for i := index + 1; i < s.size; i++ {
		s.elements[i-1] = s.elements[i]
	}
	s.elements[s.size-1] = ""
	s.size--

	if s.size > 0 && s.size <= len(s.elements)/4 {
		s.resize(len(s.elements) / 2)
	}

	return res
}

func (s *Stack) Peek() string {
	return s.Get(s.size - 1)
}

func (s *Stack) Get(index int) string {
	if index < 0 || index >= s.size {
		panic("Find index must >= 0 and < Array's size")
	}

	return s.elements[index]
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func isValid(s string) bool {
	st := New()
	for _, v := range s {
		item := string(v)
		if item == "(" || item == "[" || item == "{" {
			st.Push(item)
		} else {
			if st.IsEmpty() {
				return false
			}
			if item == ")" && st.Pop() != "(" {
				return false
			}

			if item == "]" && st.Pop() != "[" {
				return false
			}

			if item == "}" && st.Pop() != "{" {
				return false
			}
		}
	}

	return st.IsEmpty()
}

func main() {
	fmt.Println(isValid("String"))
	fmt.Println(isValid("([{}])}])"))
	fmt.Println(isValid("([{}])()"))
	fmt.Println(isValid("([{]}])"))
	fmt.Println(isValid("]"))
}
