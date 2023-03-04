/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
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

// 快慢指针，快指针每次走两步，慢指针每次走一步，如果链表有环，快慢指针一定相遇。相遇时快指针走
// 的步数是慢指针走的步数+n*环的节点数。假设快慢指针在距离入环点k个节点处相遇，此时将慢指针移至
// 起点，然后两个指针每次往前走一步，最终快慢指针一定在环起点相遇
func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	// 无环返回 nil
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

// @lc code=end
