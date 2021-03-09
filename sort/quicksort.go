package sort

import "math/rand"

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	p := partition(arr, l, r)
	quickSort(arr, l, p-1)
	quickSort(arr, p+1, r)
}

func partition(arr []int, l, r int) int {
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

//func partition(arr []int, l, r int) int {
//	v := arr[l]
//	j := l
//	// 循环不变量：arr[l...j] < v && arr[j+1...i-1] >= v
//	for i := l; i <= r; i++ {
//		if arr[i] < v {
//			j++
//			arr[i], arr[j] = arr[j], arr[i]
//		}
//	}
//	arr[l], arr[j] = arr[j], arr[l]
//	return j
//}
