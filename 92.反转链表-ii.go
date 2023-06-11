/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
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

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{
		Next: head,
	}
	first, second := dummyNode, dummyNode
	for left > 1 {
		first = first.Next
		left--
	}
	for right > 0 {
		second = second.Next
		right--
	}
	tmp := second.Next
	second.Next = nil
	second = tmp
	l1 := first.Next
	first.Next = nil
	head = l1
	l2 := l1.Next
	l1.Next = nil
	for l1 != nil && l2 != nil {
		tmp := l2.Next
		l2.Next = l1
		l1 = l2
		l2 = tmp
	}
	first.Next = l1
	head.Next = second
	return dummyNode.Next
}

// @lc code=end
