// link: https://leetcode-cn.com/problems/revers e-nodes-in-k-group

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
	h := l
	for h != nil {
		buf.WriteString(strconv.Itoa(h.Val))
		buf.WriteString("->")
		h = h.Next
	}
	buf.WriteString("NULL")
	return buf.String()
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}
	dummyHead := &ListNode{}
	curr, tail, prev := head, dummyHead, dummyHead
	for curr != nil {
		for i := 1; i <= k && curr != nil; i++ {
			next := curr.Next
			curr.Next, prev.Next, curr, prev = nil, curr, next, curr
			if i == k {
				tail.Next, prev = reverse(tail.Next)
				tail = prev
			}
		}
	}
	return dummyHead.Next
}

func reverse(l *ListNode) (h, t *ListNode) {
	h = &ListNode{}
	t = l
	for l != nil {
		next := l.Next
		l.Next, h.Next, l = h.Next, l, next
	}
	return h.Next, t
}

func main() {
	h := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	fmt.Println(reverseKGroup(h, 2))
}
