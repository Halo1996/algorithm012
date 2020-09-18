package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return dfs(candidates, target)
}

func dfs(nums []int, target int) [][]int {
	ret := [][]int{}
	for i, n := range nums {
		if i > 0 && nums[i-1] == n {
			continue
		} else if target < n {
			break
		} else if target == n {
			ret = append(ret, []int{n})
			continue
		}
		for _, v := range dfs(nums[i+1:], target-n) {
			fmt.Println(v)
			ret = append(ret, append([]int{n}, v...))
		}
	}
	return ret
}
func main() {
	nums:=[]int{10,1,2,7,6,1,5}
	result:=combinationSum2(nums,8)
	fmt.Println("result",result)
}