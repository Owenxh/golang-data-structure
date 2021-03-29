// link: https://leetcode-cn.com/problems/merge-k-sorted-lists
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

func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

// 合并 lists[l, r] 个有序链表并返回新的链表
func merge(lists []*ListNode, l, r int) *ListNode {
	if l > r {
		return nil
	}
	if l == r {
		return lists[l]
	}
	//if r-l == 1 {
	//	return merge0(lists[l], lists[r])
	//}
	//mid := l + (r-l+1)/2
	mid := l + (r-l)/2
	l1 := merge(lists, l, mid)
	l2 := merge(lists, mid+1, r)
	return merge0(l1, l2)
}

func merge0(l1, l2 *ListNode) *ListNode {
	h := &ListNode{}
	tail := h
	for l1 != nil || l2 != nil {
		if l2 == nil || (l1 != nil && l1.Val <= l2.Val) {
			next := l1.Next
			l1.Next, tail.Next = nil, l1
			l1 = next
		} else {
			next := l2.Next
			l2.Next, tail.Next = nil, l2
			l2 = next
		}
		tail = tail.Next
	}
	return h.Next
}

func main() {
	lists := []*ListNode{
		{1, &ListNode{4, &ListNode{5, nil}}},
		{1, &ListNode{3, &ListNode{4, nil}}},
		{2, &ListNode{6, nil}},
	}
	fmt.Println(mergeKLists(lists))
}
