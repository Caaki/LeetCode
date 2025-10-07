package main

import "fmt"

func main(){
	
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))
}
func gcdOfStrings(str1 string, str2 string) string {
	
	if str1[0]!= str2[0]{
		return ""
	}

	if len(str1)< len(str2){
		str1,str2 = str2,str1
	}


	for i:=len(str2) ;i>=1; i-- {
			
		word := str2[0:i]
		if len(str1) % len(word)!=0 || len(str2)%len(word)!=0{
			continue
		}
		if checkIfDivisor(str1, word) && checkIfDivisor(str2, word){
			return word
		}
	}
	return ""
}

func checkIfDivisor(str, word string) bool {

	for i:=len(word); i<=len(str); i+=len(word){
		temp := str[i-len(word):i]
		if temp!=word{
			return false
		}
	} 
	return true
}
