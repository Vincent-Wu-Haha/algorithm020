package _29

var dirX = []int{0, 1, 0, -1, 1, 1, -1, -1}
var dirY = []int{1, 0, -1, 0, 1, -1, 1, -1}

func updateBoard(board [][]byte, click []int) [][]byte {
	n := len(board)
	m := len(board[0])
	i := click[0]
	j := click[1]

	update(board, i, j, n, m)
	return board
}

func update(board [][]byte, i, j int, n, m int) {
	clicked(board, i, j, n, m)
	if board[i][j] == 'B' {
		for k := 0; k < len(dirX); k++ {
			if inBound(i+dirX[k], j+dirY[k], n, m) && board[i+dirX[k]][j+dirY[k]] == 'E' {
				update(board, i+dirX[k], j+dirY[k], n, m)
			}
		}
	}

}

func clicked(board [][]byte, i, j int, n, m int) {
	if !inBound(i, j, n, m) {
		return
	}
	switch board[i][j] {
	case 'M':
		board[i][j] = 'X'
	case 'E':
		var count byte
		for k := 0; k < len(dirX); k++ {
			if isM(board, i+dirX[k], j+dirY[k], n, m) {
				count++
			}
		}
		if count == 0 {
			board[i][j] = 'B'
		} else {
			board[i][j] = '0' + count
		}
	}

}

func inBound(i int, j int, n int, m int) bool {
	return i >= 0 && j >= 0 && i < n && j < m
}

func isM(board [][]byte, i, j int, n, m int) bool {
	if inBound(i, j, n, m) {
		return board[i][j] == 'M'
	}
	return false
}
