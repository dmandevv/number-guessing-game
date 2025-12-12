package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"slices"
	"strconv"
	"strings"
	"text/tabwriter"
)

var difficultyLevel DifficultyLevel
var guessesLeft int

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Printf("I'm thinking of a number between %v and %v.\n", NUMBER_RANGE_LOW, NUMBER_RANGE_HIGH)
	fmt.Println("Please select the difficulty level:")

	tabWriter := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(tabWriter, "1. %v\t(%v\tchances)\n", getColouredText(EASY), EASY)
	fmt.Fprintf(tabWriter, "2. %v\t(%v\tchances)\n", getColouredText(MEDIUM), MEDIUM)
	fmt.Fprintf(tabWriter, "3. %v\t(%v\tchances)\n", getColouredText(HARD), HARD)
	tabWriter.Flush()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter your choice: ")
		scanner.Scan()
		//retrieve and clean user input
		input := scanner.Text()
		input = strings.ToLower(input)
		input = strings.TrimSpace(input)

		if slices.Contains([]string{"easy", "1"}, input) {
			difficultyLevel = EASY
			break
		} else if slices.Contains([]string{"medium", "2"}, input) {
			difficultyLevel = MEDIUM
			break
		} else if slices.Contains([]string{"hard", "3"}, input) {
			difficultyLevel = HARD
			break
		}
	}

	fmt.Printf("Great! You have selected %v difficulty.\n", getColouredText(difficultyLevel))
	fmt.Println("Let's start the game!")

	guessesLeft = int(difficultyLevel)
	randNumber := NUMBER_RANGE_LOW + rand.IntN(NUMBER_RANGE_HIGH-NUMBER_RANGE_LOW)
	for guessesLeft > 0 {
		fmt.Printf("Enter your guess (%v left): ", guessesLeft)

		scanner.Scan()
		guessText := scanner.Text()
		guessText = strings.TrimSpace(guessText)

		guessNumber, err := strconv.Atoi(guessText)
		if err != nil {
			fmt.Printf("%v is not a valid guess!\n", guessText)
			continue
		}
		if guessNumber < NUMBER_RANGE_LOW || guessNumber > NUMBER_RANGE_HIGH {
			fmt.Printf("%v is not within guessing range: %d - %d\n", guessNumber, NUMBER_RANGE_LOW, NUMBER_RANGE_HIGH)
			continue
		}

		guessesLeft -= 1
		if guessNumber == randNumber {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", int(difficultyLevel)-guessesLeft)
			os.Exit(0)
		} else if guessNumber < randNumber {
			fmt.Printf("Incorrect! The number is greater than %d\n", guessNumber)
		} else if guessNumber > randNumber {
			fmt.Printf("Incorrect! The number is less than %d\n", guessNumber)
		}
	}

	fmt.Printf("Sorry! You're out of chances. The number was: %d\n", randNumber)
}

const (
	GreenText  = "\033[0;32m"
	YellowText = "\033[0;33m"
	RedText    = "\033[0;31m"
	NormalText = "\033[0m" // Reset color
)

func getColouredText(dl DifficultyLevel) string {
	switch dl {
	case EASY:
		return GreenText + dl.String() + NormalText
	case MEDIUM:
		return YellowText + dl.String() + NormalText
	case HARD:
		return RedText + dl.String() + NormalText
	default:
		return dl.String()
	}
}
