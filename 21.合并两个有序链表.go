/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
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

// 遍历
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	preHead := &ListNode{}
	head := preHead
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next
	}
	if list1 != nil {
		head.Next = list1
	}
	if list2 != nil {
		head.Next = list2
	}
	return preHead.Next
}

// 递归
func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists1(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists1(list1, list2.Next)
		return list2
	}
}

// @lc code=end
