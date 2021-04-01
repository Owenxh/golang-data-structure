// link: https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/

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
 *     val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead, curr := &ListNode{}, head
	last, first := 0, true
	for curr != nil {
		val := curr.Val
		next := curr.Next

		if first || last != curr.Val {
			curr.Next, dummyHead.Next = dummyHead.Next, curr
		} else if last == curr.Val && dummyHead.Next != nil && dummyHead.Next.Val == curr.Val {
			dummyHead.Next = dummyHead.Next.Next
		}
		last = val
		curr = next
		if first {
			first = false
		}
	}

	return reverse(dummyHead.Next)
}

func reverse(head *ListNode) *ListNode {
	dummyHead, curr := &ListNode{}, head
	for curr != nil {
		next := curr.Next
		curr.Next, dummyHead.Next = dummyHead.Next, curr
		curr = next
	}

	return dummyHead.Next
}

func main() {
	head := &ListNode{0, &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}}}}
	fmt.Println(deleteDuplicates(head))

	head = &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{2, nil}}}}
	fmt.Println(deleteDuplicates(head))
}
