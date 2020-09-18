package main
//144. 二叉树的前序遍历
//给定一个二叉树，返回它的 前序 遍历。
//
// 示例:
//
//输入: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
//输出: [1,2,3]
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	if root==nil{
		return []int{}
	}
	res:=[]int{}
	res=append(res,root.Val)
	res=append(res,preorderTraversal(root.Left)...)
	res=append(res,preorderTraversal(root.Right)...)
	return res
}
