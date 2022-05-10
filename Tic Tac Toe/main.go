package main

import (
	"fmt"
)

var ttt_dict = map[string]map[string]string{
	"top":    {"left": " ", "middle": " ", "right": " "},
	"middle": {"left": " ", "middle": " ", "right": " "},
	"bottom": {"left": " ", "middle": " ", "right": " "},
}

func drawBoard() {

	fmt.Println(ttt_dict["top"]["left"], "|", ttt_dict["top"]["middle"], "|", ttt_dict["top"]["right"])
	fmt.Println("-----------")
	fmt.Println(ttt_dict["middle"]["left"], "|", ttt_dict["middle"]["middle"], "|", ttt_dict["middle"]["right"])
	fmt.Println("-----------")
	fmt.Println(ttt_dict["bottom"]["left"], "|", ttt_dict["bottom"]["middle"], "|", ttt_dict["bottom"]["right"])
	fmt.Println()
}

func main() {
	drawBoard()
}

// def players(player, symbol):
//     """Asks the player to specify the spot where they want their mark and check if that spot if free or not."""
//     while True:
//         row, column = input(f"{player}: What row and column do you want to put an {symbol} to? ").split()
//         if ttt_dict[row][column] == " ":
//             ttt_dict[row][column] = symbol
//             print()
//             draw_board()
//             break
//         else:
//             print("Spot already taken on the board - select another one please.")
//             continue

// def check_win():
//     """Check all the possible win combinations, each one in a different line."""
//     if (ttt_dict["top"]["left"] == ttt_dict["top"]["middle"] == ttt_dict["top"]["right"] != " " or
//             ttt_dict["middle"]["left"] == ttt_dict["middle"]["middle"] == ttt_dict["middle"]["right"] != " " or
//             ttt_dict["bottom"]["left"] == ttt_dict["bottom"]["middle"] == ttt_dict["bottom"]["right"] != " " or
//             ttt_dict["top"]["left"] == ttt_dict["middle"]["middle"] == ttt_dict["bottom"]["right"] != " " or
//             ttt_dict["top"]["right"] == ttt_dict["middle"]["middle"] == ttt_dict["bottom"]["left"] != " " or
//             ttt_dict["top"]["right"] == ttt_dict["middle"]["right"] == ttt_dict["bottom"]["right"] != " " or
//             ttt_dict["top"]["middle"] == ttt_dict["middle"]["middle"] == ttt_dict["bottom"]["middle"] != " " or
//             ttt_dict["top"]["left"] == ttt_dict["middle"]["left"] == ttt_dict["bottom"]["left"] != " "):
//         return True

// print("Welcome to a text based Tic Tac Toe game. \n"
//       "The rows are called top, middle and bottom.\n"
//       "The columns are called left, middle and right.\n"
//       "To choose where to put your mark type in the row and column separated by a space, for example: 'top middle'.")
// print("-" * 110)

// while True:
//     players("Player 1", "X")
//     if check_win() is True:
//         print("Player 1 wins.")
//         break
//     players("Player 2", "O")
//     if check_win() is True:
//         print("Player 2 wins.")
//         break
