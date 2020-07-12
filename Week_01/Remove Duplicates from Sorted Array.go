package main

import "fmt"

//给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
//不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
//示例 1:
//给定数组 nums = [1,1,2],
//函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。
//你不需要考虑数组中超出新长度后面的元素。
//解题思路：双指针法
func main() {
	//arr:=[]int{1,1,2}
	arr:=[]int{0,0,0,0,0,1,1,1,1,3,3,3,3,4,4,4,12}
	result:=removeDuplicates(arr)
	fmt.Println(result)
}
func removeDuplicates(nums []int) int{
	if nums==nil{
		return 0
	}
	i:=0
	for j:=1;j<len(nums);j++{
		if nums[i]!=nums[j]{
			i++
			nums[i]=nums[j]
		}
	}
	return i+1
}