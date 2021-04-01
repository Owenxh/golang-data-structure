// link: https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates2(head *ListNode) *ListNode {
	res := head
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = nil
		if prev != nil && prev.Val != head.Val {
			prev.Next = head
			prev = head
		} else if prev == nil {
			prev = head
		}
		head = next
	}
	return res
}

func deleteDuplicates(head *ListNode) *ListNode {
	deleteDuplicates0(head, nil)
	return head
}

func deleteDuplicates0(head, prev *ListNode) {
	if head == nil {
		return
	}

	next := head.Next
	head.Next = nil

	if prev == nil || prev.Val != head.Val {
		if prev != nil {
			prev.Next = head
		}
		prev = head
	}

	deleteDuplicates0(next, prev)
}

//func deleteDuplicates0(head, prev *ListNode) {
//	if head == nil {
//		return
//	}
//
//	next := head.Next
//	head.Next = nil
//	if prev != nil {
//		prev.Next = head
//	}
//	prev, head = head, next
//	for head != nil && prev.val == head.val {
//		head = head.Next
//	}
//	deleteDuplicates0(head, prev)
//}

func main() {

}
