package main

import "math"

//72. 编辑距离
//给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。
//
//你可以对一个单词进行如下三种操作：
//
//插入一个字符
//删除一个字符
//替换一个字符
//
//
//示例 1：
//
//输入：word1 = "horse", word2 = "ros"
//输出：3
//解释：
//horse -> rorse (将 'h' 替换为 'r')
//rorse -> rose (删除 'r')
//rose -> ros (删除 'e')
func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)
	if len1 == 0 || len2 == 0 {
		return len1 + len2
	}

	//当字符串word1的长度为len1，字符串word2的长度为len2时，将word1转化为word2所使用的最少操作次数为dp[len1][len2]
	dp := make([][]int, len1+1) //初始化第一维数组，从1开始
	for i := range dp {
		dp[i] = make([]int, len2+1) //初始化第二维数组，从1开始
	}
	dp[0][1], dp[1][0], dp[1][1] = 1, 1, 1 //预初始化基本值，因为下面循环在len为1时不会进入
	for i := 1; i < len1+1; i++ {
		dp[i][0] = dp[i-1][0] + 1 //每次字符串长度加一，就相当于添加一个字符，因此需要加一步
	}
	for j := 1; j < len2+1; j++ {
		dp[0][j] = dp[0][j-1] + 1
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] //字符相等，不需要进行操作
			} else {
				dp[i][j] = int(math.Min(float64(dp[i-1][j-1]), math.Min(float64(dp[i-1][j]), float64(dp[i][j-1])))) + 1 //多出一步运算
			}
		}
	}
	return dp[len1][len2]
}
func main() {
	
}
