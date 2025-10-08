package main

import "strings"

func reverseVowels(s string) string {
	
	type pair struct {
		volw string
		place int
	}

	vowels:= make([]pair,0)
	vString := "aeiouAEIOU"

	for i,v := range s {
		if strings.Contains(vString, string(v)){
			vowels = append(vowels,pair{volw: string(v),place: i} )
		}
	}
	
	nString :=""
	count := len(vowels)-1
	for _,v:=range s{
		if strings.Contains(vString, string(v)){
			nString+=vowels[count].volw
			count--
			continue
		}
		nString+=string(v)
	}
	return nString
}
