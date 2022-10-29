package search

// BinarySearch 二分查找前提：数组中元素必须有序
func BinarySearch(arr []int, target int) int {
	return binarySearch(arr, 0, len(arr), target)
}

// 在 [l, r) 中搜索
func binarySearch(arr []int, l, r, target int) int {
	if l >= r {
		return -1
	}
	mid := (r-l)/2 + l
	if arr[mid] == target {
		return arr[mid]
	} else if arr[mid] > target {
		return binarySearch(arr, l, mid, target)
	} else {
		return binarySearch(arr, mid+1, r, target)
	}
}
