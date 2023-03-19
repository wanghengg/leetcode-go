/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	l1 := head
	l2 := head
	for l2 != nil {
		if l1.Val == l2.Val {
			l2 = l2.Next
		} else {
			l1.Next = l2
			l1 = l1.Next
		}
	}
	l1.Next = l2
	return head
}

// @lc code=end
