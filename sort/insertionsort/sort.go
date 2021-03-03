package insertionsort

//func Sort(arr []int) {
//	for i := 0; i < len(arr); i++ {
//		for j := i; j-1 >= 0; j-- {
//			if arr[j] < arr[i] {
//				swap(arr, j, i)
//			} else {
//				break
//			}
//		}
//	}
//}

func Sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j-1 >= 0 && arr[j] < arr[j-1]; j-- {
			swap(arr, j, j-1)
		}
	}
}

func swap(arr []int, i, j int) {
	t := arr[i]
	arr[i] = arr[j]
	arr[j] = t
}

func Sort2(arr []int) {
	for i := 0; i < len(arr); i++ {
		t := arr[i]
		var j int
		// j := i cause bug because of variable scope problem
		// for j := i; j-1 >= 0 && t < arr[j-1]; j-- {
		for j = i; j-1 >= 0 && t < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = t
	}
}
