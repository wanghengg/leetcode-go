/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
package leetcodego

func isValid(s string) bool {
	stack := make([]rune, 0, len(s))
	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		} else {
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				if (top == '(' && v == ')') ||
					(top == '{' && v == '}') ||
					(top == '[' && v == ']') {
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			} else {
				return false
			}
		}
	}
	if len(stack) > 0 {
		return false
	} else {
		return true
	}
}

// @lc code=end
