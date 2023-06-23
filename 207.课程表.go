/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

// @lc code=start
package leetcodego

func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	visited := make([]int, numCourses)
	var valid = true
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}
	var dfs func(u int)
	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 1 {
				valid = false
				return
			} else if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			}
		}
		visited[u] = 2
	}
	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid
}

// @lc code=end
