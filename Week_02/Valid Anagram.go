package main

//有效的字母异位词
//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
//示例 1:
//
//输入: s = "anagram", t = "nagaram"
//输出: true
//方法一：哈希表
//思路：
//方法二：排序
//思路：
func main() {
	s := "anagram"
	t := "nagaram"
	isAnagram(s,t)
}
func isAnagram(s string, t string) bool {
	if len(s)!=len(t){
		return false
	}
	setChar:=make([]int,26,26)
	for i:=0;i<len(s);i++{
		setChar[s[i]-'a']++
	}
	for i:=0;i<len(t);i++{
		setChar[t[i]-'a']--
		if setChar[t[i]-'a']<0{
			return false
		}
	}
	return true
}
