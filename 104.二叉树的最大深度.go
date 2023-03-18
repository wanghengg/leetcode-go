/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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

import "container/list"

func maxDepth(root *TreeNode) int {
	res := 0
	depth := 0
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			res = max(res, depth)
			return
		}
		depth++
		traverse(node.Left)
		traverse(node.Right)
		depth--
	}
	traverse(root)
	return res
}

// 深度优先搜索
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 二叉树的最大深度=max(左子树的最大深度，右子树的最大深度) + 1
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 广度优先搜索
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	l := list.New()
	l.PushBack(root)
	for l.Len() > 0 {
		eleNum := l.Len()
		for i := eleNum; i > 0; i-- {
			element := l.Front()
			node := element.Value.(*TreeNode)
			if node.Left != nil {
				l.PushBack(node.Left)
			}
			if node.Right != nil {
				l.PushBack(node.Right)
			}
			l.Remove(element)
		}
		res++
	}
	return res
}

// @lc code=end
