package main

func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	n := &ListNode{}
	first := n

	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		n.Val = list1.Val
		list1 = list1.Next
	} else {
		n.Val = list2.Val
		list2 = list2.Next
	}

	for !(list1 == nil || list2 == nil) {
		if list1.Val < list2.Val {
			n.Next = &ListNode{Val: list1.Val}
			n = n.Next
			list1 = list1.Next
		} else {
			n.Next = &ListNode{Val: list2.Val}
			n = n.Next
			list2 = list2.Next
		}
	}
	if list1 == nil && list2 == nil {
		return first
	}
	if list1 == nil {
		n.Next = list2
		return first
	}

	if list2 == nil {
		n.Next = list1
		return first
	}
	return first
}
