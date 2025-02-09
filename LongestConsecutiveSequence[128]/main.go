package main

type Values struct {
	Lower  []int16
	Higher []int16
}

type Values2 struct {
	Lowest  int
	Highest int
}

func main() {

}

func longestConsecutive(nums []int) int {
	values := make(map[int16]*Values)
	longest := 0
	for _, num := range nums {
		if _, ok := values[int16(num)]; ok {
			continue
		}
		val := Values{Higher: []int16{}, Lower: []int16{}}
		values[int16(num)] = &val
		getFirst(values, int16(num), int16(num))
		getLast(values, int16(num), int16(num))
		if longest < len(values[int16(num)].Higher)+len(values[int16(num)].Lower)+1 {
			longest = len(values[int16(num)].Higher) + len(values[int16(num)].Lower) + 1
		}
	}
	return longest
}

func getFirst(values map[int16]*Values, num int16, appendTo int16) {
	if _, ok := values[num-1]; ok {
		values[appendTo].Lower = append([]int16{num - 1}, values[appendTo].Lower...)
		if len(values[num-1].Lower) > 0 {
			values[appendTo].Lower = append(values[num-1].Lower, values[appendTo].Lower...)
		}

		getFirst(values, values[appendTo].Lower[0], appendTo)
	}
}

func getLast(values map[int16]*Values, num int16, appendTo int16) {
	if _, ok := values[num+1]; ok {
		values[appendTo].Higher = append(values[appendTo].Higher, num+1)
		if len(values[num+1].Higher) > 0 {
			values[appendTo].Higher = append(values[appendTo].Higher, values[num+1].Higher...)
		}
		getLast(values, values[appendTo].Higher[len(values[appendTo].Higher)-1], appendTo)
	}
}
