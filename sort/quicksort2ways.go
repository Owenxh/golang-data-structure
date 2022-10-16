package sort

import "math/rand"

func QuickSort2Ways(arr []int) {
	quickSort2ways(arr, 0, len(arr)-1)
}

func quickSort2ways(arr []int, l, r int) {
	if l >= r {
		return
	}

	p := partition2ways(arr, l, r)
	quickSort2ways(arr, l, p-1)
	quickSort2ways(arr, p+1, r)
}

func partition2ways(arr []int, l, r int) int {
	k := l + rand.Intn(r-l+1)
	arr[l], arr[k] = arr[k], arr[l]

	v := arr[l]
	// 循环不变量：arr[l + 1...i-1] < v && arr[j+1...r] > v
	i, j := l+1, r
	for {
		for i <= j && arr[i] < v {
			i++
		}
		for j >= i && arr[j] > v {
			j--
		}
		if i >= j {
			break
		}

		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}

	arr[l], arr[j] = arr[j], arr[l]

	return j
}
