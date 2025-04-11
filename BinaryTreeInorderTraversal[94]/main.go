package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	arr := make([]int, 0)
	inOrder(root, &arr)
	return arr
}

func inOrder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	inOrder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inOrder(root.Right, arr)

}
