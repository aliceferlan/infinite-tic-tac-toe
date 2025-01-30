package main

import (
	"fmt"
)

var LIFETIME = 6

type Cell struct {
	value    int // 0:空白 1:プレイヤー1 2:プレイヤー2
	lifeTime int // 余命
}
type GameBoard struct {
	board [][]Cell
}

type Game struct {
	// history     []GameBoard
	GameBoard
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
	t.board = [][]Cell{
		{Cell{0, 0}, Cell{0, 0}, Cell{0, 0}},
		{Cell{0, 0}, Cell{0, 0}, Cell{0, 0}},
		{Cell{0, 0}, Cell{0, 0}, Cell{0, 0}},
	}
	t.currentTurn = 0
	t.player = []string{"X", "O"}
}

func (t *TicTacToe) checkWin() string {

	for i := 0; i < 3; i++ {
		// 横のチェック
		if t.board[i][0].value == t.board[i][1].value && t.board[i][0].value == t.board[i][2].value && t.board[i][0].value != 0 {
			return t.player[t.board[i][0].value-1]
		}

		// 縦のチェック
		if t.board[0][i].value == t.board[1][i].value && t.board[0][i].value == t.board[2][i].value && t.board[0][i].value != 0 {
			return t.player[t.board[0][i].value-1]
		}
	}

	// 斜めのチェック
	if t.board[0][0].value == t.board[1][1].value && t.board[0][0].value == t.board[2][2].value && t.board[0][0].value != 0 {
		return t.player[t.board[0][0].value-1]
	}
	if t.board[0][2].value == t.board[1][1].value && t.board[0][2].value == t.board[2][0].value && t.board[0][2].value != 0 {
		return t.player[t.board[0][2].value-1]
	}

	// 引き分けのチェック
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if t.board[i][j].value == 0 {
				return ""
			}
		}
	}
	return "none"
}

func (t *TicTacToe) outputBoard() {

	fmt.Print("\033[H\033[2J")

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			switch t.board[i][j].value {
			case 0:
				fmt.Print(" ")
			case 1:
				if t.board[i][j].lifeTime == 1 {
					fmt.Print("Y")
				} else {
					fmt.Print("X")
				}
			case 2:
				if t.board[i][j].lifeTime == 1 {
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

	// 移動が有効かチェック
	if t.board[move[0]][move[1]].value != 0 {
		return fmt.Errorf("invalid move")
	}

	// 現在のボードの全セルのライフタイムを更新
	for i := range t.board {
		for j := range t.board[i] {
			if t.board[i][j].value != 0 { // 空白でないセルのみ
				t.board[i][j].lifeTime--

				// 死んだアイコンを削除
				if t.board[i][j].lifeTime == 0 {
					t.board[i][j].value = 0
				}
			}
		}
	}

	// 移動を適用
	t.board[move[0]][move[1]].value = t.currentTurn%len(t.player) + 1
	t.board[move[0]][move[1]].lifeTime = LIFETIME

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
