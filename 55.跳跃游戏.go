/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
package leetcodego

func canJump(nums []int) bool {
	n := len(nums)
	rightMost := 0 // 能到达的最右位置
	for i := 0; i < n; i++ {
		if i <= rightMost { // 当前位置小于等于 rightMost 时，更新 rightMost
			rightMost = max(rightMost, i+nums[i])
		}
		if rightMost >= n-1 {
			return true
		}
	}
	return false
}

// @lc code=end
