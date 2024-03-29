package util

// BuildPreSum 构建前缀和
// preSum[i] 为 nums[0..i-1] 的区间和
// nums[i, j] 的区间和为 preSum[j+1] - preSum[i]
//
// 例：nums[2, 3] 的区间和为 preSum[4] - preSum[2]
// 即 nums[0, 3] 的区间和减去 nums[0, 1] 的区间和
func BuildPreSum(nums []int) []int {
	preSum := make([]int, len(nums)+1)

	for i, v := range nums {
		preSum[i+1] = preSum[i] + v
	}

	return preSum
}

// RangeSumCountBad 求解 nums 中区间和为 k 的子 slice 个数
func RangeSumCountBad(nums []int, k int) int {
	preSum := BuildPreSum(nums)

	var cnt int
	for i := 1; i < len(preSum); i++ {
		for j := 0; j < i; j++ {
			// 区间 nums[j..i-1] 的和
			if preSum[i]-preSum[j] == k {
				cnt++
			}
		}
	}
	return cnt
}

// RangeSumCount 求解 nums 中区间和为 k 的子 slice 个数
func RangeSumCount(nums []int, k int) int {
	var cnt int

	// 保存前缀和与出现次数的映射
	ht := map[int]int{}

	// base case
	ht[0] = 1

	var preSum int
	for _, v := range nums {
		preSum += v
		if c, ok := ht[preSum-k]; ok {
			cnt += c
		}
		ht[preSum]++
	}
	return cnt
}
