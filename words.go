package hangman

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
)

func LoadWordsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			words = append(words, word)
		}
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return words, nil
}
func selectRandomWord(words []string) string {
	randIndex := rand.Intn(len(words))
	return words[randIndex]
}
func updateDisplayWord(word, displayWord, guess string) string {
	if len(word) != len(displayWord) {
		return displayWord
	}
	displayWordSlice := []rune(displayWord)
	for i := range word {
		if i < len(displayWordSlice) && word[i:i+1] == guess {
			displayWordSlice[i] = []rune(guess)[0]
		}
	}
	return string(displayWordSlice)
}
func UniqueLetter(originalWord, revealedWord string, numRevealed int) string {
	rand.Seed(time.Now().UnixNano())
	revealedLetters := make(map[rune]struct{})
	for _, r := range revealedWord {
		if r != '_' {
			revealedLetters[r] = struct{}{}
		}
	}

	var unrevealedLetters []rune
	for _, r := range originalWord {
		if _, revealed := revealedLetters[r]; !revealed {
			unrevealedLetters = append(unrevealedLetters, r)
		}
	}

	randIndices := rand.Perm(len(unrevealedLetters))
	revealedIndices := randIndices[:numRevealed]
	for _, idx := range revealedIndices {
		revealedWord = replaceAtIndex(revealedWord, unrevealedLetters[idx], '_')
	}
	return revealedWord
}

func replaceAtIndex(str string, replacement rune, index int) string {
	if index < 0 || index >= len(str) {
		return str
	}
	runes := []rune(str)
	runes[index] = replacement
	return string(runes)
}
