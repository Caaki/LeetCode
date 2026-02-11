package main

func main(){}

func maxOperations (nums []int, k int) int {
	count :=0
	m := make(map[int]int,0)
	for _,v := range nums {
		m[v]++
	}
	
	for _,v := range nums {
		if _,ok2:= m[v] ; !ok2{
			continue
		}
		if _,ok := m[k-v]; ok {
			if m[k-v] == k-v {
				if m[k-v]<=1 {
					continue
				}
			}
			count++
			m[k-v]--
			m[v]--
			if m[k-v]==0{
				delete(m, k-v)
			}
			if m[v]==0{
				delete(m, v)
			}
		}
	}
	return count
}
