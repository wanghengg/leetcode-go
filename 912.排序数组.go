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

// 堆排序
func sortArray2(nums []int) []int {
	var maxHeapify func(nums []int, i, heapSize int)
	maxHeapify = func(nums []int, i, heapSize int) {
		l, r, largest := 2*i+1, 2*i+2, i
		if l < heapSize && nums[l] > nums[largest] {
			largest = l
		}
		if r < heapSize && nums[r] > nums[largest] {
			largest = r
		}
		if largest != i {
			nums[i], nums[largest] = nums[largest], nums[i]
			maxHeapify(nums, largest, heapSize)
		}
	}
	buildMaxHeap := func(nums []int, heapSize int) {
		for i := heapSize / 2; i >= 0; i-- {
			maxHeapify(nums, i, heapSize)
		}
	}
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= 0; i-- {
		heapSize--
		nums[0], nums[i] = nums[i], nums[0]
		maxHeapify(nums, 0, heapSize)
	}
	return nums
}

// @lc code=end
