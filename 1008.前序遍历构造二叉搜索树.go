/*
 * @lc app=leetcode.cn id=1008 lang=golang
 *
 * [1008] 前序遍历构造二叉搜索树
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

/*
二叉搜索树前序遍历其左子树的所有节点值都严格小于当前节点的值，右子树所有节点的值严格大于当前节点
的值，根据这个特点可以判断左子树和右子树的边界
*/
func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	index := 0
	rootVal := preorder[0]
	for index < len(preorder) && preorder[index] <= rootVal {
		index++
	}
	leftPreOrder := preorder[1:index]
	rightPreOrder := preorder[index:]
	root := &TreeNode{
		Val:   rootVal,
		Left:  bstFromPreorder(leftPreOrder),
		Right: bstFromPreorder(rightPreOrder),
	}
	return root
}

// @lc code=end
