/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */

// @lc code=start
package leetcodego

// 数组nums长度为n+1，且每个元素都处于[1,n]区间内，那么遍历数组，假设nums[i]的值为k，则应该
// 将其放在nums[k]处，所以将nums[i]与nums[k]交换，最终总会遇到想将某个数放至nums[k]时，
// 该位置已经有了等于k的元素，说明这个元素重复了
func findDuplicate(nums []int) int {
	for {
		if nums[0] == nums[nums[0]] {
			return nums[0]
		} else {
			nums[0], nums[nums[0]] = nums[nums[0]], nums[0]
		}
	}
	return 0
}

// @lc code=end
