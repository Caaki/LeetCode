package main
func main(){

}

func twoSum(nums []int, target int) []int {
    values :=make(map[int16]int16)
    for i,v:=range nums {
    dif := int16(target - v)

        if r,ok:=values[dif] ; ok {
            return []int{int(r),i}
            } 

    values[int16(v)]=int16(i)}
    return nil
}
