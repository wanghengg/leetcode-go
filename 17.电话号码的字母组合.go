/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
package leetcodego

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	m := map[rune][]rune{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	res := []string{}
	var dfs func(digits string, ind int, combine string)
	dfs = func(digits string, ind int, combine string) {
		if ind == len(digits) {
			res = append(res, combine)
			return
		}
		for _, c := range m[rune(digits[ind])] {
			combine += string(c)
			dfs(digits, ind+1, combine)
			combine = combine[:len(combine)-1]
		}
	}
	dfs(digits, 0, "")
	return res
}

// @lc code=end
