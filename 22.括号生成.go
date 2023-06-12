/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
package leetcodego

func generateParenthesis(n int) []string {
	res := []string{}
	var dfs func(n1, n2 int, combine string)
	dfs = func(n1, n2 int, combine string) {
		if n2 > n || n1 > n || n2 > n1 {
			return
		}
		if n1 == n && n2 == n {
			res = append(res, combine)
			return
		}
		dfs(n1+1, n2, combine+"(")
		dfs(n1, n2+1, combine+")")
	}
	dfs(0, 0, "")
	return res
}

// @lc code=end
