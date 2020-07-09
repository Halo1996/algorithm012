package main

type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}
func hasPathSum(root *TreeNode, sum int) bool {
	//特殊情况 如果根节点为空返回false
	if root==nil{
		return false
	}
	//声明两个数组
	queNode:=[]*TreeNode{}
	queVal:=[]int{}
	//将根节点放入数组
	queNode=append(queNode,root)
	queVal=append(queVal,root.Val)
	for len(queNode)!=0{
		now:=queNode[0]
		queNode=queNode[1:]
		temp:=queVal[0]
		queVal=queVal[1:]
		if now.Left==nil&&now.Right==nil{
			if sum==temp{
				return true
			}
			continue
		}
		if now.Left!=nil{
			queNode=append(queNode,now.Left)
			queVal=append(queVal,now.Val+temp)
		}
		if now.Right!=nil{
			queNode=append(queNode,now.Right)
			queVal=append(queVal,now.Val+temp)
		}
	}
	return false
}