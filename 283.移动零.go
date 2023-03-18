/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
package leetcodego

func moveZeroes(nums []int) {
	left := 0
	right := 0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	for left < len(nums) {
		nums[left] = 0
		left++
	}
	return
}

// @lc code=end
