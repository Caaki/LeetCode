package main


type key struct {
	l,r int
}
func stoneGame(piles []int) bool {

	mem := make(map[key] int)

	var game func(int,int) int
	game = func(l,r int) int {
		if r-l == 1{
			if piles[r]> piles[l]{
				return piles[r]
			}
			return piles[l]
		}
		k := key{l,r}
		if v,ok := mem[k]; ok {
			return v 
		}

		left := piles[l] - game(l,r-1)
		right := piles[r] - game(l+1,r)

		if left > right{
			mem[k] = left
			return left
		}
		mem[k]=right
		return right
	}

	return game(0,len(piles)-1) > 0

}

//Second solution
//func stoneGame(piles []int) bool{
//return true
//}


