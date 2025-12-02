package mission1

import "sort"

// 56. 合并区间 https://leetcode.cn/problems/merge-intervals/
func merge(intervals [][]int) [][]int {
	//按区间左边界升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	//只有一个区间 直接返回
	if len(intervals) == 1 {
		return intervals
	}

	res := make([][]int, 0)

	i := 0
	for i < len(intervals)-1 {
		//逐个比较两个区间
		//如果能合并则继续向后合并，需要将合并结果放在后一个位置
		if mergeRes, ok := helper(intervals[i], intervals[i+1]); ok {
			intervals[i+1] = mergeRes
		} else {
			//不能合并则将前一个区间追加至结果集
			res = append(res, intervals[i])
		}
		i++
		//将最后一个区间追加至结果集
		if i >= len(intervals)-1 {
			res = append(res, intervals[i])
		}
	}
	return res
}

// 合并两个区间，其中a的左边界小于b的左边界
func helper(a, b []int) ([]int, bool) {
	if a[1] >= b[0] {
		right := a[1]
		if b[1] > right {
			right = b[1]
		}
		return []int{a[0], right}, true
	} else {
		return []int{}, false
	}
}
