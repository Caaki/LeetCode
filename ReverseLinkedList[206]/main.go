package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var cur *ListNode = nil

	for head != nil {
		next := head.Next
		head.Next = cur
		cur = head
		head = next
	}

	return cur
}
