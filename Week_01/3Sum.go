package main

import (
	"fmt"
	"sort"
)

//15. 三数之和
//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
//
//示例：
//
//给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//                 i   l            r
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
//解题方法：双指针
//解题思路: 先排序防止重复 当前指针，左指针和右指针

//func threeSum(nums []int) [][]int {
//	//获取数组长度
//	n := len(nums)
//	//定义二维数组
//	var res [][]int
//	//	排序 为了遍历起来更方便
//	sort.Ints(nums)
//	//遍历
//	for i := 0; i < n; i++ {
//		//目标值为0-当前值
//		target := 0 - nums[i]
//		//左指针为当前下一个元素
//		l := i + 1
//		//右指针为最后一个元素
//		r := n - 1
//		//如果当前元素大于0说明之后没有目标值了 直接返回
//		if nums[i] > 0 {
//			return res
//		}
//		//如果是第一个元素或者当前元素不等于上一个元素，即不重复
//		if i == 0 || nums[i] != nums[i-1] {
//			//遍历直到左右指针相遇
//			for l < r {
//				//如果左指针加右指针等于目标值
//				if nums[l]+nums[r]==target{
//					//则是目标结果 将左指针，右指针和当前元素存入二维数组中
//					res=append(res,[]int{nums[i],nums[l],nums[r]})
//					//当左指针与下一个相同则左指针向右走一步
//					for nums[l] == nums[l + 1]{
//						l++
//					}
//					//当右指针与前一个元素相同则右指针向左走一步
//					for nums[r] == nums[r - 1]{
//						r--
//					}
//					l++
//					r--
//				} else if (nums[l] + nums[r]  > target) {
//					r--
//				} else {
//					l++
//				}
//			}
//		}
//
//	}
//	return res
//}
func threeSum(nums []int) [][]int{
	n:=len(nums)
	res:=[][]int{}
	sort.Ints(nums)
	for i:=0;i<n;i++{
		target:=0-nums[i]
		l:=i+1
		r:=n-1
		if nums[i]>0{
			return res
		}
		if i==0||nums[i]!=nums[i-1]{
			for l<r  {
				if nums[l]+nums[r]==target{
					res=append(res,[]int{nums[i],nums[l],nums[r]})
					if nums[l]==nums[l+1]{
						l++
					}
					if nums[r]==nums[r-1]{
						r--
					}
					l++
					r--
				}else if nums[l]+nums[r]>target{
					r--
				}else {
					l++
				}
			}
		}
	}
	return res
}
func main() {
	nums:=[]int{-1, 0, 1, 2, -1, -4}
	result:=threeSum(nums)
	fmt.Println(result)
}