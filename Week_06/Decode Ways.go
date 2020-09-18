package main

import "fmt"

//91. 解码方法
//一条包含字母 A-Z 的消息通过以下方式进行了编码：
//
//'A' -> 1
//'B' -> 2
//...
//'Z' -> 26
//给定一个只包含数字的非空字符串，请计算解码方法的总数。
//
//示例 1:
//
//输入: "12"
//输出: 2
//解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	pre, cur, tmp := 1, 1, 0
	for i := 1; i < len(s); i++ {
		switch {
		case s[i] == '0' && s[i-1] != '1' && s[i-1] != '2':
			return 0
		case s[i] == '0':
			cur = pre
		case (s[i] <= '6' && (s[i-1] == '1' || s[i-1] == '2')) || (s[i] > '6' && s[i-1] == '1'):
			tmp = cur
			cur = cur + pre
			pre = tmp
		default:
			pre = cur
		}
	}
	return cur
}
func main() {
	str:="1234512"
	result:=numDecodings(str)
	fmt.Println(result)
}
