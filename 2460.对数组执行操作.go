/*
 * @lc app=leetcode.cn id=2460 lang=golang
 *
 * [2460] 对数组执行操作
 */

// @lc code=start
package leetcodego

func applyOperations(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] = 2 * nums[i]
			nums[i+1] = 0
		}
	}
	for i, j := 0, 0; j < n; j++ {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	return nums
}

// @lc code=end
