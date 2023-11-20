// BlackTrack.go
package hangman

import (
	"fmt"

	"github.com/rivo/tview"
)

// NewHangmanBlackTrack initializes a new HangmanBlackTrack instance.
func NewHangmanBlackTrack(words []string, letterFile string) *HangmanBlackTrack {
	hangmanGame := NewHangman(words, letterFile)
	app := tview.NewApplication()

	// Create TView components
	textView := tview.NewTextView().
		SetText("Welcome to Hangman Black Track!").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetTextAlign(tview.AlignCenter)

	inputField := tview.NewInputField().
		SetLabel("Enter your guess: ").
		SetDoneFunc(func(key tview.Key) {
			if key == tview.KeyEnter {
				guess := inputField.GetText()
				hangmanGame.Play() // Update this line based on your game logic
				updateGameUI(hangmanGame, textView)
				inputField.SetText("") // Clear the input field
			}
		})

	attemptsTextView := tview.NewTextView().
		SetText(fmt.Sprintf("Attempts remaining: %d", hangmanGame.Attempts)).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	// Layout components
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(textView, 0, 1, false).
		AddItem(attemptsTextView, 1, 0, false).
		AddItem(inputField, 3, 0, true)

	// Set the root of the application
	app.SetRoot(flex, true)

	return &HangmanBlackTrack{
		App:              app,
		TextView:         textView,
		InputField:       inputField,
		AttemptsTextView: attemptsTextView,
		Game:             hangmanGame,
	}
}

// Run starts the Hangman Black Track game.
func (hb *HangmanBlackTrack) Run() error {
	updateGameUI(hb.Game, hb.TextView)
	if err := hb.App.Run(); err != nil {
		return err
	}
	return nil
}

// updateGameUI updates the UI with the current state of the Hangman game.
func updateGameUI(game *Hangman, textView *tview.TextView) {
	// Update the text view with the current state of the game
	textView.SetText(fmt.Sprintf("Word: %s\nGuessed Letters: %s", game.DisplayWord, string(game.GuessedLetters)))
}
