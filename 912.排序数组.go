/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

// @lc code=start
package leetcodego

import (
	"math/rand"
)

// 快速排序
func sortArray(nums []int) []int {
	var quickSort func(nums []int, left, right int)
	quickSort = func(nums []int, left, right int) {
		if left >= right {
			return
		}
		pivotIndex := rand.Intn(right-left+1) + left
		pivot := nums[pivotIndex]
		nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]
		i := left
		for j := left; j < right; j++ {
			if nums[j] < pivot {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
		}
		nums[i], nums[right] = nums[right], nums[i]
		quickSort(nums, left, i-1)
		quickSort(nums, i+1, right)
	}
	quickSort(nums, 0, len(nums)-1)
	return nums
}

// 归并排序
func sortArray1(nums []int) []int {
	var mergeSort func(nums []int) []int
	var merge func(left, right []int) []int
	mergeSort = func(nums []int) []int {
		if len(nums) == 1 {
			return nums
		}
		mid := len(nums) / 2
		left := mergeSort(nums[:mid])
		right := mergeSort(nums[mid:])
		return merge(left, right)
	}
	merge = func(left, right []int) []int {
		res := make([]int, 0, len(left)+len(right))
		i, j := 0, 0
		for i < len(left) && j < len(right) {
			if left[i] < right[j] {
				res = append(res, left[i])
				i++
			} else {
				res = append(res, right[j])
				j++
			}
		}
		res = append(res, left[i:]...)
		res = append(res, right[j:]...)
		return res
	}
	return mergeSort(nums)
}

// @lc code=end
