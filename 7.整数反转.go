/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
package leetcodego

import "math"

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}
	res := 0
	for x != 0 {
		bit := x % 10
		res = res*10 + bit
		x /= 10
		if x != 0 && (sign*res < math.MinInt32/10 || sign*res > math.MaxInt32/10) {
			return 0
		}
	}
	return sign * res
}

// @lc code=end
