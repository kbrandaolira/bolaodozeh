package adapter

import (
	domain "bolaodozeh/domain/model"
	"bolaodozeh/domain/repository"

	"gorm.io/gorm"
)

type MatchGoOrmAdapter struct {
	db *gorm.DB
}

func (m MatchGoOrmAdapter) FindAll() []domain.Match {
	var matches []domain.Match
	m.db.Find(&matches)
	return matches

}

func NewMatchGoOrmAdapter(db *gorm.DB) repository.MatchRepository {
	return MatchGoOrmAdapter{db: db}
}