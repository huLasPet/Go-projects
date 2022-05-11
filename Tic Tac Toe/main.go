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
var reader = bufio.NewReader(os.Stdin)
var row string
var column string

func drawBoard() {

	fmt.Println(ttt_map["top"]["left"], "|", ttt_map["top"]["middle"], "|", ttt_map["top"]["right"])
	fmt.Println("-----------")
	fmt.Println(ttt_map["middle"]["left"], "|", ttt_map["middle"]["middle"], "|", ttt_map["middle"]["right"])
	fmt.Println("-----------")
	fmt.Println(ttt_map["bottom"]["left"], "|", ttt_map["bottom"]["middle"], "|", ttt_map["bottom"]["right"])
	fmt.Println()
}

func players(player int, symbol string) {
	//Asks the player to specify the spot where they want their mark and checks if that spot is valid and free.
	for {
		drawBoard()
		fmt.Printf("Player %d, enter the row: ", player)
		row, _ = reader.ReadString('\n')
		row = strings.TrimSpace(row)
		if strings.Contains("top middle bottom", row) {
			fmt.Printf("Player %d, enter the column: ", player)
			column, _ = reader.ReadString('\n')
			column = strings.TrimSpace(column)
			if strings.Contains("left middle right", column) {
				if ttt_map[row][column] == " " {
					ttt_map[row][column] = symbol
					break
				} else if ttt_map[row][column] != " " {
					fmt.Println("Spot already taken on the board - select another one please.")
				} else {
					fmt.Println("Invalid entry, try again.")
				}
			} else {
				fmt.Println("Invalid entry, try again.")
			}

		}
	}
}

func checkWin(player string) bool {
	//Check all the win conditions and return true or false
	if ttt_map["top"]["left"] != " " && ttt_map["top"]["left"] == ttt_map["top"]["middle"] && ttt_map["top"]["left"] == ttt_map["top"]["right"] {
		return true
	} else if ttt_map["middle"]["left"] != " " && ttt_map["middle"]["left"] == ttt_map["middle"]["middle"] && ttt_map["middle"]["left"] == ttt_map["middle"]["right"] {
		return true
	} else if ttt_map["bottom"]["left"] != " " && ttt_map["bottom"]["left"] == ttt_map["bottom"]["middle"] && ttt_map["bottom"]["left"] == ttt_map["bottom"]["right"] {
		return true
	} else if ttt_map["top"]["left"] != " " && ttt_map["top"]["left"] == ttt_map["middle"]["middle"] && ttt_map["top"]["left"] == ttt_map["bottom"]["right"] {
		return true
	} else if ttt_map["bottom"]["left"] != " " && ttt_map["bottom"]["left"] == ttt_map["middle"]["middle"] && ttt_map["bottom"]["left"] == ttt_map["top"]["right"] {
		return true
	} else if ttt_map["top"]["left"] != " " && ttt_map["top"]["left"] == ttt_map["middle"]["left"] && ttt_map["top"]["left"] == ttt_map["bottom"]["left"] {
		return true
	} else if ttt_map["top"]["middle"] != " " && ttt_map["top"]["middle"] == ttt_map["middle"]["middle"] && ttt_map["top"]["middle"] == ttt_map["bottom"]["middle"] {
		return true
	} else if ttt_map["top"]["right"] != " " && ttt_map["top"]["right"] == ttt_map["middle"]["right"] && ttt_map["top"]["right"] == ttt_map["bottom"]["right"] {
		return true
	}
	return false
}

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
		if checkWin("1") {
			fmt.Println("Player 1 wins.")
			break
		}
		players(2, "O")
		if checkWin("2") {
			fmt.Println("Player 2 wins.")
			break
		}
	}

}
