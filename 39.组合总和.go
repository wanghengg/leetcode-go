/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
package leetcodego

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	var dfs func(arr []int, index, sum int)
	dfs = func(arr []int, index, sum int) {
		arr = append(arr, candidates[index])
		sum += candidates[index]
		if sum == target {
			res = append(res, append([]int{}, arr...))
			return
		} else if sum > target {
			return
		}
		for i := index; i < len(candidates); i++ {
			dfs(arr, i, sum)
		}
	}
	for i := 0; i < len(candidates); i++ {
		dfs([]int{}, i, 0)
	}
	return res
}

// @lc code=end
