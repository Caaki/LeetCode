package main

import "sort"

func main() {}

var t int32

func carFleet(target int, position []int, speed []int) int {
	t = int32(target)
	cars := make([]Car, len(position))

	for i := range len(position) {
		cars[i] = Car{position: int32(position[i]), speed: int32(speed[i])}
	}

	sort.Sort(carSorting(cars))
	carStack := CarStack{[]Car{cars[0]}}

	for i := 1; i < len(cars); i++ {
		if carStack.ArrivalTime() < cars[i].ArrivalTime() {
			carStack.Push(cars[i])
		}
	}

	return carStack.Len()
}
