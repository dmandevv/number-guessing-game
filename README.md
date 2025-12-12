# Number Guessing Game :joystick:

A simple number guessing game where the computer randomly selects a number and the user has to guess it. The user will be given a limited number of chances to guess the number. If the user guesses the number correctly, the game will end, and the user will win. Otherwise, the game will continue until the user runs out of chances. This project is an exercise from the roadmap.sh project: https://roadmap.sh/projects/expense-tracker 📚

**Prerequisites**

- Go 1.25.1 or newer installed

## Installation 📦

There are a few easy ways to get and run this project.

- **Install using Go**

```shell
go install github.com/dmandevv/number-guessing-game@latest
```
This installs the binary into your `$GOBIN` (or `$GOPATH/bin`) so you can run it globally.

- **Clone and build from source**

```shell
git clone https://github.com/dmandevv/number-guessing-game.git
cd number-guessing-game
go build -o number-guessing-game .
```

After building you can run the generated `number-guessing-game` binary.

- **Run directly (quick, no install)**

```shell
go run .
```
