package useCase

import (
	"bolaodozeh/application/handler/dto"
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

func (u UpdateGuessUseCase) Execute(dto dto.UpdateGuessDto) {
	guess := u.guessRepository.FindById(dto.GuessId)
	match := u.matchRepository.FindById(guess.MatchId)
	now := time.Now()
	if now.After(match.DateTime) {
		//TODO validar não é possível palpitar após inicio do jogo
	} else {
		guess.HomeTeamScore = &dto.HomeTeamScore
		guess.AwayTeamScore = &dto.AwayTeamScore
		guess.UpdatedAt = time.Now()
		u.guessRepository.Update(&guess)
	}
}
