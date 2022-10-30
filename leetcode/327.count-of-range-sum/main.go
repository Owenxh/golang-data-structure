package main

import "fmt"

func main() {
	nums := []int{-2, 0, 0, 2, 2, -2}
	lower, upper := -3, 1

	fmt.Println(countRangeSum(nums, lower, upper))
}

func preSumSlice(nums []int) []int {
	ret := make([]int, len(nums)+1)
	for i, v := range nums {
		ret[i+1] = ret[i] + v
	}
	return ret
}

func countRangeSum(nums []int, lower int, upper int) int {
	sum := preSumSlice(nums)
	temp := make([]int, len(sum))
	return countRangeSumRecuresive(sum, 0, len(sum)-1, lower, upper, temp)
}

func countRangeSumRecuresive(sum []int, l int, r int, lower int, upper int, temp []int) int {
	if l >= r {
		return 0
	}

	var ret int

	mid := l + (r-l)/2
	ret += countRangeSumRecuresive(sum, l, mid, lower, upper, temp)
	ret += countRangeSumRecuresive(sum, mid+1, r, lower, upper, temp)

	lt, gt := mid+1, mid+1

	for i := l; i <= mid; i++ {
		for lt <= r && sum[lt]-sum[i] < lower {
			lt++
		}
		for gt <= r && sum[gt]-sum[i] <= upper {
			gt++
		}
		ret += (gt - lt)
	}

	merge(sum, l, mid, r, temp)
	return ret
}

func merge(sum []int, l int, m int, r int, temp []int) {
	for i := l; i <= r; i++ {
		temp[i] = sum[i]
	}

	for k, i, j := l, l, m+1; i <= m || j <= r; k++ {
		if j > r || (i <= m && temp[i] <= temp[j]) {
			sum[k] = temp[i]
			i++
		} else {
			sum[k] = temp[j]
			j++
		}

		// if j > r {
		// 	sum[k] = temp[i]
		// 	i++
		// } else if i > m {
		// 	sum[k] = temp[j]
		// 	j++
		// } else if temp[i] <= temp[j] {
		// 	sum[k] = temp[i]
		// 	i++
		// } else {
		// 	sum[k] = temp[j]
		// 	j++
		// }
	}

	// fmt.Println(sum[l : r+1])
}
