package main

import (
	"slices"
)

//min 3 cars
//A-Z,a-z,0-9
//A-Z=65 - 90
//a-z=97 - 122
//0-9=48 - 57
//one vowel
//one consonant
//vowels {65, 69, 73, 79, 85, 97,101,111, 117}

func isValid(word string ) bool{

	if len(word) < 3{
		return false
	}

	vowels := []rune{65, 69, 73, 79, 85, 97,101,105, 111, 117}
	vowel := false
	consonant :=false

	for _,v := range word{
		if (v>=48 && v<=57 ){
			continue
		}else if (v >=65 && v <=90) || (v>=97 && v <=122){
			if slices.Contains(vowels,v){
				vowel = true
			}else{
				consonant = true
			}

		}else {
			return false
		}
	}
	if consonant && vowel{
		return true
	}
	return false

}
