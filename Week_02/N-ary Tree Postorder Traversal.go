package main
//590. N叉树的后序遍历
//给定一个 N 叉树，返回其节点值的后序遍历。
//
//例如，给定一个 3叉树 :
//
//
//返回其后序遍历: [5,6,3,2,4,1].
//
type Node struct {
  Val int
  Children []*Node
}

func postorder(root *Node) []int {
  if root == nil {
    return []int{}
  }

  results := []int{}
  for _, v := range root.Children {
    results = append(results, postorder(v)...)
  }

  results = append(results, root.Val)

  return results
}