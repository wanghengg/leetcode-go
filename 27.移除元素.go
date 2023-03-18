/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
package leetcodego

func removeElement(nums []int, val int) int {
	left := 0
	right := 0
	for right < len(nums) {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}

// @lc code=end
