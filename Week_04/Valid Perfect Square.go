package main
//367. 有效的完全平方数
//给定一个正整数 num，编写一个函数，如果 num 是一个完全平方数，则返回 True，否则返回 False。
//
//说明：不要使用任何内置的库函数，如  sqrt。
//
//示例 1：
//
//输入：16
//输出：True
//示例 2：
//
//输入：14
//输出：False
func isPerfectSquare(num int) bool {
	if num < 2{
		return true
	}
	min , max := 0 , num/2 + 1
	for max - min >= 0{
		mid := (min + max) / 2
		if num > mid * mid{
			min = mid + 1
		}else if num < mid * mid{
			max = mid - 1
		}else if num == mid * mid{
			return true
		}
	}
	return false
}
func main() {
	
}
