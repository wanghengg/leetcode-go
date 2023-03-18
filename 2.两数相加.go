/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	h1 := l1
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	carry := 0
	for l1.Next != nil && l2.Next != nil {
		sum := l1.Val + l2.Val + carry
		carry = sum / 10
		l1.Val = sum % 10
		l1 = l1.Next
		l2 = l2.Next
	}
	sum := l1.Val + l2.Val + carry
	carry = sum / 10
	l1.Val = sum % 10
	if l1.Next == nil && l2.Next == nil {
		if carry > 0 {
			l1.Next = &ListNode{
				Val: carry,
			}
		}
		return h1
	} else if l1.Next == nil {
		l1.Next = l2.Next
		for l1.Next != nil {
			l1 = l1.Next
			sum := l1.Val + carry
			l1.Val = sum % 10
			carry = sum / 10
		}
		if carry > 0 {
			l1.Next = &ListNode{
				Val: carry,
			}
		}
		return h1
	} else {
		for l1.Next != nil {
			l1 = l1.Next
			sum := l1.Val + carry
			l1.Val = sum % 10
			carry = sum / 10
		}
		if carry > 0 {
			l1.Next = &ListNode{
				Val: carry,
			}
		}
		return h1
	}
}

// @lc code=end
