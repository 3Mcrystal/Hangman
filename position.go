package hangman

import (
	"bufio"
	"fmt"
	"os"
)

func readHangmanPositions() []string {
	positions := make([]string, 10)

	file, err := os.Open("hangman.txt")
	if err != nil {
		fmt.Println("Error reading Hangman positions:", err)
		return positions
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	position := ""

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			position += line + "\n"
		} else {
			positions[i] = position
			position = ""
			i++
		}
		if i >= 10 {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading Hangman positions:", err)
	}

	return positions
}

func displayHangman(positions []string, attempts int) {
	if attempts >= 0 && attempts < 10 {
		fmt.Print(positions[9-attempts])
	}
}
