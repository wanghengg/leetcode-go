/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
package leetcodego

func search1(nums []int, target int) int {
	var binSearch func(nums []int, left, right int) int
	binSearch = func(nums []int, left, right int) int {
		for left <= right {
			mid := (right-left)>>1 + left
			if nums[mid] == target {
				return mid
			} else if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return -1
	}
	return binSearch(nums, 0, len(nums)-1)
}

// @lc code=end
