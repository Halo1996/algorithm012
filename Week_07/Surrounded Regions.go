package main

//130. 被围绕的区域
//给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。
//
//找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
//
//示例:
//
//X X X X
//X O O X
//X X O X
//X O X X
//运行你的函数后，矩阵变为：
//
//X X X X
//X X X X
//X X X X
//X O X X
var n, m int

func solve(board [][]byte)  {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	n, m = len(board), len(board[0])
	for i := 0; i < n; i++ {
		dfs(board, i, 0)
		dfs(board, i, m - 1)
	}
	for i := 1; i < m - 1; i++ {
		dfs(board, 0, i)
		dfs(board, n - 1, i)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, x, y int) {
	if x < 0 || x >= n || y < 0 || y >= m || board[x][y] != 'O' {
		return
	}
	board[x][y] = 'A'
	dfs(board, x + 1, y)
	dfs(board, x - 1, y)
	dfs(board, x, y + 1)
	dfs(board, x, y - 1)
}
func main() {

}
