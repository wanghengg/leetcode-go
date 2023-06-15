/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
package leetcodego

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			front := q[0]
			q = q[1:]
			if i < n-1 {
				front.Next = q[0]
			} else {
				front.Next = nil
			}
			if front.Left != nil {
				q = append(q, front.Left)
			}
			if front.Right != nil {
				q = append(q, front.Right)
			}
		}
	}
	return root
}

func connect1(root *Node) *Node {
	if root == nil {
		return nil
	}
	for leftmost := root; leftmost.Left != nil; leftmost = leftmost.Left {
		for head := leftmost; head != nil; head = head.Next {
			head.Left.Next = head.Right
			if head.Next != nil {
				head.Right.Next = head.Next.Left
			}
		}
	}
	return root
}

// @lc code=end
