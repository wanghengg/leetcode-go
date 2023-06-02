/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 反转字符串中的单词
 */

// @lc code=start
package leetcodego

import (
	"strings"
)

func reverseWords(s string) string {
	str := make([]byte, 0, len(s))
	l := len(s)
	isBegin := true
	for j := 0; j < l; j++ {
		if isBegin && s[j] == ' ' {
			continue
		}
		if s[j] != ' ' {
			str = append(str, s[j])
			isBegin = false
		} else if j+1 < l && s[j+1] != ' ' {
			str = append(str, s[j])
		}
	}
	arr := strings.Split(string(str), " ")
	var i, j int
	j = len(arr)-1
	for ; i < j; {
		tmp := arr[i]
		arr[i] = arr[j]
		arr[j] = tmp
		i++
		j--
	}
	return strings.Join(arr, " ")
}

// @lc code=end
