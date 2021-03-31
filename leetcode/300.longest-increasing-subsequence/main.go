package main

import "math"

// dp[j] 表示 nums[0...j] 中以 nums[j] 结尾的j最长上升子序列；
// 所以如果能从 dp[j] 这个状态中转移过来，那么 nums[i] 必须大于 nums[j]
// 才能将 nums[i] 放在 nums[j] 之后形成更长的上升子序列
// 最后，整个数组的最长上升子序列即所有 dp[i] 中的最大值。
// dp[i] = max(dp[j]) + 1
// 0≤ j< i,nums[j] <nums [i]
func lengthOfLIS(nums []int) (ret int) {
	if nums == nil {
		return 0
	}
	ret = 1
	var dp []int
	dp = append(dp, 1)
	for i := 1; i < len(nums); i++ {
		dp = append(dp, 1)
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
			}
		}
		ret = int(math.Max(float64(ret), float64(dp[i])))
	}
	return
}

func lengthOfLIS2(nums []int) (sz int) {
	if nums == nil {
		return 0
	}
	sz = 1
	d := map[int]int{
		1: nums[0],
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] > d[sz] {
			sz += 1
			d[sz] = nums[i]
		} else {
			l, r, pos := 1, sz, 0
			for l <= r {
				//mid := (l + r) >> 1
				mid := (l + r) / 2
				if d[mid] < nums[i] {
					pos = mid
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
			d[pos+1] = nums[i]
		}
	}
	return
}

func main() {
	lengthOfLIS([]int{0, 1, 0, 3, 2, 3})
}
