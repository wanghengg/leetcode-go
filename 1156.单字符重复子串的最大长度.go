/*
 * @lc app=leetcode.cn id=1156 lang=golang
 *
 * [1156] 单字符重复子串的最大长度
 */

// @lc code=start
package leetcodego

func maxRepOpt1(text string) int {
	count := make(map[rune]int)
	for _, c := range text {
		count[c]++
	}
	res, i := 0, 0
	n := len(text)
	for i < n {
		// 找出一段连续的[i, j)，这段全是一样的字符text[i]，长度为j-i
		j := i
		for j < n && text[j] == text[i] {
			j++
		}
		l := j - i
		k := j + 1
		// 找出一段连续的[j+1, k)，这段字符也全是text[i]，长度为k-j-1
		for k < n && text[k] == text[i] {
			k++
		}
		r := k - j - 1
		// 长度最长为count[rune(text[i])]
		res = max(res, min(l+r+1, count[rune(text[i])]))
		i = j // 下一个不同的字符
	}
	return res
}

// @lc code=end
