package main

func main (){}

var volvs =[]bool{
  false,    //A
  false,    //E
  false,    //I
  false,    //O
  false,    //U
  false,    //nothing
}

//A=1
//E=10
//I=100
//O=1000
//U=10000
//Max=11111
//l=0

func countOfSubstrings(word string, k int) int64 {
  nVolvs := make([]bool, len(volvs))
  copy(nVolvs, volvs)
  l:=0
  r:=1
  sum:=0
  kNums:=0
  count:=int64(0)
  
  if pos(word[0])==5{
    kNums++
  }else{
    sum+=val(word[0])
    nVolvs[pos(word[0])] = true
  }


  for r<len(word){
    fmt.Println(nVolvs,volvs)
    if pos(word[r])==5{
      kNums++
    }else{
      if !nVolvs[pos(word[r])]{
        sum+=val(word[r])
        nVolvs[pos(word[r])]=true
      }
    }
    if sum==11111 && kNums==k{
      count++
      sum-=val(word[l])
      nVolvs[pos(word[l])] = false
      l++
      continue
    }
    
    if kNums>k{
      l=r
      copy(nVolvs, volvs)
      sum=0
      kNums=0
    }
    // fmt.Println(nVolvs,sum,kNums)


    r++
  }
  return count
} 

func pos(r byte) int{
  if r=='a'{
    return 0
  }
  if r=='e'{
    return 1
  }
  if r=='i'{
    return 2
  }
  if r=='o'{
    return 3
  }

  if r=='u'{
  return 4
  }
  return 5
}


func val(r byte) int{
  if r=='a'{
    return 1
  }
  if r=='e'{
    return 10
  }
  if r=='i'{
    return 100
  }
  if r=='o'{
    return 1000
  }
  if r=='u'{
  return 10000
  }
  return 0
}
