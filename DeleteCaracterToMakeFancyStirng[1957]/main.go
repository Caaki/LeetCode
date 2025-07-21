package main

import "strings"

func makeFancyString(s string) string {
	var builder strings.Builder
	count :=0
	var current rune
	for _,v:=range s{
		if current==v{
			count++
			if count > 2 {
				continue
			}
			builder.WriteRune(v)
		}else{
			current = v
			count =1
			builder.WriteRune(v)
		}
	}
	return builder.String()
}
