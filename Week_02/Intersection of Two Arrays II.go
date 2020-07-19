package main

import "fmt"

//给定两个数组，编写一个函数来计算它们的交集。
//
//示例 1：
//
//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2,2]
//解法：hash
func main() {
	nums1:=[]int{1,2,2,1}
	nums2:=[]int{2,2}
	fmt.Println(intersect(nums1,nums2))
}
func intersect(nums1 []int, nums2 []int) []int {
	m:=map[int]int{}
	for _,v :=range nums1{
		m[v]+=1
	}
	k:=0
	for _,v:=range nums2{
		if m[v]>0{
			m[v]-=1
			nums2[k]=v
			k++
		}
	}
	return nums2[0:k]
}
