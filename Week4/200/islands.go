package _00

func numIslands(grid [][]byte) int {
	visited := make([][]bool, len(grid))
	for i, bytes := range grid {
		visited[i] = make([]bool, len(bytes))
	}
	res := 0
	for i, bools := range visited {
		for j := range bools {
			if !visited[i][j] {
				found := false
				backtrack(grid, visited, i, j, &found)
				if found {
					res++
				}
			}
		}
	}
	return res
}

func backtrack(grid [][]byte, visited [][]bool, startI, startJ int, found *bool) {
	if startI >= 0 && startI < len(grid) && startJ >= 0 && startJ < len(grid[startI]) {
		if visited[startI][startJ] {
			return
		}
		visited[startI][startJ] = true
		if grid[startI][startJ] == []byte("0")[0] {
			return
		} else {
			*found = true
		}

		if startJ < len(grid[startI]) {
			backtrack(grid, visited, startI, startJ+1, found)
		}
		if startJ > 0 {
			backtrack(grid, visited, startI, startJ-1, found)
		}
		if startI < len(grid) {
			backtrack(grid, visited, startI+1, startJ, found)
		}
		if startI > 0 {
			backtrack(grid, visited, startI-1, startJ, found)
		}
	}
	return
}
