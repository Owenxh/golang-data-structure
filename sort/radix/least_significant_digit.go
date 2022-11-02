package radix

func LSDSort(arr []string, W int) {
	for _, s := range arr {
		if len(s) != W {
			panic("All string's length must be the same")
		}
	}

	// TODO
}
