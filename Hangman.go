package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func PlayHangman(wordFile string) {
	wordList, err := readWordList(wordFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	rand.Seed(time.Now().UnixNano())
	word := selectRandomWord(wordList)

	attempts := 10
	revealedLetters := make([]bool, len(word))
	revealRandomLetters(word, revealedLetters)

	reader := bufio.NewReader(os.Stdin)
	guess := ""

	for attempts > 0 {
		displayWord(word, revealedLetters)

		fmt.Printf("Attempts left: %d\n", attempts)
		fmt.Print("Enter a letter: ")
		guess, _ = reader.ReadString('\n')
		guess = strings.TrimSpace(guess)

		if len(guess) != 1 || !isLetter(guess) {
			fmt.Println("Invalid input. Please enter a single letter.")
			continue
		}

		if strings.Contains(word, guess) {
			fmt.Println("Correct guess!")
			revealLetter(word, revealedLetters, guess)

			if wordFound(revealedLetters) {
				fmt.Printf("Congratulations! You've found the word: %s\n", word)
				return
			}
		} else {
			fmt.Println("Incorrect guess!")
			attempts--
		}
	}

	fmt.Printf("You lose. The correct word was: %s\n", word)
}

func readWordList(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, scanner.Err()
}

func selectRandomWord(wordList []string) string {
	randIndex := rand.Intn(len(wordList))
	return wordList[randIndex]
}

func displayWord(word string, revealedLetters []bool) {
	for i, char := range word {
		if revealedLetters[i] {
			fmt.Printf("%c", char)
		} else {
			fmt.Print("_")
		}
	}
	fmt.Println()
}

func revealLetter(word string, revealedLetters []bool, letter string) {
	for i, char := range word {
		if string(char) == letter {
			revealedLetters[i] = true
		}
	}
}

func revealRandomLetters(word string, revealedLetters []bool) {
	numToReveal := len(word)/2 - 1
	for i := 0; i < numToReveal; i++ {
		for {
			index := rand.Intn(len(word))
			if !revealedLetters[index] {
				revealedLetters[index] = true
				break
			}
		}
	}
}

func isLetter(str string) bool {
	return len(str) == 1 && str >= "a" && str <= "z"
}

func wordFound(revealedLetters []bool) bool {
	for _, revealed := range revealedLetters {
		if !revealed {
			return false
		}
	}
	return true
}
