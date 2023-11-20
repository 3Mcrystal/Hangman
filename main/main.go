package main

import (
	"fmt"
	"hangman"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <filename> [--letterFile <letterFile>] [--startWith <saveFile>]")
	}
	filename := os.Args[1]
	letterFile := ""
	saveFile := ""
	// Check command-line arguments
	for i := 2; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--letterFile":
			if i+1 < len(os.Args) {
				letterFile = os.Args[i+1]
				i++
			} else {
				log.Fatal("Missing argument for --letterFile")
			}
		case "--startWith":
			if i+1 < len(os.Args) {
				saveFile = os.Args[i+1]
				i++
			} else {
				log.Fatal("Missing argument for --startWith")
			}
		}
	}
	var game *hangman.Hangman
	// Load game from save file if provided
	if saveFile != "" {
		game, err := hangman.LoadGame(saveFile)
		if err != nil {
			log.Fatal("Error loading the game:", err)
		}
		if game != nil {
			fmt.Println("Welcome Back, you have", game.Attempts, "attempts remaining.")
			game.Play()
		} else {
			fmt.Println("Error loading the game. Starting a new game.")
		}
	} else {
		words, err := hangman.LoadWordsFromFile(filename)
		if err != nil {
			log.Fatal("Failed to load words from file:", err)
		}
		game = hangman.NewHangman(words, letterFile)
		fmt.Println("Good Luck, you have", game.Attempts, "attempts.")
		game.Play()
	}
}
