package main

import (
	"fmt"
)

type Cell struct {
	value    int // 0:空白 1:プレイヤー1 2:プレイヤー2
	lifeTime int // 余命
}
type GameBoard struct {
	board [][]Cell
}

type Game struct {
	history     []GameBoard
	currentTurn int
	player      []string
}

type TicTacToe struct {
	Game
}

type GameInterface interface {
	initialize()
	checkWin() string
	outputBoard()
	inputWating() []int
	updateBoard([]int) error

	finishing(string)
}

func (t *TicTacToe) initialize() {
	t.history = append(t.history, GameBoard{board: [][]Cell{
		{Cell{0, 0}, Cell{0, 0}, Cell{0, 0}},
		{Cell{0, 0}, Cell{0, 0}, Cell{0, 0}},
		{Cell{0, 0}, Cell{0, 0}, Cell{0, 0}},
	}})
	t.currentTurn = 0
	t.player = []string{"X", "O"}
}

func (t *TicTacToe) checkWin() string {

	current := &t.history[t.currentTurn].board

	// 横のチェック
	for i := 0; i < 3; i++ {
		if (*current)[i][0].value == (*current)[i][1].value && (*current)[i][0].value == (*current)[i][2].value && (*current)[i][0].value != 0 {
			return t.player[(*current)[i][0].value-1]
		}
	}

	// 縦のチェック
	for i := 0; i < 3; i++ {
		if (*current)[0][i].value == (*current)[1][i].value && (*current)[0][i].value == (*current)[2][i].value && (*current)[0][i].value != 0 {
			return t.player[(*current)[0][i].value-1]
		}
	}

	// 斜めのチェック
	if (*current)[0][0].value == (*current)[1][1].value && (*current)[0][0].value == (*current)[2][2].value && (*current)[0][0].value != 0 {
		return t.player[(*current)[0][0].value-1]
	}
	if (*current)[0][2].value == (*current)[1][1].value && (*current)[0][2].value == (*current)[2][0].value && (*current)[0][2].value != 0 {
		return t.player[(*current)[0][2].value-1]
	}

	// 引き分けのチェック
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (*current)[i][j].value == 0 {
				return ""
			}
		}
	}
	return "none"
}

func (t *TicTacToe) outputBoard() {

	fmt.Print("\033[H\033[2J")
	current := &t.history[t.currentTurn].board

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			switch (*current)[i][j].value {
			case 0:
				fmt.Print(" ")
			case 1:
				if (*current)[i][j].lifeTime == 1 {
					fmt.Print("Y")
				} else {
					fmt.Print("X")
				}
			case 2:
				if (*current)[i][j].lifeTime == 1 {
					fmt.Print("P")
				} else {
					fmt.Print("O")
				}
			}
			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
	}
}

func (t *TicTacToe) inputWating() []int {

	var x, y, xy int

	fmt.Println(t.player[t.currentTurn%len(t.player)], "turn")

	for {
		fmt.Print("Enter the x and y coordinates: ")
		fmt.Scan(&xy)

		x = (xy / 10) - 1
		y = (xy % 10) - 1

		if !(x < 0 || x > 2 || y < 0 || y > 2) {
			break
		}
	}
	return []int{x, y}
}

func (t *TicTacToe) updateBoard(move []int) error {

	//現在のボードをコピー
	currentBoard := make([][]Cell, len(t.history[t.currentTurn].board))
	for i := range t.history[t.currentTurn].board {
		currentBoard[i] = make([]Cell, len(t.history[t.currentTurn].board[i]))
		copy(currentBoard[i], t.history[t.currentTurn].board[i])
	}

	// 移動が有効かチェック
	if currentBoard[move[0]][move[1]].value != 0 {
		return fmt.Errorf("invalid move")
	}

	// 現在のボードの全セルのライフタイムを更新
	for i := range currentBoard {
		for j := range currentBoard[i] {
			if currentBoard[i][j].value != 0 { // 空白でないセルのみ
				currentBoard[i][j].lifeTime--

				// 死んだアイコンを削除
				if currentBoard[i][j].lifeTime == 0 {
					currentBoard[i][j].value = 0
				}
			}
		}
	}

	// 移動を適用
	currentBoard[move[0]][move[1]].value = t.currentTurn%len(t.player) + 1
	currentBoard[move[0]][move[1]].lifeTime = 6

	// 履歴に新しいボードを追加
	t.history = append(t.history, GameBoard{board: currentBoard})
	t.currentTurn++

	return nil
}

func (t *TicTacToe) finishing(winner string) {
	if winner == "none" {
		fmt.Println("Draw")
		return
	}
	fmt.Println(winner, "win")
	fmt.Println("Game Over")
}

func main() {

	var game TicTacToe
	game.initialize()

	game.outputBoard()

	for {
		if winner := game.checkWin(); winner != "" {
			game.finishing(winner)
			break
		}

		input := game.inputWating()
		game.updateBoard(input)
		game.outputBoard()
	}
}
