package main

import "fmt"

func main(){
	fmt.Println(mergeAlternately("aaaaa", "bb"))

}

func mergeAlternately(word1,word2 string) string {
	i :=0
	j:=0
	
	toReturn := []byte{}

	for i < len(word1) && j < len(word2){
		toReturn = append(toReturn, word1[i])
		toReturn = append(toReturn, word2[j])
		i++
		j++
	}
	
	toReturn = append(toReturn, word1[i:]...)
	toReturn = append(toReturn, word2[j:]...)
	
	return string(toReturn)
}
