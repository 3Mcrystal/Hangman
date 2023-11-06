package hangman

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
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

	displayWordSlice := []rune(displayWord) // Use runes for string manipulation

	for i := range word {
		if i < len(displayWordSlice) && word[i:i+1] == guess {
			displayWordSlice[i] = []rune(guess)[0]
		}
	}

	return string(displayWordSlice)
}
