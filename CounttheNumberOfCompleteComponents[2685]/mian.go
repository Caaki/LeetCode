package main

import "fmt"

func main() {

	fmt.Println(countCompleteComponents(4, [][]int{{2, 1}, {3, 0}, {3, 1}, {3, 2}}))

}

type Comp struct {
	vals  map[int]bool
	edges int
}

type Collection struct {
	individual []*Comp
	count      int
}

func (c *Collection) Add(v []int) {
	for _, val := range c.individual {
		if _, ok := val.vals[v[0]]; ok {
			if _, ok2 := val.vals[v[1]]; !ok2 {
				val.vals[v[1]] = true
				c.count++
			}
			val.edges++
			return
		}
		if _, ok := val.vals[v[1]]; ok {
			if _, ok2 := val.vals[v[0]]; !ok2 {
				val.vals[v[0]] = true
				c.count++
			}
			val.edges++
			return
		}
	}
	comp := Comp{vals: map[int]bool{v[0]: true, v[1]: true}, edges: 1}
	c.individual = append(c.individual, &comp)
	c.count += 2
}

func (c *Comp) isValid() bool {
	return len(c.vals)*(len(c.vals)-1)/2 == c.edges
}
func (c *Collection) Count(n int) int {
	if len(c.individual) == 0 {
		return n
	}
	count := n - c.count

	for _, v := range c.individual {
		if v.isValid() {
			count++
		}
	}

	return count
}

func isSubset(smaller, larger map[int]bool) bool {
	for key := range smaller {
		if larger[key] {
			return true
		}
	}
	return false
}

func (c *Collection) removeRedundantMaps() {
	var filtered []*Comp

	for i, comp1 := range c.individual {
		redundant := false
		for j, comp2 := range c.individual {
			if i != j && len(comp1.vals) <= len(comp2.vals) && isSubset(comp1.vals, comp2.vals) {
				redundant = true
				comp2.edges += comp1.edges
				c.count -= len(comp1.vals)
				for k, v := range comp1.vals {
					if _, ok := comp2.vals[k]; !ok {
						c.count++
					}
					comp2.vals[k] = v
				}
				break
			}
		}
		if !redundant {
			filtered = append(filtered, comp1)
		}
	}

	c.individual = filtered
}

func countCompleteComponents(n int, edges [][]int) int {
	c := Collection{count: 0}
	for _, v := range edges {
		c.Add(v)
	}

	c.removeRedundantMaps()

	return c.Count(n)

}
