package main
//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
//示例：
//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
//解题思路：暴力
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//	设置哨兵
	prehead:=&ListNode{}
	result:=prehead
	for l1!=nil&&l2!=nil{
		if l1.Val<l2.Val{
			prehead.Next=l1
			l1=l1.Next
		}else {
			prehead.Next=l2
			l2=l2.Next
		}
		prehead=prehead.Next
	}
	if l1!=nil{
		prehead.Next=l1
	}
	if l2!=nil{
		prehead.Next=l2
	}
	return result.Next
}
