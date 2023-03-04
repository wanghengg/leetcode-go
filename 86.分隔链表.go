/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
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

func partition(head *ListNode, x int) *ListNode {
	dummy1 := &ListNode{}
	dummy2 := &ListNode{}
	p1 := dummy1 // 小于x的链表
	p2 := dummy2 // 大于x的链表
	for head != nil {
		if head.Val < x {
			p1.Next = head
			p1 = p1.Next
		} else {
			p2.Next = head
			p2 = p2.Next
		}
		head = head.Next
	}
	p2.Next = nil
	p1.Next = dummy2.Next
	return dummy1.Next
}

// @lc code=end
