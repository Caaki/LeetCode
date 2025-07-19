package main


func main(){}

type ListNode struct {
		Val int
		Next *ListNode
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
 	
	start := head
	behind := head
	count :=0
	
	for head!=nil{
		if count > n{
			behind = behind.Next
		}
		head = head.Next
		count++
	}
	
	if count == n{
		return start.Next
	}
	
	behind.Next = behind.Next.Next
	return start

}
