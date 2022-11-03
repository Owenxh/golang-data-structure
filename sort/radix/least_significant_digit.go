package radix

func LSDSort(arr []string, W int) {
	for i := 0; i < len(arr); i++ {
		if len(arr[i]) != W {
			panic("All string's length must be the same")
		}
	}

	// 定义一共有多少种取值的可能性
	R := 256

	// 定义每种可能性取值与其数量的对应关系
	cnt := make([]int, R)

	// 额外开辟空间用于排序
	temp := make([]string, len(arr))

	// 索引数组 cnt[i] 中元素对应的排序后的下标为 arr[index[i]..index[i+1])
	index := make([]int, R+1)

	for r := W - 1; r >= 0; r-- {
		// 循环开始之前，cnt 填充 0，清除上一次循环的运算结果
		for i := 0; i < R; i++ {
			cnt[i] = 0
		}

		// O(n)
		for _, s := range arr {
			j := []rune(s)[r]
			cnt[j]++
		}

		// O(R)
		for i := 0; i < R; i++ {
			index[i+1] = index[i] + cnt[i]
		}

		// O(n)
		for _, s := range arr {
			j := []rune(s)[r]

			temp[index[j]] = s
			index[j]++
		}
		copy(arr, temp)
	}
}
