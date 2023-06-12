/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 找出字符串中第一个匹配项的下标
 */

// @lc code=start
package leetcodego

func strStr(haystack string, needle string) int {
	i := 0
	l1 := len(haystack)
	l2 := len(needle)
	for ; i <= l1-l2; i++ {
		j := i
		for k := 0; k < l2; k++ {
			if haystack[j] == needle[k] {
				j++
			} else {
				break
			}
		}
		if j-i == l2 {
			return i
		}
	}
	return -1
}

// @lc code=end
