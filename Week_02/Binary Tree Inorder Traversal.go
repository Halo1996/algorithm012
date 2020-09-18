package main
//94. 二叉树的中序遍历
//给定一个二叉树，返回它的中序 遍历。
//
//示例:
//
//输入: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
//输出: [1,3,2]

func inorderTraversal(root *TreeNode) []int {
	if root==nil{
		return []int{}
	}
	r:=append(inorderTraversal(root.Left),root.Val)
	r=append(r,inorderTraversal(root.Right)...)
	return r
}
