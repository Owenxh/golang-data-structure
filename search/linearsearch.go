package search

func LinearSearchInt(data []int, target int) int {
	for i, datum := range data {
		if datum == target {
			return i
		}
	}
	return -1
}

func LinearSearch(data []interface{}, target interface{}) int {
	return LinearSearchWith(
		func(a, b interface{}) bool {
			return a == b
		}, data, target)
}

func LinearSearchWith(equals func(a, b interface{}) bool, data []interface{}, target interface{}) int {
	for i, datum := range data {
		if equals(datum, target) {
			return i
		}
	}
	return -1
}
