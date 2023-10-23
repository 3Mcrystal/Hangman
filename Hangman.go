package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func PlayHangman() {
	word := chooseWord()
	revealedWord := revealLetters(word)
	attempts := 10
	guessedLetters := make(map[rune]bool)

	fmt.Printf("Word: %s\n", revealedWord)

	for attempts > 0 {
		fmt.Print("Guess a letter: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if len(input) != 1 || !isLetter(input) {
			fmt.Println("Invalid input. Please enter a single letter.")
			continue
		}
		letter := []rune(input)[0]

		if guessedLetters[letter] {
			fmt.Printf("You've already guessed the letter '%c'.\n", letter)
			continue
		}

		guessedLetters[letter] = true

		if strings.ContainsRune(word, letter) {
			fmt.Printf("Good guess! '%c' is in the word.\n", letter)
			revealedWord = reveal(word, revealedWord, letter)
		} else {
			attempts--
			fmt.Printf("Wrong guess! Attempts left: %d\n", attempts)
		}

		fmt.Printf("Word: %s\n", revealedWord)

		if revealedWord == word {
			fmt.Printf("Congratulations! You found the word: %s\n", word)
			break
		}
	}

	if revealedWord != word {
		fmt.Printf("Out of attempts. The word was: %s\n", word)
	}
}

func isLetter(s string) bool {
	r := []rune(s)
	return len(r) == 1 && (r[0] >= 'a' && r[0] <= 'z' || r[0] >= 'A' && r[0] <= 'Z')
}

func chooseWord() string {
	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		os.Exit(1)
	}
	defer file.Close()

	words := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))]
}

func revealLetters(word string) string {
	n := len(word)/2 - 1
	revealedWord := make([]byte, len(word))
	for i := range revealedWord {
		revealedWord[i] = '_'
	}
	for i := 0; i < int(n); i++ {
		index := rand.Intn(len(word))
		if revealedWord[index] == '_' {
			revealedWord[index] = word[index]
		}
	}
	return string(revealedWord)
}

func reveal(word, revealed string, letter rune) string {
	revealedBytes := []byte(revealed)
	wordBytes := []byte(word)

	for i, c := range wordBytes {
		if c == byte(letter) {
			revealedBytes[i] = byte(letter)
		}
	}
	return string(revealedBytes)
}
