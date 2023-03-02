/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package leetcodego

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(0, maxGain(node.Left))
		right := max(0, maxGain(node.Right))
		maxSum = max(maxSum, left+right+node.Val)
		return max(left, right) + node.Val
	}
	maxGain(root)
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
