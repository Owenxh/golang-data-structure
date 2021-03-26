// Link: https://leetcode-cn.com/problems/add-two-numbers/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
		}
		if l2 != nil {
			sum += l2.Val
		}

		node := &ListNode{Val: sum % 10}
		if head == nil {
			head = node
		} else {
			tail.Next = node
		}
		tail = node
		carry = sum / 10

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}

func main() {
	// 2->4->3  5->6->4
	l1 := &ListNode{2, &ListNode{4, &ListNode{Val: 3}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{Val: 4}}}
	res := addTwoNumbers(l1, l2)
	for ; res != nil; res = res.Next {
		fmt.Print(res.Val)
		if res.Next != nil {
			fmt.Print("->")
		}
	}
}
