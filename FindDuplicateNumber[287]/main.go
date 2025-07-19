package main


func main() {}

func findDuplicate(nums []int) int {

	values := make(map[int]bool,0)


	for _, v:= range nums{
		_,ok:=values[v]
		if ok{
			return v
		}else{
			values[v]=true
		}
	}
	return 0
}
