package main

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	cookable := map[string]bool{}
	indexes := map[string]int{}

	for _, v := range supplies {
		cookable[v] = true
	}

	for i, v := range recipes {
		indexes[v] = i
	}

	rValues := []string{}

	for _, v := range recipes {
		if canBeCreated(cookable, indexes, v, ingredients) {
			rValues = append(rValues, v)
		}
	}
	return rValues
}

func canBeCreated(cookable map[string]bool, indexes map[string]int, r string, ingredients [][]string) bool {
	if _, ok := cookable[r]; ok {
		return cookable[r]
	}
	if _, ok := indexes[r]; !ok {
		return false
	}

	cookable[r] = false

	for _, v := range ingredients[indexes[r]] {
		if !canBeCreated(cookable, indexes, v, ingredients) {
			return false
		}
	}
	cookable[r] = true
	return cookable[r]
}
