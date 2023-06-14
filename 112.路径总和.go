/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
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

func hasPathSum(root *TreeNode, targetSum int) bool {
	var isExist bool
	var dfs func(root *TreeNode, targetSum int)
	dfs = func(root *TreeNode, curSum int) {
		if root == nil || isExist {
			return
		}
		curSum += root.Val
		if curSum == targetSum && root.Left == nil && root.Right == nil {
			isExist = true
			return
		}
		dfs(root.Left, curSum)
		dfs(root.Right, curSum)
	}
	dfs(root, 0)
	return isExist
}

// @lc code=end
