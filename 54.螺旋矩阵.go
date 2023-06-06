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

func spiralOrder1(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows, columns := len(matrix), len(matrix[0])
	res := make([]int, 0, rows*columns)
	left, right, top, bottom := 0, columns-1, 0, rows-1
	for left <= right && top <= bottom {
		for column := left; column <= right; column++ {
			res = append(res, matrix[top][column])
		}
		for row := top + 1; row <= bottom; row++ {
			res = append(res, matrix[row][right])
		}
		if left < right && top < bottom { // 最后只有一行或者一列，不能继续遍历
			for column := right - 1; column >= left; column-- {
				res = append(res, matrix[bottom][column])
			}
			for row := bottom - 1; row > top; row-- {
				res = append(res, matrix[row][left])
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return res
}

// @lc code=end
