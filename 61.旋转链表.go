/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
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

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	dummyNode := &ListNode{
		Next: head,
	}
	length := 0
	node := dummyNode
	for node.Next != nil {
		length++
		node = node.Next
	}
	k %= length
	if k == 0 {
		return dummyNode.Next
	}
	node.Next = head
	k = length - k
	node = dummyNode
	for k > 0 {
		node = node.Next
		k--
	}
	head = node.Next
	node.Next = nil
	return head
}

// @lc code=end
