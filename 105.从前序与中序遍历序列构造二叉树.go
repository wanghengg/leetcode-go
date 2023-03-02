/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func build(preOrder []int, inOrder []int, preStart int, preEnd int,
	inStart int, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	rootVal := preOrder[preStart]
	// 构造根节点
	root := &TreeNode{
		Val: rootVal,
	}
	index := 0
	// 注意这里i可以等于inend，因为inEnd表示中序遍历的最后一个元素的索引
	// 这里从中序遍历数组的第一个元素开始遍历，当遍历的值等于前序遍历第一个元素（即根节点）时，
	// 说明当前索引前面的部分是构成左子树的范围，右边是构成右子树的范围
	for i := inStart; i <= inEnd; i++ {
		if inOrder[i] == rootVal {
			index = i
			break
		}
	}
	leftSize := index - inStart
	// 构造出左子树
	root.Left = build(preOrder, inOrder, preStart+1, preStart+leftSize, inStart, index-1)
	// 构造出右子树
	root.Right = build(preOrder, inOrder, preStart+leftSize+1, preEnd, index+1, inEnd)
	return root
}

// @lc code=end
