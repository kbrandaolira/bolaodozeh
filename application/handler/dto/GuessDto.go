package dto

import (
	"time"
)

type GuessDto struct {
	GuessId       int
	HomeTeamName  string
	AwayTeamName  string
	Stadium       string
	Phase         string
	DateTime      time.Time
	Finished      bool
	HomeTeamGuess *int
	AwayTeamGuess *int
	Points        *int
}
