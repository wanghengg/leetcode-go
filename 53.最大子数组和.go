/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
package leetcodego

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	ans := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		ans = max(ans, dp[i])
	}
	return ans
}

func maxSubArray1(nums []int) int {
	cur := nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		cur = max(cur+nums[i], nums[i])
		ans = max(ans, cur)
	}
	return ans
}

// @lc code=end
