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

func findKthLargest1(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap(nums []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(nums, i, heapSize)
	}
}

func maxHeapify(nums []int, i, heapSize int) {
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

// @lc code=end
