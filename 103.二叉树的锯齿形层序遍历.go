/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
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

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	levelNum := 0
	for len(q) > 0 {
		n := len(q)
		level := []int{}
		for i := 0; i < n; i++ {
			front := q[0]
			q = q[1:len(q)]
			if levelNum%2 == 0 {
				level = append(level, front.Val)
			} else {
				level = append([]int{front.Val}, level...)
			}
			if front.Left != nil {
				q = append(q, front.Left)
			}
			if front.Right != nil {
				q = append(q, front.Right)
			}
		}
		res = append(res, level)
		levelNum++
	}
	return res
}

// @lc code=end
