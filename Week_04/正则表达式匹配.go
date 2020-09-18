package main
//剑指 Offer 19. 正则表达式匹配
//请实现一个函数用来匹配包含'. '和'*'的正则表达式。模式中的字符'.'表示任意一个字符，而'*'表示它前面的字符可以出现任意次（含0次）。在本题中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但与"aa.a"和"ab*a"均不匹配。
//
//示例 1:
//
//输入:
//s = "aa"
//p = "a"
//输出: false
//解释: "a" 无法匹配 "aa" 整个字符串。
func isMatch(s string, p string) bool {
	if s == p {
		return true
	}

	if len(p) > 1 && p[1] == '*' {
		switch p[0] {
		case '.':
			for i := 0; i <= len(s); i++ {
				if isMatch(s[i:], p[2:]) {
					return true
				}
			}
		default:
			if isMatch(s, p[2:]) {
				return true
			}

			for i := 0; i < len(s) && s[i] == p[0]; i++ {
				if isMatch(s[i+1:], p[2:]) {
					return true
				}
			}
		}
	} else {
		if s == "" || p == "" {
			return false
		}
		if s[0] == p[0] || p[0] == '.' {
			return isMatch(s[1:], p[1:])
		}
	}

	return false
}
func main() {
	
}
