package main

import "fmt"

//20. 有效的括号
//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//注意空字符串可被认为是有效字符串。
//解题方法：像洋葱一样，应该使用栈
//解题思路：
// 1.条件 字符串数组长度模除2等于1则不是有效
// 2.遍历当前括号与栈顶元素匹配，匹配则出栈，不匹配则入栈
// 3.最后栈为空则有效不为空则无效

//func isValid(s string) bool {
//	n:=len(s)
//	if n%2==1{
//		return false
//	}
//	hash:=map[byte]byte{
//		')':'(',
//		']':'[',
//		'}':'{',
//	}
//	fmt.Println(hash['('])
//	stack:=[]byte{}
//	for i:=0;i<n;i++{
//		if hash[s[i]]>0{
//			if len(stack)==0||stack[len(stack)-1]!=hash[s[i]]{
//				return false
//			}
//			stack=stack[:len(stack)-1]
//		}else {
//			stack=append(stack,s[i])
//		}
//	}
//	return len(stack)==0
//}
func isValid(s string)bool{
	n:=len(s)
	if n%2==1{
		return false
	}
	stack:=[]byte{}
	for i:=0;i<len(s);i++{
		if s[i] == '(' {
			stack = append(stack, ')')
		}else if s[i] == '[' {
			stack = append(stack, ']')
		}else if s[i] == '{' {
			stack = append(stack, '}')
		}else if len(stack)==0||s[i] != stack[len(stack)-1]{
			return false
		}else if s[i] == stack[len(stack)-1]{
			stack=stack[:len(stack)-1]
		}
	}
	return len(stack)==0
}
func main() {
	s:="}{"
	result:=isValid(s)
	fmt.Println(result)
}