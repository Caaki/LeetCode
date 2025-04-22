package main

func findRepeatedDnaSequences(s string) []string {
	values := make(map[string]int16, 0)

	if len(s) < 11 {
		return []string{}
	}
	retVals := make([]string, 0)

	for i := 9; i < len(s); i++ {
		v := s[i-9 : i+1]
		values[v]++
		if values[v] == 2 {
			retVals = append(retVals, v)
		}
	}
	return retVals
}
