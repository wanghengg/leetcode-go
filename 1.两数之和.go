/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
package leetcodego

func twoSumII(nums []int, target int) []int {
	m := make(map[int]int)
	for ind, elem := range nums {
		i, ok := m[target-elem]
		if ok {
			return []int{i, ind}
		} else {
			m[elem] = ind
		}
	}
	return []int{}
}

// @lc code=end
