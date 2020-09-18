package main
//24. 两两交换链表中的节点
//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//示例:
//
//给定 1->2->3->4, 你应该返回 2->1->4->3.
type ListNode struct {
	     Val int
	     Next *ListNode
}
//func swapPairs(head *ListNode) *ListNode {
//	head = &ListNode{
//		Next: head,
//	}
//	prev:=head
//	for prev.Next!=nil&&prev.Next.Next!=nil{
//		a1,a2,a3,a4:=prev,prev.Next,prev.Next.Next,prev.Next.Next.Next
//		a1.Next,a2.Next,a3.Next=a3,a4,a2
//		prev=prev.Next
//	}
//	return head.Next
//
//}
func swapPairs(head *ListNode) *ListNode {
	head = &ListNode{
		Next: head,
	}
	prev:=head
	for prev.Next!=nil&&prev.Next.Next!=nil{
		a1,a2,a3,a4 := prev,prev.Next,prev.Next.Next,prev.Next.Next.Next
		a1.Next,a2.Next,a3.Next=a3,a4,a2
		prev=prev.Next.Next
	}
	return head.Next
}