package mission1

// 20. 有效的括号 https://leetcode.cn/problems/valid-parentheses/
func isValid(s string) bool {
	coupleMap := map[rune]rune{'(': ')', '{': '}', '[': ']'}
	stack := make([]rune, 0)
	for _, value := range s {
		//栈为空 或者 栈顶括号和当前遍历括号不匹配 进行入栈
		if len(stack) == 0 || value != coupleMap[stack[len(stack)-1]] {
			stack = append(stack, value)
		} else {
			//出栈
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
