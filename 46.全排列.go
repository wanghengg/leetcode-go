/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
package leetcodego

func permute(nums []int) [][]int {
	res := [][]int{}
	combine := []int{}
	seen := make(map[int]bool)
	var dfs func()
	dfs = func() {
		if len(combine) == len(nums) {
			res = append(res, append([]int{}, combine...))
			return
		}
		for k, v := range nums {
			if seen[k] {
				continue
			}
			combine = append(combine, v)
			seen[k] = true
			dfs()
			combine = combine[:len(combine)-1]
			delete(seen, k)
		}
	}
	dfs()
	return res
}

// @lc code=end
