package main


func main(){

}

func twoSum(nums []int, target int) []int {
    values :=make(map[int16]int16) 
    defer clear(values) 
    for index,value:=range nums {
        get,ok:=values[int16(target-value)] 
        if ok {
            return []int{int(get),index}
            } 
    values[int16(value)]=int16(index)}
    return nil
}
