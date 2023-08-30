package useCase

import (
	domain "bolaodozeh/domain/model"
	"bolaodozeh/domain/repository"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserUseCase struct {
	userRepository                 repository.UserRepository
	createBlankGuessForUserUseCase CreateBlankGuessForUserUseCase
}

func NewCreateUserUseCase(userRepository repository.UserRepository, createBlankGuessForUserUseCase CreateBlankGuessForUserUseCase) CreateUserUseCase {
	return CreateUserUseCase{userRepository: userRepository, createBlankGuessForUserUseCase: createBlankGuessForUserUseCase}
}

func (c CreateUserUseCase) Execute(user *domain.User) {
	passwordEncrypted, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(passwordEncrypted)
	user.Id = c.userRepository.Insert(user)
	c.createBlankGuessForUserUseCase.Execute(user)
}
