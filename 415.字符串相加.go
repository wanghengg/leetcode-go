/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
package leetcodego

func addStrings(num1 string, num2 string) string {
	p1, p2 := len(num1)-1, len(num2)-1
	carry := 0
	res := []byte{}
	for p1 >= 0 && p2 >= 0 {
		c1, c2 := num1[p1], num2[p2]
		v1, v2 := int(c1-'0'), int(c2-'0')
		sum := v1 + v2 + carry
		carry = sum / 10
		sum %= 10
		res = append([]byte{byte(sum) + '0'}, res...)
		p1--
		p2--
	}
	for p1 >= 0 {
		c1 := num1[p1]
		v1 := int(c1 - '0')
		sum := v1 + carry
		carry = sum / 10
		sum %= 10
		res = append([]byte{byte(sum) + '0'}, res...)
		p1--
	}
	for p2 >= 0 {
		c2 := num2[p2]
		v2 := int(c2 - '0')
		sum := v2 + carry
		carry = sum / 10
		sum %= 10
		res = append([]byte{byte(sum) + '0'}, res...)
		p2--
	}
	if carry > 0 {
		res = append([]byte{byte(carry) + '0'}, res...)
	}
	return string(res)
}

// @lc code=end
