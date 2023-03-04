/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	p1 := dummy
	p1.Next = head
	p2 := head
	for i := 0; i < n; i++ {
		p2 = p2.Next
	}
	for p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	p1.Next = p1.Next.Next
	return dummy.Next
}

// @lc code=end
