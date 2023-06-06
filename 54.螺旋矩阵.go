/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
package leetcodego

// 时间复杂度O(m*n)，空间复杂度O(m*n)
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	direction := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	seen := make([][]bool, m)
	for ind := range seen {
		seen[ind] = make([]bool, n)
	}
	i, j := 0, 0
	res := make([]int, 0, m*n)
	d := 0
	for i >= 0 && i < m && j >= 0 && j < n && !seen[i][j] {
		res = append(res, matrix[i][j])
		seen[i][j] = true
		if i+direction[d][0] >= m || i+direction[d][0] < 0 ||
			j+direction[d][1] >= n || j+direction[d][1] < 0 ||
			seen[i+direction[d][0]][j+direction[d][1]] {
			d += 1
			d %= 4
		}
		i += direction[d][0]
		j += direction[d][1]
	}
	return res
}

// @lc code=end
