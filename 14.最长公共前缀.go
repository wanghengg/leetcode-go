/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
package leetcodego

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	s1 := strs[0]
	s2 := longestCommonPrefix(strs[1:])
	common := ""
	l1 := len(s1)
	l2 := len(s2)
	l := min(l1, l2)
	for i := 0; i < l; i++ {
		if s1[i] == s2[i] {
			common += string(s1[i])
		} else {
			break
		}
	}
	return common
}

// @lc code=end
