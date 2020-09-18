package main
//980. 不同路径 III
//在二维网格 grid 上，有 4 种类型的方格：
//
//1 表示起始方格。且只有一个起始方格。
//2 表示结束方格，且只有一个结束方格。
//0 表示我们可以走过的空方格。
//-1 表示我们无法跨越的障碍。
//返回在四个方向（上、下、左、右）上行走时，从起始方格到结束方格的不同路径的数目。
//
//每一个无障碍方格都要通过一次，但是一条路径中不能重复通过同一个方格。
//
//
//
//示例 1：
//
//输入：[[1,0,0,0],[0,0,0,0],[0,0,2,-1]]
//输出：2
//解释：我们有以下两条路径：
//1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2)
//2. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2)
func uniquePathsIII(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// 待走的空格总数，用于终止条件判断
	var steps int
	// 起点坐标
	var row, col int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				row, col = i, j
			} else if grid[i][j] == 0 {
				steps++
			}
		}
	}
	// 回溯(当前坐标，已经走过的步数)
	var bfs func(row, col int, curSteps int)
	// 合法路径数
	var ans int
	var dirs = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	bfs = func(row, col int, curSteps int) {
		// 越界判断
		if row < 0 || row >= m || col < 0 || col >= n {
			return
		}
		// 重复走过的坐标
		if grid[row][col] == -1 || grid[row][col] == 1 {
			return
		}
		// 走到终点，并且已经走过的步数=待走的步数，路径数+1
		if grid[row][col] == 2 {
			if curSteps == steps {
				ans++
			}
			return
		}
		// 添加回溯标记
		grid[row][col] = -1
		for i := range dirs {
			// 上下左右深搜，已走步数+1
			bfs(row+dirs[i][0], col+dirs[i][1], curSteps+1)
		}
		// 去除回溯标记
		grid[row][col] = 0
	}
	for i := range dirs {
		// 从起点的上下左右四个方向开始搜索
		bfs(row+dirs[i][0], col+dirs[i][1], 0)
	}
	return ans
}
func main() {
	
}
