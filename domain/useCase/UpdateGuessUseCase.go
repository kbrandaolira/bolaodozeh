package useCase

import (
	domain "bolaodozeh/domain/model"
	"bolaodozeh/domain/repository"
	"time"
)

type UpdateGuessUseCase struct {
	guessRepository repository.GuessRepository
	matchRepository repository.MatchRepository
}

func NewUpdateGuessUseCase(guessRepository repository.GuessRepository, matchRepository repository.MatchRepository) UpdateGuessUseCase {
	return UpdateGuessUseCase{guessRepository: guessRepository, matchRepository: matchRepository}
}

func (u UpdateGuessUseCase) Execute(guess *domain.Guess) {
	match := u.matchRepository.FindById(guess.MatchId)
	now := time.Now()
	if now.After(match.DateTime) {
		//TODO validar não é possível palpitar após inicio do jogo
	} else {
		var NilPoints uint
		guess.Points = NilPoints
		u.guessRepository.Update(guess)
	}
}
