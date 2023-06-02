/*
 * @lc app=leetcode.cn id=1768 lang=golang
 *
 * [1768] 交替合并字符串
 */

// @lc code=start
package leetcodego

func mergeAlternately(word1 string, word2 string) string {
	l1 := len(word1)
	l2 := len(word2)
	res := make([]byte, 0, l1+l2)
	var i, j int
	for i < l1 && j < l2 {
		res = append(res, word1[i], word2[j])
		i++
		j++
	}
	if i < l1 {
		res = append(res, word1[i:l1]...)
	}
	if j < l2 {
		res = append(res, word2[j:l2]...)
	}
	return string(res)
}

// @lc code=end
