package main

import (
	"fmt"
	"tic-tac-go/helpers"
)

func main() {
	for {
		var playerChoice string
		fmt.Println("Choose X or O:")
		fmt.Scanln(&playerChoice)

		var playerSymbol, aiSymbol string
		if playerChoice == "X" || playerChoice == "x" {
			playerSymbol = helpers.PlayerX
			aiSymbol = helpers.PlayerO
		} else if playerChoice == "O" || playerChoice == "o" {
			playerSymbol = helpers.PlayerO
			aiSymbol = helpers.PlayerX
		} else {
			fmt.Println("Invalid choice. Defaulting to X for player.")
			playerSymbol = helpers.PlayerX
			aiSymbol = helpers.PlayerO
		}

		var difficulty string
		var depth int
		for {
			fmt.Println("Choose difficulty (easy, medium, hard):")
			fmt.Scanln(&difficulty)

			if difficulty == "easy" {
				depth = helpers.EasyDepth
				break
			} else if difficulty == "medium" {
				depth = helpers.MediumDepth
				break
			} else if difficulty == "hard" {
				depth = helpers.HardDepth
				break
			} else {
				fmt.Println("Invalid difficulty. Please choose easy, medium, or hard.")
			}
		}

		rematch := false

		for !rematch {
			// Initialize the board with Empty cells
			b := helpers.Board{}
			for i := range b {
				for j := range b[i] {
					b[i][j] = helpers.Empty
				}
			}
			currentPlayer := playerSymbol

			for {
				helpers.PrintBoard(b)

				if currentPlayer == playerSymbol {
					row, col := helpers.GetUserInput(b)
					b[row][col] = playerSymbol
				} else {
					row, col := helpers.AiMoveMinimax(b, depth)
					b[row][col] = aiSymbol
				}

				if helpers.CheckWin(b, currentPlayer) {
					fmt.Printf("%s wins!\n", currentPlayer)
					break
				} else if helpers.IsDraw(b) {
					fmt.Println("It's a draw!")
					break
				}

				currentPlayer = helpers.OtherPlayer(currentPlayer)
			}

			fmt.Println("Do you want to rematch? (y/n)")
			var input string
			fmt.Scanln(&input)
			if input != "y" && input != "Y" {
				rematch = true
			}
		}
	}
}
