package main

import "fmt"

func main() {

	fmt.Println(longestCommonPrefix([]string{"flower", "fkow"}))

}
func longestCommonPrefix(strs []string) string {
	prefix := []rune(strs[0])
	for _, s := range strs {
		if len(prefix) == 0 {
			return ""
		}
		temp := make([]rune, 0, len(prefix))
		for i, c := range prefix {
			if len(s) <= i {
				prefix = temp
				break
			}
			if rune(s[i]) != c {
				prefix = temp
				break
			} else {
				temp = append(temp, c)
			}
		}
	}
	return string(prefix)
}
