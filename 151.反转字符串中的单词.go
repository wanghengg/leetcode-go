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
	j = len(arr) - 1
	for i < j {
		tmp := arr[i]
		arr[i] = arr[j]
		arr[j] = tmp
		i++
		j--
	}
	return strings.Join(arr, " ")
}

/*
 * 现将整个字符串反转，然后再将每个单词反转
 */
func reverseWords1(s string) string {
	// 去除字符串中的多余空格
	arrStr := []byte(s)
	first, second := 0, 0
	for second < len(arrStr) {
		for second < len(arrStr) && arrStr[second] == ' ' {
			second++
		}
		if second > 0 && second < len(arrStr) && arrStr[second-1] == ' ' {
			if first != 0 {
				arrStr[first] = ' '
				first++
			}
		}
		if second >= len(arrStr) {
			break
		}
		arrStr[first] = arrStr[second]
		first++
		second++
	}
	arrStr = arrStr[0:first]
	// 将整个字符串反转
	for i, j := 0, len(arrStr)-1; i < j; i, j = i+1, j-1 {
		arrStr[i], arrStr[j] = arrStr[j], arrStr[i]
	}
	// 将每个单词反转
	first, second = 0, 0
	for ind := 0; ind < len(arrStr); ind++ {
		if arrStr[ind] != ' ' {
			continue
		}
		second = ind - 1
		for first < second {
			arrStr[first], arrStr[second] = arrStr[second], arrStr[first]
			first++
			second--
		}
		first = ind + 1
	}
	second = len(arrStr)-1
	for first < second {
		arrStr[first], arrStr[second] = arrStr[second], arrStr[first]
		first++
		second--
	}
	return string(arrStr)
}

// @lc code=end
