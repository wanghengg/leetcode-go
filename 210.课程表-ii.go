/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 */

// @lc code=start
package leetcodego

func findOrder(numCourses int, prerequisites [][]int) []int {
	edges := make([][]int, numCourses)
	visited := make([]int, numCourses) // 标记节点遍历状态，0-表示未搜索，1-表示搜索中，2-表示已搜索
	res := []int{}
	var valid = true
	var dfs func(u int)

	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		res = append(res, u)
	}

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}
	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if !valid {
		return []int{}
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[numCourses-i-1] = res[numCourses-i-1], res[i]
	}
	return res
}

// @lc code=end
