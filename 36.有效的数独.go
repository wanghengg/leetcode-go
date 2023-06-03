/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
package leetcodego

func isValidSudoku(board [][]byte) bool {
	row := make([]map[byte]bool, 9)
	column := make([]map[byte]bool, 9)
	grid := make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		row[i] = make(map[byte]bool)
		column[i] = make(map[byte]bool)
		grid[i] = make(map[byte]bool)
	}
	for i, line := range board {
		for j, c := range line {
			if c == '.' {
				continue
			}
			if row[i][c] {
				return false
			} else {
				row[i][c] = true
			}
			if column[j][c] {
				return false
			} else {
				column[j][c] = true
			}
			ind := i/3*3 + j/3
			if grid[ind][c] {
				return false
			} else {
				grid[ind][c] = true
			}
		}
	}
	return true
}

// @lc code=end
