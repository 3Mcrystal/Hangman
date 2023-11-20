package hangman

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadLetterFile(filename string) (map[rune]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	letterMap := make(map[rune]string)
	var currentLetter rune
	var currentLetterLines []string

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		if len(line) == 1 && currentLetter != 0 {
			letterMap[currentLetter] = strings.Join(currentLetterLines, "\n")
			currentLetter = rune(line[0])
			currentLetterLines = nil
		} else if currentLetter != 0 {
			currentLetterLines = append(currentLetterLines, line)
		} else {
			currentLetter = rune(line[0])
			currentLetterLines = nil
		}
	}

	if currentLetter != 0 {
		letterMap[currentLetter] = strings.Join(currentLetterLines, "\n")
	}

	return letterMap, nil
}

func DisplayASCIIArt(word string, letterMap map[rune]string, guessedLetters []rune) {
	for _, letter := range word {
		asciiArt, found := letterMap[letter]
		if !found {
			fmt.Printf("Letter not found for '%c'\n", letter)
			asciiArt = "Letter not found."
		}

		lines := strings.Split(asciiArt, "\n")
		for _, line := range lines {
			fmt.Println(line)
		}
		fmt.Println()
	}

	fmt.Println("Guessed Letters:")
	for _, letter := range guessedLetters {
		fmt.Printf("%c: %s\n", letter, letterMap[letter])
	}
}
