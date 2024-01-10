package hangman

import (
	"math/rand"
	"strings"
	"time"
)

func NewGamePrep(args []string) (word string, wordRune []rune) {

	word = strings.ToUpper(RandomWord(args[0]))
	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))

	var hiddenLetters, revealedLetters []int

	wordRune = []rune(word)

	for i := 0; i < len(word)/2-1; i++ {
		randTemp := randSource.Intn(len(wordRune))
		if wordRune[randTemp] != 0 {
			revealedLetters = append(revealedLetters, randTemp)
			wordRune[randTemp] = 0
		} else {
			i--
		}
	}
	for j := 0; j < len(wordRune); j++ {
		if wordRune[j] != 0 {
			hiddenLetters = append(hiddenLetters, j)
		}
	}
	for _, i := range revealedLetters {
		wordRune[i] = rune(word[i])
	}
	for _, i := range hiddenLetters {
		wordRune[i] = '_'
	}
	return word, wordRune
}
