package useCase

import (
	"bolaodozeh/application/handler/dto"
	"bolaodozeh/domain/repository"
)

type FindGuessUseCase struct {
	guessRepository repository.GuessRepository
	matchRepository repository.MatchRepository
}

func NewFindGuessUseCase(guessRepository repository.GuessRepository, matchRepository repository.MatchRepository) FindGuessUseCase {
	return FindGuessUseCase{guessRepository: guessRepository, matchRepository: matchRepository}
}

func (f FindGuessUseCase) Execute(userId int) []dto.GuessDto {
	var dtos []dto.GuessDto
	matches := f.matchRepository.FindAll()
	for i := 0; i < len(matches); i++ {
		var dto = dto.GuessDto{
			HomeTeamName: matches[i].HomeTeamName,
			AwayTeamName: matches[i].AwayTeamName,
			Stadium:      matches[i].Stadium,
			Phase:        matches[i].Phase,
			DateTime:     matches[i].DateTime,
			Finished:     (matches[i].HomeTeamScore != nil && matches[i].AwayTeamScore != nil),
		}
		guess := f.guessRepository.FindByMatchIdAndUserId(matches[i].Id, userId)
		dto.GuessId = guess.Id
		dto.HomeTeamGuess = guess.HomeTeamScore
		dto.AwayTeamGuess = guess.AwayTeamScore
		dto.Points = guess.Points
		dtos = append(dtos, dto)
	}
	return dtos
}
