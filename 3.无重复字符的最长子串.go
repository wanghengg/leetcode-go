/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
package leetcodego

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]struct{})
	res := 0
	left := 0
	right := 0
	for right < len(s) {
		if _, ok := m[s[right]]; !ok {
			m[s[right]] = struct{}{}
			right++
		} else {
			res = max(res, right-left)
			for s[left] != s[right] {
				delete(m, s[left])
				left++
			}
			left++
			right++
		}
	}
	return max(res, right-left)
}

// @lc code=end
