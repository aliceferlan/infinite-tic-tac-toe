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

	// チェック用関数
	check := func(a, b, c Cell) bool {
		return a.value == b.value && a.value == c.value && a.value != 0
	}

	for i, row := range t.board {
		// 横のチェック
		if check(row[0], row[1], row[2]) {
			return t.player[row[0].value-1]
		}

		// 縦のチェック
		if check(t.board[0][i], t.board[1][i], t.board[2][i]) {
			return t.player[t.board[0][i].value-1]
		}
	}

	// 斜めのチェック
	if check(t.board[0][0], t.board[1][1], t.board[2][2]) {
		return t.player[t.board[0][0].value-1]
	}
	if check(t.board[0][2], t.board[1][1], t.board[2][0]) {
		return t.player[t.board[0][2].value-1]
	}

	// 引き分けのチェック
	for _, row := range t.board {
		for _, cell := range row {
			if cell.value == 0 {
				return ""
			}
		}
	}
	return "none"
}

func (t *TicTacToe) outputBoard() {

	fmt.Print("\033[H\033[2J")

	for _, row := range t.board {
		for j, cell := range row {
			switch cell.value {
			case 0:
				fmt.Print(" ")
			case 1:
				if cell.lifeTime == 1 {
					fmt.Print("Y")
				} else {
					fmt.Print("X")
				}
			case 2:
				if cell.lifeTime == 1 {
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

	// 移動が有効かチェック
	if t.board[move[0]][move[1]].value != 0 {
		return fmt.Errorf("invalid move")
	}

	// 現在のボードの全セルのライフタイムを更新
	for _, row := range t.board {
		for _, cell := range row {
			if cell.value != 0 { // 空白でないセルのみ
				cell.lifeTime--

				// 死んだアイコンを削除
				if cell.lifeTime == 0 {
					cell.value = 0
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
