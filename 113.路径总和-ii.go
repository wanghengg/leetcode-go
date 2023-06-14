/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	combine := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, curSum int) {
		if node == nil {
			return
		}
		combine = append(combine, node.Val)
		curSum += node.Val
		if curSum == targetSum && node.Left == nil && node.Right == nil {
			res = append(res, append([]int{}, combine...))
			return
		}
		if node.Left == nil && node.Right == nil {
			return
		}
		if node.Left != nil {
			dfs(node.Left, curSum)
			combine = combine[:len(combine)-1]
		}
		if node.Right != nil {
			dfs(node.Right, curSum)
			combine = combine[:len(combine)-1]
		}
	}
	dfs(root, 0)
	return res
}

// @lc code=end
