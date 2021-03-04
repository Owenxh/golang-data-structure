package linearsearch

func SearchInt(data []int, target int) int {
	for i, datum := range data {
		if datum == target {
			return i
		}
	}
	return -1
}

func Search(data []interface{}, target interface{}) int {
	return SearchWith(
		func(a, b interface{}) bool {
			return a == b
		}, data, target)
}

func SearchWith(equals func(a, b interface{}) bool, data []interface{}, target interface{}) int {
	for i, datum := range data {
		if equals(datum, target) {
			return i
		}
	}
	return -1
}
