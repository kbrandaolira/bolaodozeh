package adapter

import (
	domain "bolaodozeh/domain/model"
	"bolaodozeh/domain/repository"

	"gorm.io/gorm"
)

type UserGoOrmAdapter struct {
	db *gorm.DB
}

// FindByEmailAndPassword implements repository.UserRepository.
func (u UserGoOrmAdapter) FindByEmailAndPassword(email string, password string) domain.User {
	var user domain.User
	u.db.First(&user).Where("email = ? and password = ?", email, password)
	return user
}

// Insert implements repository.UserRepository.
func (u UserGoOrmAdapter) Insert(user *domain.User) int {
	u.db.Create(user)
	return user.Id
}

func NewUserGoOrmAdapter(db *gorm.DB) repository.UserRepository {
	return UserGoOrmAdapter{db: db}
}
