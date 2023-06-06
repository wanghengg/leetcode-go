/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
package leetcodego

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for ind := range matrix {
		matrix[ind] = make([]int, n)
	}
	count := 1
	left, right, top, bottom := 0, n-1, 0, n-1
	for left <= right && top <= bottom {
		for column := left; column <= right; column++ {
			matrix[top][column] = count
			count++
		}
		for row := top + 1; row <= bottom; row++ {
			matrix[row][right] = count
			count++
		}
		for column := right - 1; column >= left; column-- {
			matrix[bottom][column] = count
			count++
		}
		for row := bottom - 1; row > top; row-- {
			matrix[row][left] = count
			count++
		}
		left++
		right--
		top++
		bottom--
	}
	return matrix
}

// @lc code=end
