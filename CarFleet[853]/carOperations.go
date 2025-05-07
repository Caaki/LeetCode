package main



type Car struct {
	position int
	speed    int
}

func (c Car) ArrivalTime() float32 {
	return float32(t-c.position) / float32(c.speed)
}

type carSorting []Car

func (c carSorting) Len() int {
	return len(c)
}

func (c carSorting) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c carSorting) Less(i, j int) bool {
	return c[i].position > c[j].position
}

type CarStack struct {
	Cars []Car
}

func (s *CarStack) Len() int {
	return len(s.Cars)
}
func (s *CarStack) Push(c Car) {
	s.Cars = append(s.Cars, c)
}

func (s *CarStack) ArrivalTime() float32 {
	if s.Len() <= 0 {
		return -1
	}
	c := s.Cars[s.Len()-1]
	return c.ArrivalTime()
}
