package main


func main(){}

type Paths struct{
  path []int
  isBob bool
}

func mostProfitablePath(edges [][]int, bob int, amount []int) int{
  if len(edges) ==1{
      return amount[0]
    }
  alieceMax := amount[0]
  alieceCurrent:=amount[0]
  amount[bob]=0


  return alieceMax
}

func findEdgeNumbers(all [][]int) map[int][]Paths{
  paths := map[int]Paths{}

  currentRoot := all[0][0]
  paths[all[0][0]] = Paths{path:[]int{all[0][0]},isBob:false}
  for i:=0;i < len(all) ; i++{
    paths[all[i][1]].path = append(paths[all[i][0]].path,all[i][1])
    if currentRoot!=all[i][0]{
      delete(paths, currentRoot)
      currentRoot=all[i][1]
    }
  }
  return paths
}

func recursiveTraversal(edges [][]int,current int, aliece int,bob *int, amount *[]int, max *int){
}


//[0] [1] 
//[1] [2,3]
//[3] [4]


