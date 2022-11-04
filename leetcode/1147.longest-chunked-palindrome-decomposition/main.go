package main

import "fmt"

func longestDecomposition(text string) int {
	return solve([]rune(text), 0, len(text)-1)
}

func solve(s []rune, left, right int) int {
	if left > right {
		return 0
	}
	i, j := left, right
	for i < j {
		if matches(s[left:i+1], s[j:right+1]) {
			return 2 + solve(s, i+1, j-1)
		}
		i++
		j--
	}
	return 1
}

func matches(a []rune, b []rune) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(longestDecomposition("ghiabcdefhelloadamhelloabcdefghi"))
}
