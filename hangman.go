package hangman

import (
	"fmt"
	"math/rand"
	"strings"
)

func NewHangman(words []string, letterFile string) *Hangman {
	wordToGuess := selectRandomWord(words)

	revealCount := len(wordToGuess)/2 - 1
	displayWord := revealRandomLetters(wordToGuess, revealCount)

	return &Hangman{
		WordToGuess:      wordToGuess,
		DisplayWord:      displayWord,
		Attempts:         10,
		HangmanPositions: readHangmanPositions(),
		LetterFile:       letterFile,
		GuessedLetters:   []rune{},
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
	fmt.Println("Good Luck, you have", h.Attempts, "attempts.")
	fmt.Println(h.DisplayWord)

	incorrectGuesses := make(map[string]bool)

	for h.Attempts > 0 {
		fmt.Print("Choose: ")
		guess := readGuess()
		guess = strings.TrimSpace(guess)

		if len(guess) == 1 {
			if _, exists := incorrectGuesses[guess]; exists {
				fmt.Printf("You've already guessed the letter '%s' incorrectly.\n", guess)
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
				incorrectGuesses[guess] = true
				h.Attempts--
				displayHangman(h.HangmanPositions, h.Attempts)
				fmt.Printf("Not present in the word, %d attempts remaining\n", h.Attempts)
			}
		} else if len(guess) >= 2 {
			if guess == h.WordToGuess {
				fmt.Println("Congrats! You've guessed the word correctly.")
				return
			} else {
				h.Attempts -= 2
				fmt.Printf("Incorrect word guess, %d attempts remaining\n", h.Attempts)
			}
		} else {
			fmt.Println("Please enter a valid single letter or word.")
		}
	}

	fmt.Printf("Out of attempts. The word was: %s\n", h.WordToGuess)
}

func readGuess() string {
	var guess string
	fmt.Scanln(&guess)
	return strings.TrimSpace(guess)
}
