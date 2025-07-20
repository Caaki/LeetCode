package main

func main(){}


type Node struct {
	 Val int
	 Next *Node
	 Random *Node
}


func copyRandomList(head *Node) *Node {

    if head == nil{
        return nil
    }

	arrVals := make([]*Node,0)
	dupedVals:=make([]*Node,0)
	mapppedVals := make(map[*Node]*Node,0)

	for head!=nil{
		arrVals = append(arrVals, head)
		nVal := &Node{Val: head.Val}
		dupedVals = append(dupedVals, nVal)
		mapppedVals[head]= nVal
        head=head.Next
	}

	for i,v :=range arrVals{
		dupedVals[i].Next=mapppedVals[v.Next]
		dupedVals[i].Random=mapppedVals[v.Random]
	}	
	return dupedVals[0]
	
}




