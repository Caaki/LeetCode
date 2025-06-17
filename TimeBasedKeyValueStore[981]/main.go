package main

func main() {}

type TimeMap struct {
	m map[string][]Pair
}

type Pair struct {
	k int
	v string
}

func Constructor() TimeMap {
	m := make(map[string][]Pair, 0)
	t := TimeMap{m: m}

	return t
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if v, ok := this.m[key]; ok {
		this.m[key] = append(v, Pair{k: timestamp, v: value})
	} else {
		pairs := make([]Pair, 1)
		pairs[0] = Pair{k: timestamp, v: value}
		this.m[key] = pairs
	}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if v, ok := this.m[key]; ok {

		return searchArray(&v, timestamp)
	} else {

		return ""
	}
}

func searchArray(arr *[]Pair, target int) string {
	l := 0
	r := len(*arr) - 1
	res := -1
	for l <= r {

		mid := l + (r-l)/2
		val := (*arr)[mid].k
		if val == target {
			return (*arr)[mid].v
		}

		if val < target {
			res = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if res == -1 {
		return ""
	}

	return (*arr)[res].v

}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
