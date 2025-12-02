package mission1

import "strconv"

// 9. 回文数 https://leetcode.cn/problems/palindrome-number
func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	length := len(s)
	for i := 0; i < length/2; i++ {
		j := length - i - 1
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
