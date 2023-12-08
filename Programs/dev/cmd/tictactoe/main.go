/*
[x][][]
[][x][o]
[][][o]
*/
package main

import (
	"fmt"
	"math/rand"
)

// tab is the short for tablero which means board in spanish
func displayBoard(tab [][]string) {
	for i := 0; i < len(tab); i++ {
		for j := 0; j < len(tab[i]); j++ {
			fmt.Printf("[%v]", tab[i][j])
		}
		fmt.Println()
	}
}

func playerMove(tab [][]string) {
	for i := 0; i < 3; i++ {
		var x, y int
		fmt.Println("Enter the horizontal and vertical coordinates of your move (example: 0 0):")

		_, err := fmt.Scan(&x, &y)

		// if x is equal to = 0, 1, 2 and y is equal to 0, 1, 2 then the move is valid
		if err != nil || x < 0 || x > 2 || y < 0 || y > 2 || tab[x][y] != "" {
			fmt.Println("Invalid move, try again")
			continue
		}

		tab[x][y] = "X"

		break
	}
}

func computerMove(tab [][]string) {
	// the computer will choose a random number between 0 and 2
	for {
		x := rand.Intn(3)
		y := rand.Intn(3)

		if tab[x][y] == "" {
			tab[x][y] = "O"
			break
		}
	}
}

func isGamerOver(tab [][]string) bool {
	winner, _ := checkWin(tab)

	return winner || boardFull(tab)
}

// Checks if the computer has won the game

//make a list of all the win conditions and check if the computer has won

func convertPieceToPlayer(piece string) string {
	if piece == "X" {
		return "You"
	}

	return "Computer"
}

func checkWin(tab [][]string) (bool, string) {
	// check rows
	for i := 0; i < 3; i++ {
		if tab[i][0] != "" &&
			tab[i][0] == tab[i][1] &&
			tab[i][1] == tab[i][2] {
			return true, convertPieceToPlayer(tab[i][0])
		}
	}

	// check columns
	for j := 0; j < 3; j++ {
		if tab[0][j] != "" && tab[0][j] == tab[1][j] && tab[1][j] == tab[2][j] {
			return true, convertPieceToPlayer(tab[0][j])
		}
	}

	if tab[0][0] != "" && tab[0][0] == tab[1][1] && tab[0][0] == tab[2][2] {
		return true, convertPieceToPlayer(tab[0][0])
	}

	if tab[0][2] != "" && tab[0][2] == tab[1][1] && tab[0][2] == tab[2][0] {
		return true, convertPieceToPlayer(tab[0][2])
	}

	return false, ""
}

// Displays who is the winner of the game
func statusMessage(tab [][]string) {
	// it will print "computer-wins" if the check win condition is true
	hasAWinner, player := checkWin(tab)

	if hasAWinner {
		fmt.Printf("Player %s wins\n", player)
		return
	}

	if boardFull(tab) {
		fmt.Println("It's a tie!")
		return
	}

	fmt.Println("Next player's turn")
}

func boardFull(tab [][]string) bool { // checks if the board is full
	for i := 0; i < len(tab); i++ {
		for j := 0; j < len(tab[i]); j++ {
			if tab[i][j] == "" {
				return false
			}
		}
	}

	return true
}

func main() {

	var board = [][]string{
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}

	for !isGamerOver(board) {
		playerMove(board)
		displayBoard(board)
		statusMessage(board)

		if hasAWinner, _ := checkWin(board); hasAWinner || isGamerOver(board) {
			break
		}

		computerMove(board)
		displayBoard(board)
		statusMessage(board)
	}
}
