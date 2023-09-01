package adapter

import (
	domain "bolaodozeh/domain/model"
	"bolaodozeh/domain/repository"

	"gorm.io/gorm"
)

type MatchGoOrmAdapter struct {
	db *gorm.DB
}

func (m MatchGoOrmAdapter) FindById(id int) domain.Match {
	var match domain.Match
	m.db.Find(&match, id)
	return match
}

func (m MatchGoOrmAdapter) FindAll() []domain.Match {
	var matches []domain.Match
	m.db.Order("date_time asc").Find(&matches)
	return matches

}

func NewMatchGoOrmAdapter(db *gorm.DB) repository.MatchRepository {
	return MatchGoOrmAdapter{db: db}
}
