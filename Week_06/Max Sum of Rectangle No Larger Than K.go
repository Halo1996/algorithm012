package main

import "math"

//363. 矩形区域不超过 K 的最大数值和
//给定一个非空二维矩阵 matrix 和一个整数 k，找到这个矩阵内部不大于 k 的最大矩形和。
//
//示例:
//
//输入: matrix = [[1,0,1],[0,-2,3]], k = 2
//输出: 2
//解释: 矩形区域 [[0, 1], [-2, 3]] 的数值和是 2，且 2 是不超过 k 的最大数字（k = 2）。
func maxSumSubmatrix(matrix [][]int, k int) int {
	a := matrix

	m := len(a)
	n := len(a[0])

	res := math.MinInt64
	for i := 0; i < m; i++ {
		list := make([]int, n)
		for p := i; p < m; p++ {
			for j := 0; j < n; j++ {
				list[j] += a[p][j]
			}

			temp := maxSumArrayWithK(list, k)
			if temp == k {
				return k
			}
			if temp < k && temp > res {
				res = temp
			}
		}
	}
	return res
}

func maxSumArrayWithK(list []int, k int) int {
	res := math.MinInt64
	sum := 0
	for i := 0; i < len(list); i++ {
		sum += list[i]
		res = max(res, sum)
		if res == k {
			return k
		}
		if sum < 0 {
			sum = 0
		}
	}
	if res < k {
		return res
	}

	res = math.MinInt64
	for i := 0; i < len(list); i++ {
		sum = 0
		for j := i; j < len(list); j++ {
			sum += list[j]
			if sum == k {
				return k
			}
			if sum > res && sum < k {
				res = sum
			}
		}
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	
}
