/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
package leetcodego

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, v := range strs {
		chars := []rune(v)
		sort.Slice(chars, func(i, j int) bool {
			if chars[i] < chars[j] {
				return true
			} else {
				return false
			}
		})
		m[string(chars)] = append(m[string(chars)], v)
	}
	res := make([][]string, 0, len(m))
	for _, value := range m {
		res = append(res, value)
	}
	return res
}

// @lc code=end
