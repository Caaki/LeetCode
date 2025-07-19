package main

import (
	"fmt"
	"slices"
	"strings"
)

func main(){
	fmt.Println(strings.HasPrefix("/a", "/a/b"))
	fmt.Println(strings.HasPrefix("/a/b", "/a"))


}


func removeSubfolders(folder []string) []string {
	slices.Sort(folder)

	result := []string{}
	for _, f := range folder {
		if len(result) == 0 || !strings.HasPrefix(f, result[len(result)-1]+"/") {
			result = append(result, f)
		}
	}
	return result
}
