package match

func RabinKarp(src string, target string) int {
	if src == "" || target == "" {
		panic("src & target can' be empty")
	}
	if len(target) > len(src) {
		return -1
	}

	const MOD, B = int64(1e9 + 7), int64(256)

	var tHash int64
	s, t := []rune(src), []rune(target)
	for i := 0; i < len(t); i++ {
		tHash = (tHash*B + int64(t[i])) % MOD
	}

	P := int64(1)
	for i := 0; i < len(t)-1; i++ {
		P = P * B % MOD
	}

	var hash int64
	for i := 0; i < len(t)-1; i++ {
		hash = (hash*B + int64(s[i])) % MOD
	}

	for i := len(t) - 1; i < len(s); i++ {
		hash = (hash*B + int64(s[i])) % MOD
		j := i - len(t) + 1
		if hash == tHash && matches(s[j:i+1], t) {
			return j
		}

		// int64(s[(i-len(t)+1)])*P%MOD 可能会大于 hash
		// hash 减去该值可能会负数，那就再加上一个 MOD 然后再取模
		hash = (hash - int64(s[j])*P%MOD + MOD) % MOD
	}
	return -1
}

func matches(a []rune, b []rune) bool {
	// return reflect.DeepEqual(a, b)
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
