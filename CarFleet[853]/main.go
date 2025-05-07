package main

import "sort"

func main() {}

var t int = 0

func carFleet(target int, position []int, speed []int) int {
	t = target
	cars := make([]Car, len(position))

	for i := range len(position) {
		cars[i] = Car{position: position[i], speed: speed[i]}
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
