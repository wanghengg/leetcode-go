/*
 * @lc app=leetcode.cn id=1171 lang=golang
 *
 * [1171] 从链表中删去总和值为零的连续节点
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

func removeZeroSumSublists(head *ListNode) *ListNode {
	dummyNode := &ListNode{}
	dummyNode.Next = head
	var remove func(dummyNode *ListNode)
	remove = func(dummyNode *ListNode) {
		preSum := make([]int, 1)
		node := dummyNode
		for node.Next != nil {
			preSum = append(preSum, preSum[len(preSum)-1]+node.Next.Val)
			node = node.Next
		}
		m := make(map[int]int)
		left, right := 0, 0
		for ind, sum := range preSum {
			v, ok := m[sum]
			if !ok {
				m[sum] = ind
			} else {
				left = v
				right = ind
			}
		}
		if left == right {
			return
		}
		leftNode, rightNode := dummyNode, dummyNode
		for left > 0 {
			leftNode = leftNode.Next
			left--
		}
		for right > 0 {
			rightNode = rightNode.Next
			right--
		}
		leftNode.Next = rightNode.Next
		remove(dummyNode)
	}
	remove(dummyNode)
	return dummyNode.Next
}

// @lc code=end
