package main

import "fmt"

//152. 乘积最大子数组
//给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
//
//
//
//示例 1:
//
//输入: [2,3,-2,4]
//输出: 6
//解释: 子数组 [2,3] 有最大乘积 6。
//示例 2:
//
//输入: [-2,0,-1]
//输出: 0
//解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
func maxProduct(nums []int) int {
	maxF:=make([]int,len(nums))
	minF:=make([]int,len(nums))
	maxF[0],minF[0]=nums[0],nums[0]
	result:=nums[0]
	for i:=1;i<len(nums);i++{
		maxF[i]=max(maxF[i-1]*nums[i],max(nums[i],minF[i-1]*nums[i]))
		minF[i]=min(maxF[i-1]*nums[i],min(nums[i],minF[i-1]*nums[i]))
		result=max(result,maxF[i])
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	nums:=[]int{2,3,-2,4}
	result:=maxProduct(nums)
	fmt.Println(result)
}
