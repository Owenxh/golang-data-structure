package radix

// MSDSort most significant digit
func MSDSort(arr []string) {
	// 额外开辟空间用于排序
	temp := make([]string, len(arr))
	internalMSDSort(arr, 0, len(arr)-1, 0, temp)
}

func internalMSDSort(arr []string, left int, right int, r int, temp []string) {
	if left >= right {
		return
	}

	// 定义一共有多少种取值的可能性
	R := 256

	// 定义每种可能性取值与其数量的对应关系
	// 需要考虑空串的情况
	cnt := make([]int, R+1)

	// 索引数组 cnt[i] 中元素对应的排序后的下标为 arr[index[i]..index[i+1])
	index := make([]int, R+2)

	// O(n)
	for i := left; i <= right; i++ {
		var j int
		// j >= len(arr[i]) 时表示字符串长度不够，视为空，算到 cnt 中第一个元素
		if r < len(arr[i]) {
			j = int(arr[i][r]) + 1
		}
		cnt[j]++
	}

	// O(R)
	for i := 0; i < R+1; i++ {
		index[i+1] = index[i] + cnt[i]
	}

	// O(n)
	for i := left; i <= right; i++ {
		var j int
		// j >= len(arr[i]) 时表示字符串长度不够，视为空，算到 cnt 中第一个元素
		if r < len(arr[i]) {
			j = int(arr[i][r]) + 1
		}

		temp[left+index[j]] = arr[i]
		index[j]++
	}

	// O(n)
	for i := left; i <= right; i++ {
		arr[i] = temp[i]
	}

	// index[1] 表示空字符串，从 index[2] 开始
	for i := 1; i <= R; i++ {
		internalMSDSort(arr, left+index[i], left+index[i+1]-1, r+1, temp)
	}
}
