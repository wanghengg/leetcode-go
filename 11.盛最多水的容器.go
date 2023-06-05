/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
package leetcodego

/*
* 设两指针i,j，指向的水槽板高度分别为h[i],h[j],此状态下水槽面积为s(i,j)，由于水槽高度
* 由两个板中的短板决定，所以s(i,j) = min(h[i],h[j])*(j-i)
* 在每个状态下，收缩挡板一格，都会使得水槽的宽度-1，而水槽的高度则取决于移动的哪个挡板
* 1. 如果收缩移动短板一格，水槽的高度可能减小也可能变大，所以水槽面积同样可能变大
* 2. 如果收缩移动长板一格，水槽的高度只可能不变或者更小，所以水槽面积只可能变小或者不变
* 根据上面的两种情况，只需要考虑短板收缩的情况

* 最初将指针i,j分别指向第一个和最后一个挡板，每次收缩短板直到 i 和 j 相遇
 */
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	res := 0
	for i < j {
		res = max(res, min(height[i], height[j])*(j-i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}

// @lc code=end
