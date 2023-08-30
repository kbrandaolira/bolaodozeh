package repository

import domain "bolaodozeh/domain/model"

type GuessRepository interface {
	Insert(guess *domain.Guess) int
}
