package infrastructure

import (
	domain "bolaodozeh/domain/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dbURL := "postgres://pg:pass@localhost:5436/bolaodozeh"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&domain.Guess{}, &domain.Match{}, &domain.User{})

	return db
}
