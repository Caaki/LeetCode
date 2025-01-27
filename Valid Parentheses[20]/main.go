package main

import (
	"fmt"
	"math/big"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func main() {

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	stringVal := ""
	for l1 != nil {
		stringVal = strconv.Itoa(l1.Val) + stringVal
		l1 = l1.Next
	}
	fmt.Println("String val is:", stringVal)
	val1, _ := new(big.Int).SetString(stringVal, 10)
	fmt.Println("Int64Val is: ", val1)
	stringVal = ""
	for l2 != nil {
		stringVal = strconv.Itoa(l2.Val) + stringVal
		l2 = l2.Next
	}

	val2, _ := new(big.Int).SetString(stringVal, 10)

	newBigInt := big.NewInt(int64(43))
	newBigInt.Add(val2, val1)
	result := newBigInt.String()
	list := LinkedList{}
	for _, c := range result {
		list.insertAtFront(int(c - '0'))
	}
	return list.head
}

func (list *LinkedList) insertAtFront(data int) {
	if list.head == nil {
		newNode := &ListNode{
			Val:  data,
			Next: nil,
		}
		list.head = newNode
		return
	}
	newNode := &ListNode{
		Val:  data,
		Next: list.head,
	}
	list.head = newNode

}
