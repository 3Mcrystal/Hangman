package hangman

import "fmt"

func DisplayHelp() {
	helpText := `
Hangman Game Help:

To start a new game:
	go run main.go <filename> [--letterFile <letterFile>]

Optional Flags:
	--letterFile <letterFile> : Specify the file containing ASCII art for letters.
	--startWith <saveFile>    : Resume the game from a previously saved state.

Game Commands:
	Enter a single letter or a word to make a guess.
	Type "STOP" to save and exit the game.

Example:
	$ go run main.go words.txt --letterFile ascii.txt
	$ go run main.go words.txt --letterFile ascii.txt --startWith save.txt
	$ go run main.go help

Attempts :
	You have ten attempts to beat the game
	A wrong guessed letter count for one attempts
	A wrong guessed words count for two attempts
	

	`
	fmt.Println(helpText)
}
