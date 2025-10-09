package main

import "strings"


func reverseWords(s string ) string {

	words := strings.Fields(s)
	
	toReturn :=words[len(words)-1]
	for i:=len(words)-2; i >=0;i-- {
		toReturn+=" "+words[i]
	}
	
	return toReturn
}
