package main

import "strconv"

type ListNode struct {
	Val int
	Next *ListNode
}

func getDecimalValue(head *ListNode)int{

	value:=""
	
	for head!= nil{
		value+= strconv.Itoa(head.Val)
		head = head.Next
	}
	rv,_ := strconv.ParseInt(value, 2, 64)
		
	return int(rv)
}
