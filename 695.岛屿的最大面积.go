/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 */

// @lc code=start
package leetcodego

func maxAreaOfIsland(grid [][]int) int {
	nr, nc := len(grid), len(grid[0])
	var dfs func(r, c int) int // 从grid[i][j]开始搜索的岛屿大小
	dfs = func(r, c int) int {
		if grid[r][c] == 0 {
			return 0
		}
		grid[r][c] = 0
		res := 1
		if r-1 >= 0 {
			res += dfs(r-1, c)
		}
		if r+1 < nr {
			res += dfs(r+1, c)
		}
		if c-1 >= 0 {
			res += dfs(r, c-1)
		}
		if c+1 < nc {
			res += dfs(r, c+1)
		}
		return res
	}
	ans := 0
	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			if grid[i][j] == 1 {
				ans = max(ans, dfs(i, j))
			}
		}
	}
	return ans
}

// @lc code=end
