/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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

func isSymmetric(root *TreeNode) bool {
	var dfs func(left, right *TreeNode) bool
	dfs = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil && right != nil {
			return false
		}
		if left != nil && right == nil {
			return false
		}
		return left.Val == right.Val && dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
	}
	return dfs(root.Left, root.Right)
}

// @lc code=end
