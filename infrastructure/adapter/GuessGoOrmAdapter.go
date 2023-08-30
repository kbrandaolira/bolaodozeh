package adapter

import (
	domain "bolaodozeh/domain/model"
	"bolaodozeh/domain/repository"

	"gorm.io/gorm"
)

type GuessGoOrmAdapter struct {
	db *gorm.DB
}

func (g GuessGoOrmAdapter) Insert(guess *domain.Guess) int {
	g.db.Create(guess)
	return guess.Id
}

func NewGuessGoOrmAdapter(db *gorm.DB) repository.GuessRepository {
	return GuessGoOrmAdapter{db: db}
}
