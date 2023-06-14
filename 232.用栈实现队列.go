/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

// @lc code=start
package leetcodego

type MyQueue struct {
	s1, s2 []int // s1存放push进去的数据，s2存放pop的数据
}

func ConstructorI() MyQueue {
	return MyQueue{
		s1: []int{},
		s2: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.s1 = append(this.s1, x)
}

func (this *MyQueue) Pop() int {
	if len(this.s2) == 0 {
		// 将s1中的所有数据搬到s2
		for len(this.s1) > 0 {
			top := this.s1[len(this.s1)-1]
			this.s1 = this.s1[:len(this.s1)-1]
			this.s2 = append(this.s2, top)
		}
	}
	if len(this.s2) == 0 {
		return -1
	}
	front := this.s2[len(this.s2)-1]
	this.s2 = this.s2[:len(this.s2)-1]
	return front
}

func (this *MyQueue) Peek() int {
	if len(this.s2) == 0 {
		// 将s1中的所有数据搬到s2
		for len(this.s1) > 0 {
			top := this.s1[len(this.s1)-1]
			this.s1 = this.s1[:len(this.s1)-1]
			this.s2 = append(this.s2, top)
		}
	}
	if len(this.s2) == 0 {
		return -1
	}
	return this.s2[len(this.s2)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.s1)+len(this.s2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end
