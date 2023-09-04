package useCase

import (
	"bolaodozeh/application/handler/dto"
	"bolaodozeh/domain/repository"
)

type FindMatchUseCase struct {
	matchRepository repository.MatchRepository
}

func NewFindMatchUseCase(matchRepository repository.MatchRepository) FindMatchUseCase {
	return FindMatchUseCase{matchRepository: matchRepository}
}

func (f FindMatchUseCase) Execute() []dto.MatchDto {
	var dtos []dto.MatchDto
	matches := f.matchRepository.FindAll()
	for i := 0; i < len(matches); i++ {
		var dto = dto.MatchDto{
			MatchId:       matches[i].Id,
			HomeTeamName:  matches[i].HomeTeamName,
			HomeTeamScore: matches[i].HomeTeamScore,
			AwayTeamName:  matches[i].AwayTeamName,
			AwayTeamScore: matches[i].AwayTeamScore,
		}
		dtos = append(dtos, dto)
	}
	return dtos
}
