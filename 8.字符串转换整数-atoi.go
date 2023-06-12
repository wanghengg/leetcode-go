/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
package leetcodego

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.Trim(s, " ")
	if s == "" {
		return 0
	}
	var sign int64 = 1 // 符号，1-表示正数，2-表示负数
	first, second := 0, 0
	if s[0] == '+' {
		sign = 1
		first, second = 1, 1
	} else if s[0] == '-' {
		sign = -1
		first, second = 1, 1
	}
	length := len(s)
	for second < length {
		if s[second] >= '0' && s[second] <= '9' {
			second++
		} else {
			break
		}
	}
	numStr := s[first:second]
	var res int64
	for _, v := range numStr {
		res = res*10 + int64(v-'0')
		if sign*res < math.MinInt32 {
			return math.MinInt32
		} else if sign*res > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return int(sign * res)
}

// @lc code=end
