/*
 * @lc app=leetcode.cn id=673 lang=golang
 *
 * [673] 最长递增子序列的个数
 */

// @lc code=start
package leetcodego

func findNumberOfLIS(nums []int) int {
	maxLen := 0
	n := len(nums)
	dp := make([]int, n) // dp[i]表示以 nums[i-1]结尾的递增子序列长度
	cnt := make([]int, n) // cnt[i]表示以 nums[i-1]结尾的递增子序列个数
	ans := 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		cnt[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j]+1
					cnt[i] = cnt[j]
				} else if dp[i] == dp[j]+1 {
					cnt[i] += cnt[j]
				}
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
			ans = cnt[i]
		} else if dp[i] == maxLen {
			ans += cnt[i]
		}
	}
	return ans
}

// @lc code=end
