package main

import "fmt"
//1. 两数之和
//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

//示例:
//
//给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
//解题思路1：使用哈希表 定义一个map 遍历nums find值为目标值减去当前值
// 如果find值在map的key里则return 否则map中存当前值与当前位置的映射关系
func twoSum1(nums []int, target int) []int{
	m:=make(map[int]int)
	for i,v:=range nums{
		find:=target-v
		if _,ok:=m[find];ok{
			return []int{m[find],i}
		}
		m[v]=i
	}
	return []int{}
}
func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result:=twoSum1(nums,target)
	fmt.Println(result)
}