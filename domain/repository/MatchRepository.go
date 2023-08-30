package repository

import domain "bolaodozeh/domain/model"

type MatchRepository interface {
	FindAll() []domain.Match
}
