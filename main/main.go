package main

import (
	"hangman"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: hangman <filename>")
	}

	words, err := hangman.LoadWordsFromFile(os.Args[1])
	if err != nil {
		log.Fatal("Failed to load words from file:", err)
	}

	game := hangman.NewHangman(words)

	game.Play()
}
