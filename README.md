
#  Hangman Game

This is a simple hangman game implemented in Golang.


## Installation

1. Clone the project:

    ```bash
    git clone https://ytrack.learn.ynov.com/git/gyael/hangman
    ```

2. Change into the project directory:

    ```bash
    cd hangman    
    ```
## Features

- Help command
- Word Variety
- ASCII Art (Not implemented)
- TView (Not implemented)
- Save and Resume
- Random Word Selection

## Usage/Examples

Usage
```bash
go run main/main.go <filename> [--letterFile <letterFile>]
```

Usage (for Help)
``` bash
go run main/main.go help
```

Example (without ASCII art)
```bash
go run main/main.go words.txt
```

Example (with ASCII art)
```bash
go run main/main.go words.txt --letterFile standard.txt
```

