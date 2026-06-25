package main

import "fmt"

func main (){
	fmt.Println(longestBalanced("abbac"))
}

func longestBalanced(s string) int {
	
	return longestSameCaracter(s)
	

}

func longestSameCaracter (s string) int{
	current :=s[0]
	max :=0
	temp:=0
	for _,v :=range s{
		if v == rune(current){
			temp++
		}else{
			current=byte(v)
			if max<temp{
				max = temp
			}
			temp = 1
		}
	}

	return max

}
