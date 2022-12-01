package util

// GCD 最大公约数
func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

// LCM 最小公倍数
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}
