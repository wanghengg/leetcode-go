/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */

// @lc code=start
package leetcodego

import "strconv"

func countAndSay(n int) string {
	var dfs func(times int, s string) string
	dfs = func(times int, s string) string {
		if times == n {
			return s
		}
		res := ""
		first, second := 0, 0
		for second < len(s) {
			if s[second] == s[first] {
				second++
			} else {
				l := second - first
				res += strconv.Itoa(l)
				res += string(s[first])
				first = second
			}
		}
		l := second - first
		res += strconv.Itoa(l)
		res += string(s[first])
		return dfs(times+1, res)
	}
	return dfs(1, "1")
}

// @lc code=end
