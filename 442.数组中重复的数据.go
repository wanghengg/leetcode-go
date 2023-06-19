/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 */

// @lc code=start
package leetcodego

func findDuplicates(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	res := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			res = append(res, nums[i])
		}
	}
	return res
}

// @lc code=end
