package useCase

import (
	domain "bolaodozeh/domain/model"
	"bolaodozeh/domain/repository"

	"encoding/base64"
)

type CreateUserUseCase struct {
	userRepository                 repository.UserRepository
	createBlankGuessForUserUseCase CreateBlankGuessForUserUseCase
}

func NewCreateUserUseCase(userRepository repository.UserRepository, createBlankGuessForUserUseCase CreateBlankGuessForUserUseCase) CreateUserUseCase {
	return CreateUserUseCase{userRepository: userRepository, createBlankGuessForUserUseCase: createBlankGuessForUserUseCase}
}

func (c CreateUserUseCase) Execute(user *domain.User) {
	passwordEncrypted := base64.StdEncoding.EncodeToString([]byte(user.Password))
	user.Password = string(passwordEncrypted)
	user.Id = c.userRepository.Insert(user)
	c.createBlankGuessForUserUseCase.Execute(user)
}
