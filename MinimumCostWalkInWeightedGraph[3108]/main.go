package main

import(
  "fmt"
)

func main(){
  fmt.Println(minimumCost(5,[][]int{{0,1,0},{1,3,7},{1,2,1}},[][]int{{0,3},{3,4}}))
}

type UF struct {
  root []int
  rank []int
}
func minimumCost(n int, edges [][]int, query [][]int) []int {
  uf:=NewUnionFind(n)

  for _,e := range edges{
    uf.Union(e[0],e[1])
  }
  
  count:=make(map[int]int)

  for _,e:=range edges{
    root:= uf.Find(e[0])
    if _,ok := count[root];!ok{
      count[root] = e[2]
    }else{
    count[uf.Find(e[0])]&=e[2]
    }
  }
  
  answers:=make([]int,len(query))

  for i,v:= range query{
    if uf.Find(v[0]) != uf.Find(v[1]){
      answers[i]=-1
    }else{
      answers[i] =count[uf.Find(v[0])]
    }
  }
  
  return answers
}

func NewUnionFind(n int) *UF {
  uf := &UF{
    root: make([]int,n+1),
    rank: make([]int,n+1),
  }
  for i:=0;i<=n; i++{
    uf.root[i]=i
    uf.rank[i]=1
  }
  return uf
}

func (uf *UF) Find(x int) int{
  if uf.root[x]!=x{
    uf.root[x] = uf.Find(uf.root[x])
  }
  return uf.root[x]
}

func (uf *UF) Union(x,y int) {
  rootX:=uf.Find(x)
  rootY:= uf.Find(y)

  if rootX!=rootY{
    if uf.rank[rootX] > uf.rank[rootY] {
      uf.root[rootY]=rootX
      return
    }
    if uf.rank[rootX]<uf.rank[rootY]{
      uf.root[rootX]=rootY
      return
    }
    uf.root[rootY]=rootX
    uf.rank[rootX]++
  }
}

