package main

import (
	"slices"
)

func compress(chars []byte) int{
	if len(chars)==1{
		return 1
	}
	count := 1
	current := chars[0]
	toWrite:=0

	for i := 1; i < len(chars); i++{
		if chars[i]==current {
			count++
		}else{
			if count == 1{
				chars[toWrite]=current
				toWrite++
				current=chars[i]
				count=1
				continue
			}
			numbers := seperateCharacters(count)
			chars[toWrite] = current
			toWrite++
			for j := 0; j < len(numbers); j++{
				chars[toWrite] = numbers[j]
				toWrite++
			}
			current = chars[i]
			count = 1
		}
	}

	if count == 1{
		chars[toWrite]=current
		toWrite++
	}else{
		numbers := seperateCharacters(count)
		chars[toWrite] = current
		toWrite++
		for j := 0; j < len(numbers); j++{
			chars[toWrite] = numbers[j]
			toWrite++
		}	
	}
	return toWrite

}

func seperateCharacters(num int) []byte{
	numbers := make([]byte,0)
	for num > 0 {
		numbers = append(numbers, byte('0'+num%10))
		num/=10
	}
	slices.Reverse(numbers)
	return numbers
}
