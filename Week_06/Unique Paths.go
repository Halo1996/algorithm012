package main
//62. 不同路径
//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
//
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
//
//问总共有多少条不同的路径？
//示例 1:
//
//输入: m = 3, n = 2
//输出: 3
//解释:
//从左上角开始，总共有 3 条路径可以到达右下角。
//1. 向右 -> 向右 -> 向下
//2. 向右 -> 向下 -> 向右
//3. 向下 -> 向右 -> 向右
//示例 2:
//
//输入: m = 7, n = 3
//输出: 28
/*
DP:
a 重复性（分治）
    每次只能向下或者向右
    a[i-1][j] + a[i][j-1]
b 定义状态数组
c DP方式
dp[i][j] = dp[i-1][j] + a[i][j-1]
动态规划和递归或者分治没有本质的区别，关键自安于有无
最优的子结构
共性： 找到重复子问题
差异性：最优子结构，中途可以淘汰次优解
*/

func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0{
		return 0
	}
	dp := make([][]int,m)
	for i := 0;i < m;i++{
		dp[i] = make([]int,n)
		dp[i][0] = 1
	}

	for i := 0;i < n;i++{
		dp[0][i] = 1
	}
	for i := 1;i < m;i++{
		for j := 1;j < n;j++{
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
func main() {
	
}
