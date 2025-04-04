package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {}

func maxDepth(root *TreeNode) int {

	return findMax(root)

}

func findMax(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + int(math.Max(float64(findMax(root.Left)), float64(findMax(root.Right))))

}
