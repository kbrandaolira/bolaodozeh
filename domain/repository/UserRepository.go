package repository

import domain "bolaodozeh/domain/model"

type UserRepository interface {
	Insert(user *domain.User) int
	FindByEmailAndPassword(email string, password string) domain.User
}
