package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	var board [9][9]int

	for r := 0; r < 9; r++ {
		line := os.Args[r+1]
		if len(line) != 9 {
			fmt.Println("Error")
			return
		}

		for c := 0; c < 9; c++ {
			ch := line[c]
			if ch == '.' {
				board[r][c] = 0
			} else if ch >= '1' && ch <= '9' {
				board[r][c] = int(ch - '0')
			} else {
				fmt.Println("Error")
				return
			}
		}
	}

	if solve(&board) {
		printBoard(board)
	} else {
		fmt.Println("Error")
	}
}

func isValid(board *[9][9]int, row, col, num int) bool {
	for c := 0; c < 9; c++ {
		if board[row][c] == num {
			return false
		}
	}

	for r := 0; r < 9; r++ {
		if board[r][col] == num {
			return false
		}
	}

	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if board[startRow+r][startCol+c] == num {
				return false
			}
		}
	}
	return true
}

func solve(board *[9][9]int) bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == 0 {
				for num := 1; num <= 9; num++ {
					if isValid(board, r, c, num) {
						board[r][c] = num
						if solve(board) {
							return true
						}
						board[r][c] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func printBoard(board [9][9]int) {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			fmt.Print(board[r][c])
			if c < 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
