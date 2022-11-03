package radix

import "math"

func BucketSort(arr []int, B int) {
	if B <= 1 {
		panic("B must > 1")
	}

	// 临时的额外空间
	temp := make([]int, len(arr))
	bucketSort(arr, 0, len(arr)-1, B, temp)
}

// bucketSort 对 arr[left, right] 区间进行桶排序
func bucketSort(arr []int, left int, right int, B int, temp []int) {
	// 递归到底的情况
	if left >= right {
		return
	}

	min, max := math.MaxInt32, math.MinInt32
	for i := left; i <= right; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	if max == min {
		return
	}

	// 一共有多少个桶
	d := (max - min + 1) / B
	if (max-min+1)%B > 0 {
		d += 1
	}

	// TODO
}
