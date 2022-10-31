package main

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})

	n := len(people)
	ret := make([][]int, n)
	for _, p := range people {
		space := p[1] + 1

		i := 0
		for ; i < n; i++ {
			if ret[i] == nil {
				space--
			}
			if space == 0 {
				ret[i] = p
				break
			}
		}
	}
	return ret
}

func main() {

}
