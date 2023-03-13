/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
package leetcodego

func combine(n int, k int) [][]int {
	res := [][]int{}
	tmp := []int{}
	var dfs func(left int)
	dfs = func(left int) {
		/*
			剪枝
			tmp的长度加上区间[left, n]的长度小于k，不可能构造出长度为k的数组，提前返回
		*/
		if left+k-len(tmp) > n+1 {
			return
		}
		// tmp的长度刚好等于k，得到一个合理的结果
		if len(tmp) == k {
			com := make([]int, k)
			copy(com, tmp)
			res = append(res, com)
			return
		}
		// 保证后面append的数大于之前已有的值，不会出现逆序对，这样保证了不会重复，例如
		// [1,2]和[2,1]这样的情况
		for i := left; i <= n; i++ {
			tmp = append(tmp, i)
			dfs(i + 1)
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(1)
	return res
}

// @lc code=end
