package main
//279. 完全平方数
//给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
//
//示例 1:
//
//输入: n = 12
//输出: 3
//解释: 12 = 4 + 4 + 4.
//示例 2:
//
//输入: n = 13
//输出: 2
//解释: 13 = 4 + 9.
func numSquares(n int) int {
	//把平方数写入切片
	sum :=[]int{}
	for i:=0;i<=n/2+2;i++{
		sum=append(sum,i*i)
	}
	//dp存放0-n的最小平方数个数
	dp:=make([]int,0,1000)
	dp=append(dp,0)
	//第一个循环求0-n的最小平方数个数
	for i:=1;i<=n;i++{
		mid,p,j := n,n,1
		//第二个循环求i的最小平方数个数
		for ;i>=sum[j];j++{
			//dp[i]=dp[i-sum[j]]+1
			//p=min(dp[i]+1,mid)
			//缩写为：p=min(dp[i-sum[j]]+1,mid)
			p=min(dp[i-sum[j]]+1,mid)
			mid=p
		}
		dp=append(dp,mid)
	}
	return dp[n]
}
func min(a,b int)int{
	if(a>=b){
		return b
	}
	return a
}
func main() {
	
}
