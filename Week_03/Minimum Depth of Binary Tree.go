package main
//111. 二叉树的最小深度
//给定一个二叉树，找出其最小深度。
//
//最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//
//说明: 叶子节点是指没有子节点的节点。
//
//示例:
//
//给定二叉树 [3,9,20,null,null,15,7],
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//返回它的最小深度  2.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func minDepth(root *TreeNode) int {
	if root==nil{
		return 0
	}
	dl:=minDepth(root.Left)
	dr:=minDepth(root.Right)
	if root.Left==nil{
		return dr+1
	}else if root.Right==nil{
		return dl+1
	}else {
		return min(dl,dr)+1
	}
}

func min(a, b int) int{
	if a < b{
		return a
	}
	return b
}
func main() {
	
}
