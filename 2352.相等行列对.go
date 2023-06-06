/*
 * @lc app=leetcode.cn id=2352 lang=golang
 *
 * [2352] 相等行列对
 */

// @lc code=start
package leetcodego

import "strings"

func equalPairs(grid [][]int) int {
	rows := make(map[string]int)
	columns := make(map[string]int)
	n := len(grid)
	for i := 0; i < n; i++ {
		tmp := make([]string, 0, n)
		for j := 0; j < n; j++ {
			tmp = append(tmp, string(grid[i][j]))
		}
		rows[strings.Join(tmp, "_")]++
	}
	for j := 0; j < n; j++ {
		tmp := make([]string, 0, n)
		for i := 0; i < n; i++ {
			tmp = append(tmp, string(grid[i][j]))
		}
		columns[strings.Join(tmp, "_")]++
	}
	res := 0
	for k, v := range rows {
		res += v * columns[k]
	}
	return res
}

// @lc code=end
