/*
 * @lc app=leetcode.cn id=1026 lang=golang
 *
 * [1026] 节点与其祖先之间的最大差值
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

func maxAncestorDiff(root *TreeNode) int {
	res := 0
	var dfs func(node *TreeNode, maximum, minimum int)
	// maximum表示前面已经遍历过的祖先最大值，minimum表示已经遍历过祖先的最小值
	dfs = func(node *TreeNode, maximum, minimum int) {
		if node == nil {
			return
		}
		res = max(res, max(maximum-node.Val, node.Val-minimum))
		maximum = max(maximum, node.Val) // 已经遍历过的祖先的最大值
		minimum = min(minimum, node.Val) // 已经遍历过的祖先的最小值
		dfs(node.Left, maximum, minimum)
		dfs(node.Right, maximum, minimum)
	}
	dfs(root, root.Val, root.Val)
	return res
}

// @lc code=end
