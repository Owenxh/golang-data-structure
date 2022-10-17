package sort

func SelectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[i] {
				swap(arr, j, i)
			}
		}
	}
}

func SelectSort2(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			if arr[j] > arr[i] {
				swap(arr, j, i)
			}
		}
	}
}
