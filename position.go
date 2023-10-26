package hangman

import (
	"bufio"
	"fmt"
	"os"
)

func DisplayHangmanPosition(attempts int, hangmanPositions [10]string) {
	if attempts > 10 || attempts < 1 {
		fmt.Println("Nombre d'essais invalide.")
		return
	}

	position := hangmanPositions[10-attempts]

	fmt.Println(position)
}

func LoadHangmanPositions() ([10]string, error) {
	var positions [10]string
	file, err := os.Open("hangman.txt")
	if err != nil {
		return positions, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() && i < 10 {
		positions[i] = scanner.Text()
		i++
	}

	return positions, scanner.Err()
}
