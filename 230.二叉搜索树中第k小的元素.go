/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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

// 因为二叉搜索树的中序遍历是有序的，所以找第k大的元素很容易想到应该使用中序遍历做
func kthSmallest(root *TreeNode, k int) int {
	res := 0
	rank := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		rank++
		if rank == k {
			res = node.Val
			return
		}
		dfs(node.Right)
	}
	dfs(root)
	return res
}

// 方法2
func kthSmallest1(root *TreeNode, k int) int {
	arr := []int{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		arr = append(arr, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return arr[k-1]
}

// @lc code=end
