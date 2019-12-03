package main

import "sort"

import "fmt"

/*
	回溯法+排序剪枝
	为了对算法进行剪枝处理，首先对candidates排序

	特判，若candidates为空，则返回[][]

	回溯函数helper()helper()，传入参数：下一加和索引ii，当前已加和数组tmptmp，下一目标targettarget

	若target==0target==0，说明当前和满足条件，将当前加和数组tmptmp加入resres，并return。
	剪枝 因为已经将candidates排序，所以当下一目标小于下一待加和数时，return。并且当下一待加和索引i==ni==n时，return。为了防止数组越界，将条件i==ni==n放在target<candidates[i]target<candidates[i]之前，进行截断。
	因为可重复调用元素，所以helper(i,tmp+[candidates[i],target-candidates[i]])helper(i,tmp+[candidates[i],target−candidates[i]])，继续重复调用自身。
	调用数组中下一元素，寻找新答案。helper(i+1,tmp,target])helper(i+1,tmp,target])。
	执行helper(0,[],target)helper(0,[],target)，并返回resres

	复杂度分析
	时间复杂度：O(2^n)
	空间复杂度：O(1)
*/

func dfs(candidates []int, begin, target int, path []int, res [][]int) [][]int {
	// # 先写递归终止的情况
	if target == 0 {
		elm := []int{}
		elm = append(elm, path...)
		res = append(res, elm)
		// 	res = append(res, path)  直接这样子写会有问题  因为path的内容会被改变，导致 res内容被改变
		return res
	}

	for i := begin; i < len(candidates) && target-candidates[i] >= 0; i++ {
		path = append(path, candidates[i])
		res = dfs(candidates, i, target-candidates[i], path, res)
		path = path[:len(path)-1]
	}
	return res
}

func combinationSum(candidates []int, target int) (res [][]int) {
	nlen := len(candidates)
	if nlen == 0 {
		return
	}
	/*
		# 剪枝的前提是数组元素排序
		# 深度深的边不能比深度浅的边还小
		# 要排序的理由：1、前面用过的数后面不能再用；2、下一层边上的数不能小于上一层边上的数。
	*/
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	// # 在遍历的过程中记录路径，一般而言它是一个栈
	path := []int{}
	// # 注意要传入 size ，在 range 中， size 取不到
	return dfs(candidates, 0, target, path, res)
}

func main() {
	arr := []int{2, 3, 5}
	arr1 := combinationSum(arr, 8)
	fmt.Println(arr1)
}

/*
	给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

	candidates 中的数字可以无限制重复被选取。

	说明：

	所有数字（包括 target）都是正整数。
	解集不能包含重复的组合。
	示例 1:

	输入: candidates = [2,3,6,7], target = 7,
	所求解集为:
	[
		[7],
		[2,2,3]
	]
	示例 2:

	输入: candidates = [2,3,5], target = 8,
	所求解集为:
	[
	  [2,2,2,2],
	  [2,3,3],
	  [3,5]
	]

*/
