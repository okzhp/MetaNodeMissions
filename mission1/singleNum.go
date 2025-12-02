package mission1

// 136. 只出现一次的数字 https://leetcode.cn/problems/single-number
func singleNumber(nums []int) int {
	//计数map
	cntMap := make(map[int]int, len(nums)/2+1)
	for _, num := range nums {
		//已出现一次的key 直接删除
		if _, ok := cntMap[num]; ok {
			delete(cntMap, num)
		} else {
			cntMap[num] = 1
		}
	}

	var result int
	//最终map只剩一个出现一次的元素
	for k, _ := range cntMap {
		result = k
		break
	}
	return result
}
