/*
 * @lc app=leetcode.cn id=1022 lang=golang
 *
 * [1022] 从根到叶的二进制数之和
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

func sumRootToLeaf(root *TreeNode) int {
	sum := 0
	val := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node.Left == nil && node.Right == nil {
			val <<= 1
			val |= node.Val
			sum += val
			val >>= 1
			return
		}
		val <<= 1
		val |= node.Val
		if node.Left != nil {
			dfs(node.Left)
		}
		if node.Right != nil {
			dfs(node.Right)
		}
		val >>= 1
	}
	dfs(root)
	return sum
}

// @lc code=end
