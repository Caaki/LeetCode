package main


type TreeNode struct{

	Val int
	Left *TreeNode
	Right *TreeNode

}

func isBalanced(root *TreeNode)bool{
	
	balanced,_ := dfs(root)
	return balanced
}

func dfs(root *TreeNode)(bool,int){
	if root == nil {
		return true,0
	}
	
	leftCheck,leftVal := dfs(root.Left)
	rightCheck,rightVal := dfs(root.Right)

	if leftCheck==false || rightCheck == false{
		return false,-1
	}

	if rightVal < leftVal{
		rightVal,leftVal = leftVal,rightVal
	}
	if rightVal-leftVal >=2{
		return false,-1
	}
	
	return true,1+rightVal
}

