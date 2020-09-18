package main
//46. 全排列
//给定一个 没有重复 数字的序列，返回其所有可能的全排列。
//
//示例:
//
//输入: [1,2,3]
//输出:
//[
//  [1,2,3],
//  [1,3,2],
//  [2,1,3],
//  [2,3,1],
//  [3,1,2],
//  [3,2,1]
//]

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	option := make([]int, len(nums))
	var dfs func(int)
	dfs = func(idx int) {
		if idx == len(nums) {
			ans = append(ans, append([]int{}, option...))
			return
		}
		for i := idx; i < len(nums); i++ {
			option[idx] = nums[i]
			nums[i], nums[idx] = nums[idx], nums[i]
			dfs(idx + 1)
			nums[i], nums[idx] = nums[idx], nums[i]
		}
	}
	dfs(0)
	return ans
}
