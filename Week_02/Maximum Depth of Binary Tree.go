package main

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root==nil{
		return 0
	}else {
		ldepth:=maxDepth(root.Left)+1
		rdepth:=maxDepth(root.Right)+1
		if ldepth>rdepth{
			return ldepth
		}else {
			return rdepth
		}
	}
}
