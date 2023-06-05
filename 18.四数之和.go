/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
package leetcodego

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := [][]int{}
	if n < 4 {
		return res
	}
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, n-1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
					for l < n && nums[l] == nums[l-1] {
						l++
					}
					for r >= 0 && nums[r] == nums[r+1] {
						r++
					}
				} else if sum < target {
					l++
					for l < n && nums[l] == nums[l-1] {
						l++
					}
				} else {
					r--
					for r >= 0 && nums[r] == nums[r+1] {
						r--
					}
				}
			}
		}
	}
	return res
}

// @lc code=end
