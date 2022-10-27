package main

import (
	"fmt"
)

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func getSkyline(buildings [][]int) [][]int {
	if len(buildings) == 0 {
		return nil
	}

	return getSkyline0(buildings, 0, len(buildings)-1)
}

func getSkyline0(buildings [][]int, l, r int) [][]int {
	if l >= r {
		return [][]int{
			{buildings[l][0], buildings[l][2]},
			{buildings[l][1], 0},
		}
	}

	mid := l + (r-l)/2
	left := getSkyline0(buildings, l, mid)
	right := getSkyline0(buildings, mid+1, r)
	ret := merge(left, right)
//	fmt.Println(ret)
	return ret
}

func push(ret [][]int, next []int) [][]int {
	if len(ret) == 0 {
		return [][]int{next}
	}

	// h 相等，不追加
	if ret[len(ret)-1][1] == next[1] {
		return ret
	}

	return append(ret, next)
}

func merge(l, r [][]int) [][]int {
	fmt.Printf("l:%v\n", l)
	fmt.Printf("r:%v\n", r)
	var ret [][]int

	i, j := 0, 0
	hl, hr := 0, 0
	for i < len(l) || j < len(r) {
		var x int

		if j >= len(r) || (i < len(l) && l[i][0] < r[j][0]) {
			x, hl = l[i][0], l[i][1]
			h := max(hl, hr)
			ret = push(ret, []int{x, h})
			i++
		} else if i >= len(l) || (j < len(r) && l[i][0] > r[j][0]) {
			x, hr = r[j][0], r[j][1]
			h := max(hl, hr)
			ret = push(ret, []int{x, h})
			j++
		} else {
			x, hl, hr = l[i][0], l[i][1], r[j][1]
			h := max(hl, hr)
			ret = push(ret, []int{x, h})
			i++
			j++
		}
	}
	return ret
}

func main() {
	buildings := [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}
	fmt.Print(getSkyline(buildings))

	buildings = [][]int{{0, 2, 3}, {2, 5, 3}}
	fmt.Print(getSkyline(buildings))
}
