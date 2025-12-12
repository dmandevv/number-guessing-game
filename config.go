package main

import "fmt"

var NUMBER_RANGE_LOW int = 1
var NUMBER_RANGE_HIGH int = 100

type DifficultyLevel int

const (
	EASY   = 10
	MEDIUM = 5
	HARD   = 3
)

func (dl DifficultyLevel) String() string {
	switch dl {
	case EASY:
		return "Easy"
	case MEDIUM:
		return "Medium"
	case HARD:
		return "Hard"
	default:
		return fmt.Sprintf("DifficultyLevel(%d)", dl)
	}
}
