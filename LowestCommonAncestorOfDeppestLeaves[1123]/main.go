package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {}

func lcaDeppestLeaves(root *TreeNode) *TreeNode {

	v, _ := FindBestNode(root)
	return v
}

func FindBestNode(root *TreeNode) (value *TreeNode, h int) {

	if root == nil {
		return nil, 0
	}

	lNum, lh := FindBestNode(root.Left)
	rNum, rh := FindBestNode(root.Right)
	if lh == rh {
		return root, 1 + lh
	}

	if lh > rh {
		return lNum, lh + 1
	}

	return rNum, rh + 1

}
