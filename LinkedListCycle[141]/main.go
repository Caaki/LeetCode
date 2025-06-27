package main

func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head.Next
	slow := head
	for fast != nil {
		fast = fast.Next
		if fast == nil {
			return false
		} else {
			fast = fast.Next
		}
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
