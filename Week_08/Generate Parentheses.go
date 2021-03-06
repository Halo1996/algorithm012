package main
//22. 括号生成
//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
//示例：
//
//输入：n = 3
//输出：[
//       "((()))",
//       "(()())",
//       "(())()",
//       "()(())",
//       "()()()"
//     ]
func generateParenthesis(n int) []string {
	Output := new([]string)
	_generate(0, 0, n, "", Output)
	return *Output
}

func _generate(left int, right int, max int, s string, Output *[]string){
	// 递归终止条件
	if left == right && left ==  max{
		*Output = append(*Output, s)
		return
	}
	// 递归主体
	if left < max{
		_generate(left+1, right, max, s + "(", Output)
	}
	if right < left{
		_generate(left, right+1,  max, s + ")", Output)
	}
}
func main() {
	
}
