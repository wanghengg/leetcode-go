/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
package leetcodego

import "math/rand"

func findKthLargest(nums []int, k int) int {
	var partition func(nums []int, left, right, index int) int
	partition = func(nums []int, left, right, index int) int {
		pivotIndex := rand.Intn(right-left+1) + left
		pivot := nums[pivotIndex]
		nums[right], nums[pivotIndex] = nums[pivotIndex], nums[right]
		i, j := left, left
		for j < right {
			if nums[j] < pivot {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
			j++
		}
		nums[i], nums[right] = nums[right], nums[i]
		if i == index {
			return pivot
		} else if i < index {
			return partition(nums, i+1, right, index)
		} else {
			return partition(nums, left, i-1, index)
		}
	}
	return partition(nums, 0, len(nums)-1, len(nums)-k)
}

// @lc code=end
