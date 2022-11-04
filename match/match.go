package match

func BruteForce(src string, t string) int {
	if src == "" || t == "" {
		panic("src & t can't be empty")
	}
	if len(t) > len(src) {
		return -1
	}

	source, target := []rune(src), []rune(t)
	for i := 0; i+len(target) <= len(source); i++ {
		j := 0
		for j < len(target) {
			if target[j] != source[j+i] {
				break
			}
			j++
		}
		if j == len(target) {
			return i
		}
	}
	return -1
}
