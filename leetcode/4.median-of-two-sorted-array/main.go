// link: https://leetcode-cn.com/problems/median-of-two-sorted-arrays

package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	res := make([]int, len(nums1)+len(nums2))
	for i, j, k := 0, 0, 0; i < len(nums1) || j < len(nums2); k++ {
		if i >= len(nums1) && j < len(nums2) || j < len(nums2) && nums2[j] <= nums1[i] {
			res[k] = nums2[j]
			j++
		} else {
			res[k] = nums1[i]
			i++
		}
	}
	if len(res)%2 != 0 {
		return float64(res[len(res)/2])
	} else {
		return (float64(res[len(res)/2-1] + res[len(res)/2])) / 2
	}
}

func main() {
	nums1, nums2 := []int{1, 2}, []int{3, 4}
	fmt.Print(findMedianSortedArrays(nums1, nums2))
}
