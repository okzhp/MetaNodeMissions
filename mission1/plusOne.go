package mission1

// 66. 加一 https://leetcode.cn/problems/plus-one/
func plusOne(digits []int) []int {
	length := len(digits)
	res := make([]int, length+1)

	var cur int
	var sum int
	var ret int = 1
	for i := length - 1; i >= 0; i-- {
		sum = digits[i] + ret
		ret = sum / 10
		cur = sum % 10
		res[i+1] = cur
	}
	if ret == 1 {
		res[0] = 1
	} else {
		res = res[1:]
	}
	return res
}
