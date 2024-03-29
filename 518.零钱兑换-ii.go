/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */

// @lc code=start
package leetcodego

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	// 先遍历 coins，保证不会出现重复顺序
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

// @lc code=end
