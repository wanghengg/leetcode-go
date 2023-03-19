/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
package leetcodego

func longestPalindrome(s string) string {
	res := ""
	length := len(s)
	for i := 0; i < length; i++ {
		s1 := palindrome(s, i, i)
		s2 := palindrome(s, i, i+1)
		if len(res) < len(s1) {
			res = s1
		}
		if len(res) < len(s2) {
			res = s2
		}
	}
	return res
}

func palindrome(s string, l, r int) string {
	length := len(s)
	for l >= 0 && r < length && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}

// 动态规划
func longestPalindrome1(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}
	maxLength := 1
	begin := 0
	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
		dp[i][i] = true
	}
	// 先枚举子串长度
	for l := 2; l <= length; l++ {
		// 再枚举左边界
		for i := 0; i < length; i++ {
			j := l + i - 1
			if j >= length {
				break
			}
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j == i+1 { // 这里是针对s[i]==s[i+1]的情况
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && j-i+1 > maxLength {
				maxLength = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLength]
}

// @lc code=end
