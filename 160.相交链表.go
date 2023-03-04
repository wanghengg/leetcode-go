/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1 := headA
	l2 := headB
	for l1 != nil && l2 != nil {
		if l1 == l2 {
			return l1
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 == nil && l2 == nil {
		return nil
	}
	steps := 0
	isL1Longer := false
	if l1 == nil {
		for l2 != nil {
			steps++
			l2 = l2.Next
		}
	} else {
		isL1Longer = true
		for l1 != nil {
			steps++
			l1 = l1.Next
		}
	}
	l1 = headA
	l2 = headB
	if isL1Longer {
		for i := 0; i < steps; i++ {
			l1 = l1.Next
		}
	} else {
		for i := 0; i < steps; i++ {
			l2 = l2.Next
		}
	}
	for l1 != nil && l2 != nil {
		if l1 == l2 {
			return l1
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return nil
}

// @lc code=end
