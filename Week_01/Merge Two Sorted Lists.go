package main
//21. 合并两个有序链表
//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
//示例：
//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
//解题思路：定义一个哑节点，循环，比较l1,l2，小的拼到后面然后往后走一步，哑节点每次往后走一步
//其中一个没了，把另一个剩下的拼过来
//时间复杂度：O(n+m)
//空间复杂度：O(1)
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
