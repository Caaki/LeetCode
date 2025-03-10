package main

import(
  "strings"
)

func CountOfSubstrings(word string,k int) int64{
    
  return count2(word,k) - count2(word,k+1)
}

func count2(word string, k int) int64{
  vowel:= make(map[byte]int,0)
  nv:=0
  res:=int64(0)
  l:=0

  for r:=0;r<len(word);r++{
    if strings.ContainsRune("aeiou",rune(word[r])){
      vowel[word[r]]+=1
    }else{
      nv+=1
    }
    for len(vowel)==5 && nv >=k{
      res +=int64((len(word)-r))
      if strings.ContainsRune("aeiou",rune(word[l])){
        vowel[word[l]]-=1
      }else{
        nv-=1
      }
      if vowel[word[l]]==0{
        delete(vowel,word[l])
      }
      l+=1
    }
  }
  return res

}
