// https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/

package main

import (
	"fmt"
	"math"
)

func reversePairs(nums []int) int {
	tmp := make([]int, len(nums), len(nums))
	copyArray(nums, 0, len(nums)-1, tmp)

	res, n := 0, len(nums)

	for sz := 1; sz < n; sz += sz {
		for i := 0; i+sz < n; i += sz + sz {
			if nums[i+sz-1] > nums[i+sz] {
				res += merge(nums, i, i+sz-1, int(math.Min(float64(i+sz+sz-1), float64(n-1))), tmp)
			}
		}
	}
	return res
}

func merge(arr []int, l, m, r int, tmp []int) int {
	copyArray(arr, l, r, tmp)

	res := 0
	i, j := l, m+1
	for k := l; i <= m || j <= r; k++ {
		if i > m {
			arr[k] = tmp[j]
			j++
		} else if j > r {
			arr[k] = tmp[i]
			i++
		} else if tmp[i] > tmp[j] {
			res += m - i + 1
			arr[k] = tmp[j]
			j++
		} else {
			arr[k] = tmp[i]
			i++
		}
	}

	return res
}

func copyArray(source []int, l, r int, dest []int) {
	for i := l; i <= r; i++ {
		dest[i] = source[i]
	}
}

func main() {
	arr := []int{7, 5, 6, 4}
	fmt.Println(reversePairs(arr))
	fmt.Println(arr)

	var prev int
	for i, v := range arr {
		if i > 0 && v < prev {
			panic("Not a sorted slice")
		}
		prev = v
	}
}
