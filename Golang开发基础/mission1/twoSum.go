package mission1

// 1. 两数之和 https://leetcode.cn/problems/two-sum/
func twoSum(nums []int, target int) []int {
	//方式一
	// var curTarget int
	// //双重循环，遍历求和
	// for i := 0; i < len(nums) - 1; i++ {
	//     curTarget = target - nums[i]
	//     for j := i + 1; j < len(nums); j++ {
	//         if nums[j] == curTarget {
	//             return []int{i,j}
	//         }
	//     }
	// }
	// return []int{}

	//方式二
	mp := make(map[int]int, len(nums))
	var curTarget int
	for i, num := range nums {
		curTarget = target - num
		if index, ok := mp[curTarget]; ok {
			return []int{i, index}
		} else {
			mp[num] = i
		}
	}
	return []int{}

}
