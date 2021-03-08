package sort

func QuickSort(arr []int) {
	internalQuickSort(arr, 0, len(arr)-1)
}

func internalQuickSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	p := partition(arr, l, r)
	internalQuickSort(arr, l, p-1)
	internalQuickSort(arr, p+1, r)
}

func partition(arr []int, l, r int) int {
	v := arr[l]
	j := l
	// 循环不变量：arr[l...j] < v && arr[j+1...i-1] >= v
	for i := l; i <= r; i++ {
		if arr[i] < v {
			j++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}
