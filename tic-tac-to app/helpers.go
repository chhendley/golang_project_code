package helpers

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	EasyDepth   = 2
	MediumDepth = 4
	HardDepth   = 6
)

type Board [3][3]string

const (
	Empty   = " "
	PlayerX = "X"
	PlayerO = "O"
)

func IsDraw(b Board) bool {
	for _, row := range b {
		for _, cell := range row {
			if cell == Empty {
				return false
			}
		}
	}
	return true
}

func PrintBoard(b Board) {
	for _, row := range b {
		for _, cell := range row {
			fmt.Printf("| %s ", cell)
		}
		fmt.Println("|")
	}
}

func OtherPlayer(player string) string {
	if player == PlayerX {
		return PlayerO
	}
	return PlayerX
}

func CheckWin(b Board, player string) bool {
	// Check rows
	for i := 0; i < 3; i++ {
		if b[i][0] == player && b[i][1] == player && b[i][2] == player {
			return true
		}
	}

	// Check columns
	for i := 0; i < 3; i++ {
		if b[0][i] == player && b[1][i] == player && b[2][i] == player {
			return true
		}
	}

	// Check diagonals
	if b[0][0] == player && b[1][1] == player && b[2][2] == player {
		return true
	}
	if b[0][2] == player && b[1][1] == player && b[2][0] == player {
		return true
	}

	return false
}

func AiMove(b Board) (int, int) {
	for {
		row := rand.Intn(3)
		col := rand.Intn(3)
		if b[row][col] == Empty {
			return row, col
		}
	}
}

func AiMoveMinimax(b Board, difficulty int) (int, int) {
	bestScore := -100
	bestRow, bestCol := -1, -1

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b[i][j] == Empty {
				b[i][j] = PlayerO
				score := alphaBeta(b, difficulty, -math.MaxInt32, math.MaxInt32, false)
				b[i][j] = Empty
				if score > bestScore {
					bestScore = score
					bestRow, bestCol = i, j
				}
			}
		}
	}

	return bestRow, bestCol
}

func alphaBeta(b Board, depth int, alpha int, beta int, isMaximizingPlayer bool) int {
	// Base case: terminal state (win, loss, or draw)
	if CheckWin(b, PlayerX) {
		return -10
	}
	if CheckWin(b, PlayerO) {
		return 10
	}
	if IsDraw(b) || depth == 0 {
		return evaluateBoard(b) // Heuristic evaluation for non-terminal states
	}

	// Recursive case
	if isMaximizingPlayer {
		bestValue := -math.MaxInt32
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if b[i][j] == Empty {
					b[i][j] = PlayerO
					value := alphaBeta(b, depth-1, alpha, beta, false)
					b[i][j] = Empty
					if value > bestValue {
						bestValue = value
					}
					if bestValue > alpha {
						alpha = bestValue
					}
					if beta <= alpha {
						break
					}
				}
			}
		}
		return bestValue
	} else {
		bestValue := math.MaxInt32
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if b[i][j] == Empty {
					b[i][j] = PlayerX
					value := alphaBeta(b, depth-1, alpha, beta, true)
					b[i][j] = Empty
					if value < bestValue {
						bestValue = value
					}
					if bestValue < beta {
						beta = bestValue
					}
					if beta <= alpha {
						break
					}
				}
			}
		}
		return bestValue
	}
}

func evaluateBoard(b Board) int {
	score := 0

	// Reward for three in a row
	for i := 0; i < 3; i++ {
		if b[i][0] == b[i][1] && b[i][1] == b[i][2] {
			if b[i][0] == PlayerO {
				score += 100
			} else if b[i][0] == PlayerX {
				score -= 100
			}
		}
		if b[0][i] == b[1][i] && b[1][i] == b[2][i] {
			if b[0][i] == PlayerO {
				score += 100
			} else if b[0][i] == PlayerX {
				score -= 100
			}
		}
	}

	// Reward for diagonals
	if b[0][0] == b[1][1] && b[1][1] == b[2][2] {
		if b[0][0] == PlayerO {
			score += 100
		} else if b[0][0] == PlayerX {
			score -= 100
		}
	}
	if b[0][2] == b[1][1] && b[1][1] == b[2][0] {
		if b[0][2] == PlayerO {
			score += 100
		} else if b[0][2] == PlayerX {
			score -= 100
		}
	}

	// Reward for center control
	if b[1][1] == PlayerO {
		score += 10
	} else if b[1][1] == PlayerX {
		score -= 10
	}

	// Reward for number of pieces
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b[i][j] == PlayerO {
				score += 1
			} else if b[i][j] == PlayerX {
				score -= 1
			}
		}
	}

	return score
}

func GetUserInput(b Board) (int, int) {
	var row, col int

	for {
		fmt.Print("Enter your move (row and column, separated by a space): ")
		_, err := fmt.Scanf("%d %d", &row, &col)
		if err != nil {
			fmt.Println("Invalid input. Please enter two integers separated by a space.")
			// Clear input buffer to prevent issues with subsequent inputs
			fmt.Scanln()
			continue
		}

		if row < 0 || row >= 3 || col < 0 || col >= 3 {
			fmt.Println("Invalid position. Row and column must be between 0 and 2.")
			continue
		}

		if b[row][col] != Empty {
			fmt.Println("Cell already occupied. Please choose an empty cell.")
			continue
		}

		return row, col
	}
}
