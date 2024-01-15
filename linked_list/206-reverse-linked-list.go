package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// if LL has 0 or 1 element
	if head == nil || head.Next == nil {
		return head
	}

	curr := head
	var prev, next *ListNode

	for curr != nil {
		next = curr.Next
		curr.Next = prev

		prev = curr
		curr = next
	}

	return prev
}
