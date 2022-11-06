package match

func GetLPSBytes(s string) []int {
	if len(s) == 0 {
		return nil
	}

	lps := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		// 最长的 border
		a := lps[i-1]
		for a > 0 && s[i] != s[a] {
			// 找次长的 border
			a = lps[a-1]
		}
		if s[i] == s[a] {
			lps[i] = a + 1
		}
	}

	return lps
}

func GetLPS(src string) []int {
	if len(src) == 0 {
		return nil
	}

	s := []rune(src)

	lps := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		// 最长的 border
		a := lps[i-1]
		for a > 0 && s[i] != s[a] {
			// 找次长的 border
			a = lps[a-1]
		}
		if s[i] == s[a] {
			lps[i] = a + 1
		}
	}

	return lps
}

// Knuth-Morris-Pratt
// KMP match
func KMP(s string, t string) int {
	if s == "" || t == "" {
		panic("src & target can't be empty")
	}

	if len(s) < len(t) {
		return -1
	}

	lps := GetLPS(t)
	return kmp0([]rune(s), []rune(t), lps)
}

func kmp0(s []rune, t []rune, lps []int) int {
	var i, j int

	// src 字符串下标 i 不往回退
	for i < len(s) {
		// 字符匹配，s 和 t 的下标向前移动
		if s[i] == t[j] {
			i++
			j++

			// 目标字符串 t 移动到末尾，字符串匹配
			if j == len(t) {
				return i - len(t)
			}
		} else if j > 0 {
			// lps[j-1] 为目标字符串 t 从 [0..j-1] 的最长的 border 的长度；
			// 该 border 对应字符串 t 的索引范围是 [0, lps[j-1] - 1]；它的下一个索引是 lps[j-1]
			j = lps[j-1]
		} else {
			// j == 0；从原字符串 s 的下一个坐标开始匹配
			i++
		}
	}

	return -1
}

// KMPUtf8Bad 处理中文时返回的下标非目标字符串在字符串中的下标位置
// 性能接近 golang 标准库
func KMPUtf8Bad(s string, t string) int {
	if s == "" || t == "" {
		panic("s & t can't be empty")
	}

	if len(s) < len(t) {
		return -1
	}

	lps := GetLPSBytes(t)

	var i, j int

	for i < len(s) {
		if s[i] == t[j] {
			i++
			j++
			if j == len(t) {
				return i - len(t)
			}
		} else if j > 0 {
			j = lps[j-1]
		} else {
			i++
		}
	}

	return -1
}

func KMP2(s string, t string) int {
	if s == "" || t == "" {
		panic("s & t can't be empty")
	}

	if len(s) < len(t) {
		return -1
	}

	lps := GetLPSBytes(t)

	var i, j int

	for i < len(s) {
		if s[i] == t[j] {
			i++
			j++
			if j == len(t) {
				index := i - len(t)
				return len([]rune(s[:index]))
			}
		} else if j > 0 {
			j = lps[j-1]
		} else {
			i++
		}
	}

	return -1
}
