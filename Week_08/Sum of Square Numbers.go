package main

import "math"

//633. 平方数之和
//给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a2 + b2 = c。
//
//示例1:
//
//输入: 5
//输出: True
//解释: 1 * 1 + 2 * 2 = 5
func judgeSquareSum(c int) bool {
	i,j := 0,int(math.Sqrt(float64(c)))
	for i <= j {
		sum := i*i+j*j
		if sum == c {
			return true
		} else if sum < c {
			i++
		} else {
			j--
		}
	}
	return false
}
func main() {
	
}
