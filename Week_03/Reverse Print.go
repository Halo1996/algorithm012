package main

type ListNode struct {
	     Val int
	     Next *ListNode
}
func reversePrint(head *ListNode) []int {
	var s []int
	var rs []int
	for head!=nil{
		s=append(s,head.Val)
		head=head.Next
	}
	for i:=len(s);i>0;i--{
		rs=append(rs,s[i])
	}
	return rs
}
