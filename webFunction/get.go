package hangman

import "fmt"

func GetInput() (letter string) {
	_, err := fmt.Scan(&letter)
	if err != nil {
		fmt.Println(err)
	}
	return letter
}
