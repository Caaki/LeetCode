package main

import (
	"fmt"
	"sort"
)

type Car struct {
	pos   int
	speed int
	same  bool
}

type CarArray []*Car

func (c CarArray) Less(i, j int) bool {
	return c[i].pos > c[j].pos
}

func (c CarArray) Len() int {
	return len(c)
}
func (c CarArray) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func carFleet(target int, position []int, speed []int) int {
	end := 0

	cars := make([]*Car, len(position))

	for i, v := range speed {
		cars[i] = &Car{pos: position[i], speed: v}
	}
	c := CarArray(cars)
	sort.Sort(c)
	for _, v := range cars {
		fmt.Print(*v)
	}
	fmt.Println()

	fleets := 0
	for len(cars) > 0 {
		atEnd := false
		for i := 0; i < len(cars); i++ {
			cars[i].pos += cars[i].speed
			if cars[i].pos >= target {
				atEnd = true
				continue
			}
			if i != 0 {
				if cars[i-1].pos <= cars[i].pos {
					cars[i].pos = cars[i-1].pos
					if cars[i].speed < cars[i-1].speed {
						cars[i-1].speed = cars[i].speed
					} else {
						cars[i].speed = cars[i-1].speed
					}
				}
			}
		}

		for _, v := range cars {
			fmt.Print(*v)
		}
		fmt.Println()
		if atEnd {
			for len(cars) > 0 && cars[0].pos >= target {
				end = 0
				first := cars[0].pos
				for _, v := range cars {
					if first != target && (v.pos != first || v.speed != cars[0].speed) && v.pos > target {
						fleets++
					}
					if v.pos >= target {
						end++
					}
				}
				fleets++
				cars = cars[end:]
			}

			for _, v := range cars {
				fmt.Print(*v)
			}
			fmt.Println()
			fmt.Println(fleets)
		}
	}

	return fleets

}
