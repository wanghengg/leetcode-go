/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
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

func deleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummyNode := &ListNode{
		Next: head,
	}
	node := dummyNode
	first := dummyNode
	for node.Next != nil {
		if node.Next.Val == first.Next.Val {
			node = node.Next
		} else {
			if node == first.Next {
				first = first.Next
			} else {
				first.Next = node.Next
			}
		}
	}
	if node != first.Next && node.Val == first.Next.Val {
		first.Next = node.Next
	}
	return dummyNode.Next
}

// @lc code=end
