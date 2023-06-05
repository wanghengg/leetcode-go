/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
package leetcodego

import "sort"

/*
 * 先排序，然后使用双指针
 */
func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	n := len(nums)
	res := [][]int{}
	for i := 0; i < n-2; i++ {
		num1 := nums[i]
		if i > 0 && num1 == nums[i-1] {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			tmp := nums[j] + nums[k]
			if tmp == -num1 {
				res = append(res, []int{num1, nums[j], nums[k]})
				j++
				for j < n && nums[j] == nums[j-1] {
					j++
				}
				k--
				for k > j && nums[k] == nums[k+1] {
					k--
				}
			} else if tmp > -num1 {
				k--
			} else {
				j++
			}
		}
	}
	return res
}

func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := [][]int{}
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < n && nums[j] == nums[j-1] {
					j++
				}
				for k >= 0 && nums[k] == nums[k+1] {
					k--
				}
			} else if sum < 0 {
				j++
				for j < n && nums[j] == nums[j-1] {
					j++
				}
			} else {
				k--
				for k >= 0 && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}
	return res
}

// @lc code=end
