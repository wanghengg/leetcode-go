/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
package leetcodego

import "math"

func minSubArrayLen(target int, nums []int) int {
	preSum := make([]int, len(nums)+1)
	for ind, v := range nums {
		preSum[ind+1] = preSum[ind] + v
	}
	res := math.MaxInt
	for i, j := 0, 1; i < j && j <= len(nums); {
		for j <= len(nums) &&  preSum[j]-preSum[i] < target {
			j++
		}
		if j > len(nums) {
			break
		}
		for preSum[j]-preSum[i+1] >= target {
			i++
		}
		diff := j - i
		if diff < res {
			res = diff
		}
		j++
	}
	if res == math.MaxInt {
		return 0
	} else {
		return res
	}
}

// @lc code=end
