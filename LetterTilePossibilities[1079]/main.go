package main

func main(){}


func factorial (n int,values map[int]int) int{
  if n ==1{
    return 1
  }
  if val,ok:= values[n] ; ok{
  return values[val]
}

  factorialOfNumber:= n * factorial(n-1, values)
  values[n] = factorialOfNumber
  return factorialOfNumber
}

func numTilePossilities(tiles string) int {
  caracters := make(map[rune]int,0)

  for _,v := range tiles {
    caracters[v]+=1
  }
  
  facNumbs:= make(map[int]int,0)
  sum :=0
  for i := 1 ; i <= len(tiles); i++{
    
  }
}
