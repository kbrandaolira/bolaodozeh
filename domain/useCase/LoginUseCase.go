package useCase

import (
	"bolaodozeh/application/handler/dto"
	"bolaodozeh/domain/repository"
	"encoding/base64"
)

type LoginUseCase struct {
	userRepository repository.UserRepository
}

func NewLoginUseCase(userRepository repository.UserRepository) LoginUseCase {
	return LoginUseCase{userRepository: userRepository}
}

func (l LoginUseCase) Execute(dto *dto.LoginDto) bool {
	passwordEncrypted := base64.StdEncoding.EncodeToString([]byte(dto.Password))
	user := l.userRepository.FindByEmailAndPassword(dto.Email, string(passwordEncrypted))
	return user != nil
}
