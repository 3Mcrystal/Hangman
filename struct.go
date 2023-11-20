package hangman

import "github.com/rivo/tview"

type Hangman struct {
	WordToGuess      string
	DisplayWord      string
	Attempts         int
	HangmanPositions []string
	LetterFile       string
	GuessedLetters   []rune
}

type HangmanBlackTrack struct {
	App              *tview.Application
	TextView         *tview.TextView
	InputField       *tview.InputField
	AttemptsTextView *tview.TextView
	Game             *Hangman
}
