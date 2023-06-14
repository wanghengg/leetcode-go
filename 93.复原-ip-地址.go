/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 */

// @lc code=start
package leetcodego

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	res := []string{}
	var backtrace func(combine []string, ind int)
	backtrace = func(combine []string, ind int) {
		if ind == len(s) && len(combine) == 4 {
			res = append(res, strings.Join(combine, "."))
			return
		}
		if len(combine) > 4 || ind >= len(s) {
			return
		}
		if s[ind] == '0' {
			combine = append(combine, "0")
			backtrace(combine, ind+1)
		} else {
			for i := ind; i < len(s); i++ {
				subStr := string(s[ind : i+1])
				if value, _ := strconv.Atoi(subStr); value <= 255 {
					combine = append(combine, subStr)
					backtrace(combine, i+1)
				} else {
					break
				}
				combine = combine[:len(combine)-1]
			}
		}
	}
	backtrace([]string{}, 0)
	return res
}

// @lc code=end
