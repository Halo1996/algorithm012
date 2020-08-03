package main
//200. 岛屿数量
//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
//岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。
//
//此外，你可以假设该网格的四条边均被水包围。
//
//示例 1:
//
//输入:
//[
//['1','1','1','1','0'],
//['1','1','0','1','0'],
//['1','1','0','0','0'],
//['0','0','0','0','0']
//]
//输出: 1
func numIslands(grid [][]byte) int {
	count := 0
	for line := range grid {
		for column := range grid[line] {
			if grid[line][column] == '0' || grid[line][column] == '2' {
				continue
			}

			count++
			dfs(grid, line, column)
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
		return
	}

	if grid[i][j] != '1' {
		return
	}

	grid[i][j] = '2'
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}
func main() {
	
}
