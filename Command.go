// In command.go

package hangman

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func ProcessCommand(game *Hangman) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		if command == "STOP" {
			saveGame(game)
			fmt.Println("Game Saved in Save/save.txt.")
			os.Exit(0)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading command:", err)
	}
}

func saveGame(game *Hangman) {
	saveFile := "Save/save.txt"
	saveData, err := json.Marshal(game)
	if err != nil {
		fmt.Println("Error encoding game data:", err)
		return
	}

	err = os.WriteFile(saveFile, saveData, 0644)
	if err != nil {
		fmt.Println("Error saving game data:", err)
	}
}

func LoadGame(filename string) (*Hangman, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var game Hangman
	err = json.Unmarshal(content, &game)
	if err != nil {
		return nil, err
	}

	return &game, nil
}
