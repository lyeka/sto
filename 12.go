package sto

func exist(board [][]byte, word string) bool {
	for i, _ := range board {
		for j, _ := range board[0] {
			if dfs(i, j, 0, board, word) {
				return true
			}
		}
	}
	return false

}

func dfs(i, j, k int, board [][]byte, word string) bool {
	if i < 0 || i > len(board) || j < 0 || j > len(board[0]) || board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}
	board[i][j] = ' '
	res := dfs(i+1, j, k+1, board, word) || dfs(i-1, j, k+1, board, word) || dfs(i, j-1, k+1, board, word) || dfs(i, j+1, k+1, board, word)
	board[i][j] = word[k]
	return res
}
