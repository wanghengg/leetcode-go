/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
package leetcodego

import (
	"math"
)

func intToRoman(num int) string {
	m := map[int]string{
		1:    "I",
		2:    "II",
		3:    "III",
		4:    "IV",
		5:    "V",
		6:    "VI",
		7:    "VII",
		8:    "VIII",
		9:    "IX",
		10:   "X",
		20:   "XX",
		30:   "XXX",
		40:   "XL",
		50:   "L",
		60:   "LX",
		70:   "LXX",
		80:   "LXXX",
		90:   "XC",
		100:  "C",
		200:  "CC",
		300:  "CCC",
		400:  "CD",
		500:  "D",
		600:  "DC",
		700:  "DCC",
		800:  "DCCC",
		900:  "CM",
		1000: "M",
		2000: "MM",
		3000: "MMM",
	}
	res := ""
	count := 0
	for num > 0 {
		count++
		bit := num % 10
		num = num / 10
		if bit == 0 {
			continue
		}
		res = m[bit*int(math.Pow10(count-1))] + res
	}
	return res
}

// @lc code=end
