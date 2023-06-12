/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
package leetcodego

func subsets(nums []int) [][]int {
	res := [][]int{}
	combine := []int{}
	var dfs func(ind int)
	dfs = func(ind int) {
		res = append(res, append([]int{}, combine...))
		if ind >= len(nums) {
			return
		}
		for i := ind; i < len(nums); i++ {
			combine = append(combine, nums[i])
			dfs(i + 1)
			combine = combine[:len(combine)-1]
		}
	}
	dfs(0)
	return res
}

// @lc code=end
