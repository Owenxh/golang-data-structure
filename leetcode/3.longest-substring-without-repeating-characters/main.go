// Link: https://leetcode-cn.com/problems/longest-substring-without-repeating-characters

package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	arr := []rune(s)
	index, res := 0, 0
	for i, c := range arr {
		sub := string(arr[index:i])
		if last := strings.LastIndex(sub, string(c)); last != -1 {
			index += last + 1
		} else if len(sub)+1 > res {
			res = len(sub) + 1
		}
	}
	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring("aaa"))
	fmt.Println(lengthOfLongestSubstring("abcdef"))
	fmt.Println(lengthOfLongestSubstring("abcabc"))
	fmt.Println(lengthOfLongestSubstring(""))
}
