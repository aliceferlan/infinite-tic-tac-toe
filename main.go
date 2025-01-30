package main

import "fmt"

func initialize() [][]int {
	return [][]int{
		{0, 0, 0}, {0, 0, 0}, {0, 0, 0},
	}
}

func checkWin(board [][]int) bool {
	return false
}

func main() {

	history := [][][]int{} // 3D array to store the history of the game
	currentTurn := 0

	player := []string{"X", "O"}

	for true {

		if checkWin(board) {
			fmt.Println("Player " + player[0] + " wins!")
			break
		}

	}

}
