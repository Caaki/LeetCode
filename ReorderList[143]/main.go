package main

func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	positions := make([]*ListNode, 0)

	for head != nil {
		positions = append(positions, head)
		head = head.Next
	}
	n := len(positions)
	next := positions[0]
	last := positions[n-1]
	finish := positions[n/2]

	n--
	for last != finish {
		temp := next.Next
		next.Next = last
		last.Next = temp
		next = temp
		n--
		last = positions[n]
	}

	finish.Next = nil

}
