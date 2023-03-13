/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
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

func reverseList(head *ListNode) *ListNode {
	var node1 *ListNode = nil
	node2 := head
	for node2 != nil {
		tmp := node2.Next
		node2.Next = node1
		node1 = node2
		node2 = tmp
	}
	return node1
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := head.Next
	head.Next = nil
	newHead := reverseList1(tmp)
	tmp.Next = head
	return newHead
}

// @lc code=end
