package main

import "fmt"

func main(){
	fmt.Println(defangIPaddr("1.1.1.1"))
}


func defangIPaddr(address string) string {
    newVal:=""
		for _,v :=range address{
			if v!='.'{
				newVal+=string(v)
			}else{
				newVal+="[.]"
			}
		}
		return newVal
}
