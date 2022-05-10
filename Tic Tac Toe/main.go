package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var ttt_map = map[string]map[string]string{
	"top":    {"left": " ", "middle": " ", "right": " "},
	"middle": {"left": " ", "middle": " ", "right": " "},
	"bottom": {"left": " ", "middle": " ", "right": " "},
}

func drawBoard() {

	fmt.Println(ttt_map["top"]["left"], "|", ttt_map["top"]["middle"], "|", ttt_map["top"]["right"])
	fmt.Println("-----------")
	fmt.Println(ttt_map["middle"]["left"], "|", ttt_map["middle"]["middle"], "|", ttt_map["middle"]["right"])
	fmt.Println("-----------")
	fmt.Println(ttt_map["bottom"]["left"], "|", ttt_map["bottom"]["middle"], "|", ttt_map["bottom"]["right"])
	fmt.Println()
}

func players(player int, symbol string) {
	//Asks the player to specify the spot where they want their mark and check if that spot if free or not.
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Player %d, enter the row: ", player)
		row, _ := reader.ReadString('\n')
		fmt.Printf("Player %d, enter the column: ", player)
		column, _ := reader.ReadString('\n')
		row = strings.TrimSpace(row)
		column = strings.TrimSpace(column)

		//Add row, column value error checking here
		if ttt_map[row][column] == " " {
			ttt_map[row][column] = symbol
			break
		} else {
			fmt.Println("Spot already taken on the board - select another one please.")
		}
	}

	drawBoard()
}

func checkWin() bool {
	if ttt_map["top"]["left"] == ttt_map["top"]["middle"] == ttt_map["top"]["right"] != " " ||
		ttt_map["middle"]["left"] == ttt_map["middle"]["middle"] == ttt_map["middle"]["right"] != " " ||
		ttt_map["bottom"]["left"] == ttt_map["bottom"]["middle"] == ttt_map["bottom"]["right"] != " " ||
		ttt_map["top"]["left"] == ttt_map["middle"]["middle"] == ttt_map["bottom"]["right"] != " " ||
		ttt_map["top"]["right"] == ttt_map["middle"]["middle"] == ttt_map["bottom"]["left"] != " " ||
		ttt_map["top"]["right"] == ttt_map["middle"]["right"] == ttt_map["bottom"]["right"] != " "|
			ttt_map["top"]["middle"] == ttt_map["middle"]["middle"] == ttt_map["bottom"]["middle"] != " " ||
		ttt_map["top"]["left"] == ttt_map["middle"]["left"] == ttt_map["bottom"]["left"] != " " {
		return true
	}
}

// def check_win():
//     """Check all the possible win combinations, each one in a different line."""
//     if (ttt_map["top"]["left"] == ttt_map["top"]["middle"] == ttt_map["top"]["right"] != " " or
//             ttt_map["middle"]["left"] == ttt_map["middle"]["middle"] == ttt_map["middle"]["right"] != " " or
//             ttt_map["bottom"]["left"] == ttt_map["bottom"]["middle"] == ttt_map["bottom"]["right"] != " " or
//             ttt_map["top"]["left"] == ttt_map["middle"]["middle"] == ttt_map["bottom"]["right"] != " " or
//             ttt_map["top"]["right"] == ttt_map["middle"]["middle"] == ttt_map["bottom"]["left"] != " " or
//             ttt_map["top"]["right"] == ttt_map["middle"]["right"] == ttt_map["bottom"]["right"] != " " or
//             ttt_map["top"]["middle"] == ttt_map["middle"]["middle"] == ttt_map["bottom"]["middle"] != " " or
//             ttt_map["top"]["left"] == ttt_map["middle"]["left"] == ttt_map["bottom"]["left"] != " "):
//         return True

func main() {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
	fmt.Print("----------------------------------------------\n",
		"Welcome to a text based Tic Tac Toe game.\n",
		"The rows are called top, middle and bottom.\n",
		"The columns are called left, middle and right.\n",
		"----------------------------------------------\n")
	for {
		players(1, "X")
		players(2, "O")
	}

}

// while True:
//     players("Player 1", "X")
//     if check_win() is True:
//         print("Player 1 wins.")
//         break
//     players("Player 2", "O")
//     if check_win() is True:
//         print("Player 2 wins.")
//         break
