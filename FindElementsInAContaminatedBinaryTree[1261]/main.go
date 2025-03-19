package main

func main(){}

  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
 
type FindElements struct {
  values map[int]bool
}

func (this *FindElements) Add(v int){
  this.values[v] = true
}

func Constructor(root *TreeNode) FindElements {
  findElements:= FindElements{values:make(map[int]bool,100)}
  findElements.Add(0)
  root.Val = 0
  recursiveAdd(root, findElements)
  return findElements
}

func recursiveAdd(root *TreeNode,elm FindElements)(){
  if root.Left!= nil{
    root.Left.Val = (root.Val*2+1)
    elm.Add(root.Left.Val)
    recursiveAdd(root.Left, elm)
  }
  if root.Right!=nil{
    root.Right.Val = (root.Val*2+2)
    elm.Add(root.Right.Val)
    recursiveAdd(root.Right,elm)
  }
}

func (this *FindElements) Find(target int) bool {
  _,ok:= this.values[target]
  return ok
}

