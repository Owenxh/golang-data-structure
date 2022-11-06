package match

func BruteForce(src string, target string) int {
	if src == "" || target == "" {
		panic("src & target can't be empty")
	}
	if len(target) > len(src) {
		return -1
	}

	s, t := []rune(src), []rune(target)
	for i := 0; i+len(t) <= len(s); i++ {
		j := 0
		for j < len(t) {
			if t[j] != s[j+i] {
				break
			}
			j++
		}
		if j == len(t) {
			return i
		}
	}
	return -1
}
