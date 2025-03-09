package main 


func main (){}


func numberOfLaternatingGroups(colors []int, k int) int{
    count:=0
    n:=len(colors)
  for i,j:=0,1; j<len(colors)+(k-1); j++{
    if colors[j%n] == colors[(j-1)%n]{
        i=j
    }
    if j-i+1>=k{
        count++
        i++
    }
  }
  return count
}
