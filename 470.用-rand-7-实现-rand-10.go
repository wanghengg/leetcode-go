/*
 * @lc app=leetcode.cn id=470 lang=golang
 *
 * [470] 用 Rand7() 实现 Rand10()
 */

// @lc code=start
package leetcodego

func rand10() int {
	first, second := rand7(), rand7()
	for first > 6 {
		first = rand7()
	}
	for second > 5 {
		second = rand7()
	}
	if first&1 == 1 {
		return second
	} else {
		return 5 + second
	}
}

func rand7() int {
	return 0
}

// @lc code=end
