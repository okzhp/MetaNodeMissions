package mission1

import "math"

// 14. 最长公共前缀 https://leetcode.cn/problems/longest-common-prefix/description/
func longestCommonPrefix(strs []string) string {
	var index int = math.MaxInt
	//找到最短字符串长度
	for _, s := range strs {
		if len(s) < index {
			index = len(s)
		}
	}
	var done bool
	var maxLen int

	for i := 0; i < index; i++ {
		var c uint8
		//遍历字符串数组 并比较第i个字符是否一致
		for _, str := range strs {
			if c == 0 {
				c = str[i]
			} else if c != str[i] {
				done = true
				break
			}
		}
		if done {
			break
		}
		maxLen++
	}

	return strs[0][:maxLen]
}
