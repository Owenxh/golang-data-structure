package sort

func QuickSort1Way(arr []int) {
	quickSort2ways(arr, 0, len(arr)-1)
}

func quickSort1Way(arr []int, l, r int) {
	if l >= r {
		return
	}

	p := partition1way(arr, l, r)
	quickSort1Way(arr, l, p-1)
	quickSort1Way(arr, p+1, r)
}

func partition1way(arr []int, l, r int) int {
	// 循环不变量 v = arr[l]; arr[l+1...j] < v && arr[j+1...i-1] > v;  [i..r] 还没有遍历的

	v := arr[l]
	j := l

	for i := l; i <= r; i++ {
		if arr[i] < v {
			j++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}

	arr[j], arr[l] = arr[l], arr[j]
	return j
}
