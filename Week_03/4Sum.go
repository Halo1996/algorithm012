package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	//排序
	sort.Ints(nums)
	//获取数组长度
	n := len(nums)
	var res [][]int
	//从0遍历到数组长度-3
	for i := 0; i < n-3; i++ {
		//如果当前元素与前一元素相同则跳出循环
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		//否则a指针=当前的下一个元素
		a := i+1
		//从i+1遍历到数组长度-2
		for j := i+1; j < n-2; j++ {
			//如果当前的元素与前一元素相同则跳出循环
			if j > a && nums[j] == nums[j-1] {
				continue
			}
			//否则左指针=当前元素的下一个 右指针=最后一个元素
			l, r := j + 1, n - 1
			//遍历左指针到右指针
			for l < r {
				//临时值=四个指针元素之和
				tmp := nums[i] + nums[j] + nums[l] + nums[r]
				//如果四个指针元素之和等于目标值
				if tmp == target {
					//就append到结果中
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					//左指针右移右指针左移
					l++
					r--
					//如果左指针相邻左侧元素相同则继续右移
					for l < r && nums[l] == nums[l-1] {
						l++
					}
					//如果右指针相邻右侧元素相同则继续左移
					for l < r && nums[r] == nums[r+1] {
						r--
					}
				//	如果四数之和小于目标值则左指针右移
				} else if tmp < target {
					l++
					//如果左指针相邻左侧元素相同则继续右移
					for l < r && nums[l] == nums[l-1] {
						l++
					}
					//	如果四数之和大于目标值则右指针左移
				} else {
					r--
					//如果右指针相邻右侧元素相同则继续左移
					for l < r && nums[r] == nums[r+1] {
						r--
					}
				}
			}

		}
	}
	return res
}