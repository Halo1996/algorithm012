package main
//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
//示例:
//
//输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]
//时间复杂度：O(n)
//空间复杂度：O(1)
func main() {
	nums:=[]int{3,5,0,8,6}
	moveZeroes(nums)
}
func moveZeroes(nums []int)  {
	if nums==nil{
		return
	}
	var zeroCount int
	for i,n:=range nums{
		if n==0{
			zeroCount++
		}else {
			nums[i-zeroCount]=nums[i]
		}
	}
	for i:=0;i<zeroCount;i++{
		nums[len(nums)-i-1]=0
	}
}
//func moveZeroes(nums []int)  {
//	if nums==nil{
//		return
//	}
//	var j int
//	for i:=0;i<len(nums);i++{
//		if nums[i]==0{
//			continue
//		}
//		nums[i],nums[j]=nums[j],nums[i]
//		j++
//	}
//}

