/*
 * @lc app=leetcode.cn id=2559 lang=golang
 *
 * [2559] 统计范围内的元音字符串数
 */

// @lc code=start
package leetcodego

func vowelStrings(words []string, queries [][]int) []int {
	res := make([]int, 0, len(queries))
	preSum := make([]int, len(words)+1)
	vowel := map[uint8]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	for ind, word := range words {
		if vowel[word[0]] && vowel[word[len(word)-1]] {
			preSum[ind+1] = preSum[ind] + 1
		} else {
			preSum[ind+1] = preSum[ind]
		}
	}
	for _, v := range queries {
		res = append(res, preSum[v[1]+1]-preSum[v[0]])
	}
	return res
}

// @lc code=end
