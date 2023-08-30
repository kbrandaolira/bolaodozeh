package adapter

import (
	"gorm.io/gorm"
)

type adapter struct {
	db *gorm.DB
}

func New(db *gorm.DB) adapter {
	return adapter{db: db}
}
