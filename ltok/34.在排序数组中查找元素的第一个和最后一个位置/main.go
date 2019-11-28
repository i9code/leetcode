package main

import "fmt"

//  O(logN)O(logN) 的时间复杂度，基本可以断定本题是需要使用二分查找
func searchRange(nums []int, target int) (ret []int) {
	nlen := len(nums)
	start := 0
	end := nlen - 1

	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			fmt.Println(mid)
			r1 := mid
			r2 := mid
			for r1 >= 0 && nums[r1] == target {
				r1 -= 1
			}
			for r2 < nlen && nums[r2] == target {
				r2 += 1
			}
			ret = append(ret, r1+1)
			ret = append(ret, r2-1)
			return ret
		}
		// target 在 前半部分
		if target >= nums[start] && target < nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	ret = append(ret, -1)
	ret = append(ret, -1)
	return ret
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(searchRange(arr, 1))
}

/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
*/
