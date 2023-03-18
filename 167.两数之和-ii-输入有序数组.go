/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
package leetcodego

// 注意这里给出的数组是有序的，所以可以使用双指针，当左右指针指向的和大于target时，right--，
// 小于target时，left++
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	res := []int{}
	for left < right {
		if numbers[left]+numbers[right] == target {
			res = append(res, left+1)
			res = append(res, right+1)
			break
		} else if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}
	}
	return res
}

// @lc code=end
