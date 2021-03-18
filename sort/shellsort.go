package sort

func ShellSort(arr []int) {
	h := len(arr) / 2
	for h >= 1 {
		for k := 0; k < h; k++ {
			for i := k; i < len(arr); i += h {
				t, j := arr[i], i
				for ; j-h >= 0 && t < arr[j-h]; j -= h {
					arr[j] = arr[j-h]
				}
				arr[j] = t
			}
		}
		h /= 2
	}
}
