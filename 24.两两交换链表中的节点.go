/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package leetcodego

func swapPairs(head *ListNode) *ListNode {
	dummyNode := &ListNode{
		Next: head,
	}
	node := dummyNode
	for dummyNode.Next != nil && dummyNode.Next.Next != nil {
		first, second := dummyNode.Next, dummyNode.Next.Next
		first.Next = second.Next
		second.Next = first
		dummyNode.Next = second
		dummyNode = dummyNode.Next.Next
	}
	return node.Next
}

// @lc code=end
