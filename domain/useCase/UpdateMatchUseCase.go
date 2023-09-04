package useCase

import (
	"bolaodozeh/application/handler/dto"
	"bolaodozeh/domain/repository"
	"time"
)

type UpdateMatchUseCase struct {
	matchRepository repository.MatchRepository
}

func NewUpdateMatchUseCase(matchRepository repository.MatchRepository) UpdateMatchUseCase {
	return UpdateMatchUseCase{matchRepository: matchRepository}
}

func (u UpdateMatchUseCase) Execute(dto dto.UpdateMatchDto) {
	match := u.matchRepository.FindById(dto.MatchId)
	match.HomeTeamScore = &dto.HomeTeamScore
	match.AwayTeamScore = &dto.AwayTeamScore
	match.UpdatedAt = time.Now()
	u.matchRepository.Update(&match)
}
