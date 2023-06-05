/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
package leetcodego

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var (
		n    = len(nums)
		best = math.MaxInt
	)
	update := func(cur int) {
		if abs(cur-target) < abs(best-target) {
			best = cur
		}
	}
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return target
			}
			update(sum)
			if sum > target {
				k--
			} else {
				j++
			}
		}
	}
	return best
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

// @lc code=end
