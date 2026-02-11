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

func dfs (root *TreeNode) (bool,int){
	if root==nil{
		return true, 0
	}
	
	leftC,leftV := dfs(root.Left)
	rightC,rightV := dfs(root.Right)

	if leftC ==false || rightC == false{
		return false,-1
	}
	
	if rightV<leftV{
		rightV,leftV = leftV,rightV
	}
	
	if rightV - leftV >=2{
		return false,-1
	}
	return true , 1+rightV
}
