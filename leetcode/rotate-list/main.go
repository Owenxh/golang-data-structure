// link: https://leetcode-cn.com/problems/rotate-list/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var values []int
	c := l
	for c != nil {
		values = append(values, c.Val)
		c = c.Next
	}
	return fmt.Sprint(values)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
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
	if head == nil {
		return nil
	}
	size := 0
	cur, last := head, head
	for cur != nil {
		size++
		last = cur
		cur = cur.Next
	}

	if k%size == 0 {
		return head
	}

	last.Next = head
	retPrev := head
	h := size - (k % size)
	for i := 0; i < h-1; i++ {
		retPrev = retPrev.Next
	}

	retHead := retPrev.Next
	retPrev.Next = nil

	return retHead
}

func main() {
	head := &ListNode{0, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}}}}}
	//for i := 0; i < 16; i++ {
	//	fmt.Println(rotateRight(head, i))
	//}
	fmt.Println(rotateRight(head, 1))
}
