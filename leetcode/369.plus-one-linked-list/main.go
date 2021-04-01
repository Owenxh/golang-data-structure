// link: https://leetcode-cn.com/problems/plus-one-linked-list/

package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var buf bytes.Buffer
	for c := l; c != nil; c = c.Next {
		buf.WriteString(strconv.Itoa(c.Val))
		buf.WriteString("->")
	}
	buf.WriteString("NULL")
	return buf.String()
}

//func (l *ListNode) String() string {
//	var res string
//	for c := l; c != nil; c = c.Next {
//		res += strconv.Itoa(c.val) + "->"
//	}
//	res += "NULL"
//	return res
//}

func plusOne(head *ListNode) *ListNode {
	l := reverse(head)
	c, carry := l, 1
	for c != nil {
		val := c.Val + carry
		c.Val = val % 10
		if carry = val / 10; carry == 0 {
			break
		}
		c = c.Next
	}
	res := reverse(l)
	if carry > 0 {
		return &ListNode{carry, res}
	}
	return res
}

func reverse(head *ListNode) *ListNode {
	dummy := &ListNode{}
	for head != nil {
		next := head.Next
		head.Next, dummy.Next = dummy.Next, head
		head = next
	}
	return dummy.Next
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	fmt.Println(head)
	fmt.Println(plusOne(head))

	head = &ListNode{9, &ListNode{9, &ListNode{9, nil}}}
	fmt.Println(head)
	fmt.Println(plusOne(head))
}
