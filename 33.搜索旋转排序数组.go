/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
package leetcodego

import "sort"

func search(nums []int, target int) int {
	index := -1
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			index = i
		}
	}
	if index == -1 {
		loc := sort.Search(n, func(i int) bool {
			return nums[i] >= target
		})
		if loc == n {
			return -1
		} else if nums[loc] == target {
			return loc
		} else {
			return -1
		}
	}
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			if mid <= index {
				if target >= nums[0] {
					right = mid - 1
				} else {
					left = index + 1
				}
			} else {
				left = index + 1
				right = mid - 1
			}
		} else {
			if mid <= index {
				left = mid + 1
				right = index
			} else {
				if target >= nums[0] {
					right = index
				} else {
					left = mid + 1
				}
			}
		}
	}
	return -1
}

// @lc code=end
