package sort

// 自顶向下归并排序
func MergeSort(arr []int) {
	tmp := make([]int, len(arr), len(arr))
	copyArray(arr, 0, len(arr)-1, tmp)
	internalMergeSort(arr, 0, len(arr)-1, tmp)
}

func internalMergeSort(arr []int, l, r int, tmp []int) {
	if l >= r {
		return
	}
	m := (r-l)/2 + l

	internalMergeSort(arr, l, m, tmp)
	internalMergeSort(arr, m+1, r, tmp)
	if arr[m] > arr[m+1] {
		merge(arr, l, m, r, tmp)
	}
}

func MergeSort2(arr []int) {
	tmp := make([]int, len(arr), len(arr))
	copyArray(arr, 0, len(arr)-1, tmp)
	internalMergeSort2(arr, 0, len(arr)-1, tmp)
}

// 排序个数 16 以内使用插入排序以提高性能
func internalMergeSort2(arr []int, l, r int, tmp []int) {
	if l >= r {
		return
	}

	if r-l+1 <= 16 {
		for i := l; i <= r; i++ {
			t := arr[i]
			j := i
			for ; j-1 >= l && t < arr[j-1]; j-- {
				arr[j] = arr[j-1]
			}
			arr[j] = t
		}
		return
	}

	m := (r-l)/2 + l

	internalMergeSort2(arr, l, m, tmp)
	internalMergeSort2(arr, m+1, r, tmp)
	if arr[m] > arr[m+1] {
		merge(arr, l, m, r, tmp)
	}
}

func merge(arr []int, l, m, r int, tmp []int) {
	copyArray(arr, l, r, tmp)

	i, j := l, m+1
	for k := l; i <= m || j <= r; k++ {
		if i > m {
			arr[k] = tmp[j]
			j++
		} else if j > r {
			arr[k] = tmp[i]
			i++
		} else if tmp[i] < tmp[j] {
			arr[k] = tmp[i]
			i++
		} else {
			arr[k] = tmp[j]
			j++
		}
	}
}

func copyArray(source []int, l, r int, dest []int) {
	for i := l; i <= r; i++ {
		dest[i] = source[i]
	}
}
