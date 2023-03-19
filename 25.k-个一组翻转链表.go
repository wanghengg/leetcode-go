/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
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

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyNode := &ListNode{}
	dummyNode.Next = head
	nextNode := getKthNode(dummyNode, k)
	if nextNode == nil { // 剩余链表不足k个节点
		return dummyNode.Next
	}
	if nextNode.Next == nil { // 链表刚好k个节点
		return reverseList(dummyNode.Next)
	}
	newHead := nextNode.Next
	nextNode.Next = nil
	firstStage := reverseList(dummyNode.Next)
	secondStage := reverseKGroup(newHead, k)
	head.Next = secondStage
	return firstStage
}

func getKthNode(head *ListNode, k int) *ListNode {
	for i := k; i > 0 && head != nil; i-- {
		head = head.Next
	}
	return head
}

// func reverseList(head *ListNode) *ListNode {
// 	var node1 *ListNode = nil
// 	node2 := head
// 	for node2 != nil {
// 		tmp := node2.Next
// 		node2.Next = node1
// 		node1 = node2
// 		node2 = tmp
// 	}
// 	return node1
// }

// @lc code=end
