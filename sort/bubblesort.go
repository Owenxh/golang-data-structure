package sort

func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		var sorted bool
		for !sorted {
			sorted = true
			for j := 0; j+1 < len(arr)-i; j++ {
				if arr[j] > arr[j+1] {
					swap(arr, j, j+1)
					sorted = false
				}
			}
		}
	}
}
