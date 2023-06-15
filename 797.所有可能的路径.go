/*
 * @lc app=leetcode.cn id=797 lang=golang
 *
 * [797] 所有可能的路径
 */

// @lc code=start
package leetcodego

func allPathsSourceTarget(graph [][]int) [][]int {
	res := [][]int{}
	path := []int{}
	var dfs func(index int)
	dfs = func(index int) {
		path = append(path, index)
		if index == len(graph)-1 {
			res = append(res, append([]int{}, path...))
			return
		}
		for _, v := range graph[index] {
			dfs(v)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}

// @lc code=end
