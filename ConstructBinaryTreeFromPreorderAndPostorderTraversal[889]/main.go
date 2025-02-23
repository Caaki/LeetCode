package main


func main(){}

type TreeNode struct{
  Val int
  Left *TreeNode
  Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
  root:= TreeNode{Val:preorder[0]}
  if len(preorder) <=3{
    return lessThan3(&root,preorder,postorder)
  }

  recursiveAdd(&root,preorder,postorder)
  return &root
}

func lessThan3(root *TreeNode,preorder []int, postorder []int) *TreeNode{
    if len(preorder)==1{
    return root
  }
  if len(preorder) ==2{
    root.Left= &TreeNode{Val:preorder[1]}
    return root
  }
if len(preorder) ==3{
    if preorder[1] == postorder[1]{
      root.Left= &TreeNode{Val:preorder[1],Left:&TreeNode{Val:preorder[2]}}
      return root
    }
    root.Left= &TreeNode{Val:preorder[1]}
    root.Right= &TreeNode{Val:preorder[2]}
    return root
  }
  return root
}

func recursiveAdd(root *TreeNode, preorder []int, postorder []int){
  if len(preorder)<= 3{
    lessThan3(root,preorder,postorder)
    return
  }
  count:=0
  for i:=0; i<len(postorder) ; i++{
    if preorder[1] != postorder[i]{
      count++
      continue
    }
    break
  }
  leftSide:=postorder[:count+1]
  rightSide:=postorder[count+1:len(postorder)-1]
  root.Left= &TreeNode{Val:preorder[1]}

    if len(rightSide)>0{
  root.Right=&TreeNode{Val:rightSide[len(rightSide)-1]}
    }
  if len(leftSide)>1{
  recursiveAdd(root.Left, preorder[1:count+2],postorder[:count+1])
  }
  if len(rightSide)>1{
  recursiveAdd(root.Right, preorder[count+2:],postorder[count+1:len(postorder)-1])
  }
}

