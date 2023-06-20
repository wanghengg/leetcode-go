/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
package leetcodego

func numIslands(grid [][]byte) int {
	var dfs func(row, col int)
	dfs = func(row, col int) {
		nrow, ncol := len(grid), len(grid[0])
		grid[row][col] = '0'
		if row-1 >= 0 && grid[row-1][col] == '1' {
			dfs(row-1, col)
		}
		if row+1 < nrow && grid[row+1][col] == '1' {
			dfs(row+1, col)
		}
		if col-1 >= 0 && grid[row][col-1] == '1' {
			dfs(row, col-1)
		}
		if col+1 < ncol && grid[row][col+1] == '1' {
			dfs(row, col+1)
		}
	}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res += 1
				dfs(i, j)
			}
		}
	}
	return res
}

// @lc code=end
