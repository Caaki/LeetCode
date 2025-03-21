package main

func main() {}

type Ingredient struct {
	name   *string
	needed *[]string
	basic  bool
}

func createBasicIngredients(name string) *Ingredient {
	return &Ingredient{name: &name,
		basic: true}
}

func createNonBasicIngredients(name *string, needed *[]string) *Ingredient {
	return &Ingredient{name: name, needed: needed, basic: false}
}

func (ing *Ingredient) CanBeCreated(recipes map[string]*Ingredient, ingredients map[string]*Ingredient) bool {
	for _, v := range *(ing.needed) {
		if _, ok := ingredients[v]; ok {
			continue
		}
		if _, ok := recipes[v]; !ok {
			return false
		}
		if !recipes[v].CanBeCreated(recipes, ingredients) {
			return false
		}
	}
	return true
}

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	basicIngredients := make(map[string]*Ingredient, len(supplies))
	for _, v := range supplies {
		basicIngredients[v] = createBasicIngredients(v)
	}
	needed := make(map[string]*Ingredient, len(recipes))
	for i, v := range recipes {
		needed[v] = createNonBasicIngredients(&v, &ingredients[i])
	}

	returnVals := []string{}
	for _, v := range recipes {
		if needed[v].CanBeCreated(needed, basicIngredients) {
			basicIngredients[v] = needed[v]
			returnVals = append(returnVals, v)
		}
	}
	return returnVals
}
