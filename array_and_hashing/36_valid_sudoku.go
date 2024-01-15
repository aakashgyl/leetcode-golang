package main

import "fmt"

func main() {
	var board [][]byte
	board = [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	fmt.Println(isValidSudoku(board))
}

func isValidSudoku(board [][]byte) bool {
	rowMap := make(map[int]map[int]bool)  // {row_num : {val : true}}
	colMap := make(map[int]map[int]bool)  // {col_num : {val : true}}
	gridMap := make(map[int]map[int]bool) // {grid_num : {val : true}}
	var gridNum int

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			num := int(board[r][c]) - 48
			if num < 0 || num > 9 {
				continue
			}

			gridNum = 3*(r/3) + c/3 // this will map each grid from 0 to 8.

			if rowMap[r] == nil {
				rowMap[r] = make(map[int]bool)
			}
			if colMap[c] == nil {
				colMap[c] = make(map[int]bool)
			}
			if gridMap[gridNum] == nil {
				gridMap[gridNum] = make(map[int]bool)
			}

			if rowMap[r][num] == true {
				return false
			} else {
				rowMap[r][num] = true
			}

			if colMap[c][num] == true {
				return false
			} else {
				colMap[c][num] = true
			}

			if gridMap[gridNum][num] == true {
				return false
			} else {
				gridMap[gridNum][num] = true
			}
		}
	}
	return true
}
