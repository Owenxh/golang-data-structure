package sort

import "math/rand"

func QuickSort3Ways(arr []int) {
	quickSort3Ways(arr, 0, len(arr)-1)
}

func quickSort3Ways(arr []int, l, r int) {
	if l >= r {
		return
	}

	k := l + rand.Intn(r-l+1)
	arr[l], arr[k] = arr[k], arr[l]
	v := arr[l]

	// 循环不变量
	// ① [l+1, lt] < v
	// ② [lt + 1, i - 1] == v
	// ③ [gt, r] > v
	lt, gt := l-1, r+1
	for i := l + 1; i <= r && i < gt; {
		if arr[i] < v {
			lt++
			arr[i], arr[lt] = arr[lt], arr[i]
			i++
		} else if arr[i] > v {
			gt--
			arr[i], arr[gt] = arr[gt], arr[i]
		} else {
			i++
		}
	}
	// 循环之后
	// ① [l+1, lt] < v
	// ② [lt + 1, gt - 1] == v
	// ③ [gt, r] > v

	quickSort3Ways(arr, l, lt)
	quickSort3Ways(arr, gt, r)
}
