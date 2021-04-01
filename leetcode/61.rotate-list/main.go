// link: https://leetcode-cn.com/problems/rotate-list/
package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var r []int
	c := l
	for c != nil {
		r = append(r, c.Val)
		c = c.Next
	}
	return fmt.Sprint(r)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     val int
 *     Next *ListNode
 * }
 */
func rotateRightUsingTempSlice(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	var tmp []int
	size := 0
	cur := head
	for cur != nil {
		tmp = append(tmp, cur.Val)
		size++
		cur = cur.Next
	}

	if k%size == 0 {
		return head
	}

	dummy := &ListNode{}
	prev := dummy
	h := size - (k % size)
	for i := 0; i < size; i++ {
		prev.Next = &ListNode{tmp[h], nil}
		h = (h + 1) % size
		prev = prev.Next
	}
	return dummy.Next
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	size := 1
	tail := head
	for tail.Next != nil {
		size++
		tail = tail.Next
	}

	if k%size == 0 {
		return head
	}

	tail.Next = head
	temp := head
	for i, h := 1, size-(k%size); i < h; i++ {
		temp = temp.Next
	}

	res := temp.Next
	temp.Next = nil
	return res
}

func main() {
	for i := 0; i < 16; i++ {
		head := &ListNode{0, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}}}}}
		res1 := rotateRightUsingTempSlice(head, i)
		res2 := rotateRight(head, i)
		if !reflect.DeepEqual(res1, res2) {
			panic("The two algorithm implementation are not same")
		}
		fmt.Printf("rotate %2d => %v\n", i, res1)
	}
}
