package main

func main() {}
func divideString(s string, k int, fill byte) []string {
	elements := make([]string, 0)
	current := ""
	for i := 0; i < len(s); i++ {
		if (i+1)%k == 0 {
			current += string(s[i])
			elements = append(elements, current)
			current = ""
			continue
		}
		current += string(s[i])
	}
	l := len(current)
	if l < k && l != 0 {
		for l < k {
			current += string(fill)
			l = len(current)
		}
		elements = append(elements, current)
	}

	return elements

}
