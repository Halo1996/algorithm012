package main
//589. N叉树的前序遍历
//给定一个 N 叉树，返回其节点值的前序遍历。
//
//例如，给定一个 3叉树 :
//
//返回其前序遍历: [1,3,5,6,2,4]。
//
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
//解法：递归 深度优先遍历
//解题思路：将根节点值存入数组 遍历子节点

func preorder(root *Node) []int {
	if root==nil{
		return []int{}
	}
	res:=[]int{}
	res=append(res,root.Val)
	for _,v:=range root.Children{
		res = append(res, preorder(v)...)
	}
	return res
}

