package main

import "fmt"

func main() {
	nums := []int{1, 3, 2, 3, 1}
	fmt.Println(reversePairs(nums))

	nums = []int{-5, -5}
	fmt.Println(reversePairs(nums))
}

func reversePairs(nums []int) int {
	temp := make([]int, len(nums))
	return reversePairsRecursive(nums, 0, len(nums)-1, temp)
}

func reversePairsRecursive(nums []int, l, r int, temp []int) int {
	if l == r {
		return 0
	}

	mid := l + (r-l)/2

	var ret int
	ret += reversePairsRecursive(nums, l, mid, temp)
	ret += reversePairsRecursive(nums, mid+1, r, temp)
	ret += merge(nums, l, mid, r, temp)
	return ret
}

func merge(nums []int, l, m, r int, temp []int) int {
	for i := l; i <= r; i++ {
		temp[i] = nums[i]
	}

	var ret int
	for k, i, j := l, l, m+1; i <= m || j <= r; k++ {
		if j > r {
			nums[k] = temp[i]
			i++
		} else if i > m {
			nums[k] = temp[j]
			j++
		} else if i <= m && temp[i] <= temp[j] {
			nums[k] = temp[i]
			ret += countRightReversePairs(temp, j, r, nums[k])
			i++
		} else {
			nums[k] = temp[j]
			ret += countLeftReversePairs(temp, i, m, nums[k])
			j++
		}
	}
	return ret
}

func countLeftReversePairs(nums []int, l int, r int, v int) int {
	if l > r {
		return 0
	}
	i := getFirstReversePairIndex(nums, l, r, v)
	if i < 0 {
		return 0
	}
	return r - i + 1
}

// 在区间 [l..r] 中找到第一个值大于 2 * v 的下标
// 找不到返回 -1
func getFirstReversePairIndex(nums []int, l int, r int, v int) int {
	if l == r {
		if nums[l] > 2*v {
			return l
		} else {
			return -1
		}
	}

	mid := l + (r-l)/2
	if nums[mid] > 2*v {
		return getFirstReversePairIndex(nums, l, mid, v)
	}
	return getFirstReversePairIndex(nums, mid+1, r, v)
}

func countRightReversePairs(nums []int, l int, r int, v int) int {
	if l > r {
		return 0
	}
	i := getLastReversePairIndex(nums, l, r, v)
	if i < 0 {
		return 0
	}
	return i - l + 1
}

// 在区间 [l..r] 中找到最后一个值小于 v/2 的下标
// 找不到返回 -1
func getLastReversePairIndex(nums []int, l int, r int, v int) int {
	if l == r {
		if v > nums[l]*2 {
			return l
		} else {
			return -1
		}
	}

	mid := l + (r-l)/2
	if v > nums[mid]*2 {
		return max(mid, getLastReversePairIndex(nums, mid+1, r, v))
	}
	return getLastReversePairIndex(nums, l, mid, v)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
