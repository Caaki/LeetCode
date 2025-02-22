package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	str := "1234"
	fmt.Println(myAtoi(&str))
	fmt.Println(str)

	node := recoverFromPreorder("1-401--349---90--88")
	node.printTree()
}
func (tree *TreeNode) printTree() {
	fmt.Println(tree.Val)
	if tree.Left != nil {
		tree.Left.printTree()
	}
	if tree.Right != nil {
		tree.Right.printTree()
	}
}
func recoverFromPreorder(traversal string) *TreeNode {
	num := myAtoi(&traversal)
	root := TreeNode{Val: num}
	addNumToTree(&root, &traversal, 1)
	return &root
}

func addNumToTree(root *TreeNode, str *string, currentBranch int) {
	if len((*str)) == 0 {
		return
	}
	bd := countDashes(str)
	if bd < currentBranch {
		return
	}
	if bd == currentBranch {
		if root.Left == nil {
			*str = (*str)[bd:]
			tn := TreeNode{Val: myAtoi(str)}
			root.Left = &tn
			addNumToTree(root.Left, str, currentBranch+1)
		}

		if root.Right == nil && len((*str)) >= 1 && countDashes(str) == bd {
			*str = (*str)[bd:]
			tn := TreeNode{Val: myAtoi(str)}
			root.Right = &tn
			addNumToTree(root.Right, str, currentBranch+1)
		}
	}
}

func countDashes(str *string) int {
	count := 0
	for i := 0; i < len(*str); i++ {
		if (*str)[i] == '-' {
			count++
			continue
		}
		break
	}
	return count
}

func myAtoi(s *string) int {
	var n int32 = 0
	i := 0
	for _, c := range *s {
		if c > 57 || c < 48 {
			break
		}
		if i > 10 {
			return math.MaxInt32
		}
		if i >= 9 {
			if (int(n)*10)+int(c)-48 > math.MaxInt32 {
				return math.MaxInt32
			}
			n = n*10 + c - 48
			i++
			continue
		}
		n = n*10 + c - 48
		if n != 0 {
			i++
		}
	}
	*s = (*s)[i:]
	return int(n)
}
