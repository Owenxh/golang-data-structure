package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	Head *ListNode
	Tail *ListNode
}

func (l *List) add(n *ListNode) {
	if l.Head == nil {
		l.Head, l.Tail = n, n
	} else {
		l.Tail.Next = n
		l.Tail = l.Tail.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &List{}
	for l1 != nil || l2 != nil {
		if (l1 == nil && l2 != nil) || (l2 != nil && l1.Val >= l2.Val) {
			list.add(l2)
			l2 = l2.Next
		} else {
			list.add(l1)
			l1 = l1.Next
		}
	}

	return list.Head
}
