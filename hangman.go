package hangman

import (
	"fmt"
	"math/rand"
	"strings"
)

func NewHangman(words []string) *Hangman {
	wordToGuess := selectRandomWord(words)

	revealCount := len(wordToGuess)/2 - 1
	displayWord := revealRandomLetters(wordToGuess, revealCount)

	return &Hangman{
		WordToGuess:      wordToGuess,
		DisplayWord:      displayWord,
		Attempts:         10,
		HangmanPositions: readHangmanPositions(),
	}
}

func revealRandomLetters(word string, count int) string {
	if count <= 0 {
		return strings.Repeat("_", len(word))
	}
	revealed := make([]byte, len(word))
	for i := 0; i < len(word); i++ {
		revealed[i] = '_'
	}
	for count > 0 {
		index := rand.Intn(len(word))
		if revealed[index] == '_' {
			revealed[index] = word[index]
			count--
		}
	}
	return string(revealed)
}

func (h *Hangman) Play() {
	fmt.Println("Good Luck, you have 10 attempts.")
	fmt.Println(h.DisplayWord)

	incorrectGuesses := make(map[string]bool) // Keep track of incorrect guesses

	for h.Attempts > 0 {
		fmt.Print("Choose: ")
		guess := readGuess()
		guess = strings.TrimSpace(guess)

		if len(guess) != 1 {
			fmt.Println("Please enter a single letter.")
			continue
		}

		if strings.Contains(h.WordToGuess, guess) {
			h.DisplayWord = updateDisplayWord(h.WordToGuess, h.DisplayWord, guess)
			fmt.Println(h.DisplayWord)
			if h.DisplayWord == h.WordToGuess {
				fmt.Println("Congrats!")
				return
			}
		} else {
			if _, exists := incorrectGuesses[guess]; exists {
				fmt.Printf("You've already guessed '%s' incorrectly.\n", guess)
			} else {
				incorrectGuesses[guess] = true // Record incorrect guess
				h.Attempts--
				displayHangman(h.HangmanPositions, h.Attempts)
				fmt.Printf("Not present in the word, %d attempts remaining\n", h.Attempts)
			}
		}
	}

	fmt.Printf("Out of attempts. The word was: %s\n", h.WordToGuess)
}

func readGuess() string {
	var guess string
	fmt.Scanln(&guess)
	return strings.TrimSpace(guess)
}
