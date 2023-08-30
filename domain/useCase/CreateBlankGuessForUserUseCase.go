package useCase

import (
	domain "bolaodozeh/domain/model"
	"bolaodozeh/domain/repository"
)

type CreateBlankGuessForUserUseCase struct {
	guessRepository repository.GuessRepository
	matchRepository repository.MatchRepository
}

func NewCreateBlankGuessForUserUseCase(guessRepository repository.GuessRepository, matchRepository repository.MatchRepository) CreateBlankGuessForUserUseCase {
	return CreateBlankGuessForUserUseCase{guessRepository: guessRepository, matchRepository: matchRepository}
}

func (c CreateBlankGuessForUserUseCase) Execute(user *domain.User) {
	matches := c.matchRepository.FindAll()
	for i := 0; i < len(matches); i++ {
		c.guessRepository.Insert(&domain.Guess{MatchId: matches[i].Id, UserId: user.Id})
	}
}
